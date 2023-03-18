package main

import (
	"fmt"

	"github.com/google/go-github/v50/github"
)

const confirmNoBackportLabel = "confirm-no-backport"

func generateBackportLabelForRelease(releaseTag string) string {
	return fmt.Sprintf("backport %s", releaseTag)
}

func checkForLabel(pullRequestLabels []*github.Label, labelToCheck string) bool {
	for _, label := range pullRequestLabels {
		if *label.Name == labelToCheck {
			return true
		}
	}

	return false
}
