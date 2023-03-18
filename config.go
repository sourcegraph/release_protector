package main

import (
	"errors"
	"os"
)

type config struct {
	token      string
	eventPath  string
	releaseTag string
}

func getConfig() (config, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return config{}, errors.New("token is not provided")
	}

	eventPath := os.Getenv("GITHUB_EVENT_PATH")
	if eventPath == "" {
		return config{}, errors.New("event path is not provided")
	}

	releaseTag := os.Getenv("INPUT_RELEASETAG")
	if releaseTag == "" {
		return config{}, errors.New("release tag not provided")
	}

	return config{
		token:      token,
		eventPath:  eventPath,
		releaseTag: releaseTag,
	}, nil
}
