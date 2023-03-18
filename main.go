package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/go-github/v50/github"
)

func main() {
	ctx := context.Background()

	config, err := getConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	client := newClient(ctx, config.token)

	eventFile, err := os.Open(config.eventPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var pullRequestEvent github.PullRequestEvent
	if err := json.NewDecoder(eventFile).Decode(&pullRequestEvent); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	releaseTag := os.Getenv("INPUT_RELEASETAG")
	releaseBackportLabel := generateBackportLabelForRelease(releaseTag)

	shouldBackport := checkForLabel(pullRequestEvent.PullRequest.Labels, releaseBackportLabel)
	shouldNotBackport := checkForLabel(pullRequestEvent.PullRequest.Labels, confirmNoBackportLabel)

	// if the PullRequest doesn't have the `confirm-no-backport` label or `backport ${version}` label, we comment on the PullRequest.
	if !shouldBackport && !shouldNotBackport {
		var commentBody = fmt.Sprintf(`❌ Label 'confirm-no-backport' or '%s' is absent

	👉 We're in the next Sourcegraph release code freeze period. If you are 100%% sure your changes should get released or provide no risk to the release, add the 'backport 5.0' label your PR. If you don't want to include this change, add 'confirm-no-backport' to merge into main without backport.
	To learn more about backporting, see the handbook https://handbook.sourcegraph.com/departments/engineering/dev/tools/backport/#how-should-i-use-the-backporting-tool
	`, releaseTag)

		comment := &github.IssueComment{
			Body: github.String(commentBody),
		}
		owner := pullRequestEvent.Repo.Owner.GetLogin()
		repo := pullRequestEvent.Repo.GetName()
		prNumber := pullRequestEvent.Number
		_, _, err := client.Issues.CreateComment(ctx, owner, repo, *prNumber, comment)
		if err != nil {
			fmt.Printf("Error creating comment: %v\n", err)
			return
		}
		os.Exit(1)
	}

	os.Exit(0)
}
