package main

import (
	"testing"

	"github.com/google/go-github/v50/github"
)

func TestGennerateBackportLabelForRelease(t *testing.T) {
	expected := "backport 2.0"
	got := generateBackportLabelForRelease("2.0")

	if got != expected {
		t.Fatalf("expected %s got %s", expected, got)
	}
}

func TestCheckForLabel(t *testing.T) {
	labels := []*github.Label{
		{Name: strToPtr("Label 1")},
		{Name: strToPtr("Label 2")},
		{Name: strToPtr("Label 3")},
	}

	t.Run("nonexistent label", func(t *testing.T) {
		labelExists := checkForLabel(labels, "Label 4")
		if labelExists {
			t.Fatal("expected label to not exist")
		}
	})

	t.Run("existing label", func(t *testing.T) {
		labelExists := checkForLabel(labels, "Label 2")
		if !labelExists {
			t.Fatal("expected label to exist")
		}
	})
}

func strToPtr(input string) *string {
	if input == "" {
		return nil
	}
	return &input
}
