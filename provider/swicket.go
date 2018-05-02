package provider

import (
	"fmt"
	"os"
)

// Swicket : Provider for Swicket.io
type Swicket struct {
	clientID,
	clientSecret string
}

// NewSwicket Creates a new instance of Swicket provider using env variables
func NewSwicket() *Swicket {
	return &Swicket{
		clientID:     os.Getenv("SWICKET_CLIENT_ID"),
		clientSecret: os.Getenv("SWICKET_CLIENT_SECRET"),
	}
}

func (s *Swicket) valid() bool {
	return s.clientID != "" && s.clientSecret != ""
}

// IsActive tests whether or not this provider can be used
func (s Swicket) IsActive() bool {
	return s.valid()
}

// GetLatestVersion retrieves the latest available version of alf.io
// by checking on Swicket.io server. Swicket allows to schedule an installation so that
// the new version will be visible only from a given time, and this process can be completely
// automatized.
func (Swicket) GetLatestVersion() (string, error) {
	return "WIP", nil
}

// GetArtifactURL returns the artifact URL. This URL can either redirect to the GitHub repository or to an s3 bucket, depending on
// the configuration.
func (Swicket) GetArtifactURL(version string) string {
	return fmt.Sprintf("https://go.swicket.io/%s", version)
}

// Name provider's name
func (Swicket) Name() string {
	return "Swicket.io"
}
