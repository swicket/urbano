package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/github"
)

// GitHub the GitHub Provider, that retrieves information directly from alf.io's GitHub repository
type GitHub struct{}

// IsActive - GitHub is always active
func (GitHub) IsActive() bool {
	return true
}

// GetLatestVersion retrieves the latest alf.io version from GitHub
func (GitHub) GetLatestVersion() (string, error) {

	client := github.NewClient(nil)
	releases, _, err := client.Repositories.ListReleases(context.Background(), "alfio-event", "alf.io", nil)

	if err != nil {
		log.Printf("could not fetch version from Github: %v", err)
		return "", err
	}
	release, err := findLatestRelease(releases)
	if err != nil {
		log.Print(err)
		return "", err
	}
	return (*release.TagName), nil
}

func findLatestRelease(releases []*github.RepositoryRelease) (*github.RepositoryRelease, error) {
	for _, r := range releases {
		if !(*r.Draft) && !(*r.Prerelease) {
			return r, nil
		}
	}
	return nil, Error{"Could not find a stable release"}
}

// GetArtifactURL returns the uploaded artifact URL
func (GitHub) GetArtifactURL(version string) string {
	return fmt.Sprintf("https://github.com/alfio-event/alf.io/releases/download/%s/alfio-%s-boot.war", version, version)
}

// Name - the provider's name
func (GitHub) Name() string {
	return "GitHub Alf.io Repository"
}
