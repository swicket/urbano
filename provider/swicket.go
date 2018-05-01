package provider

import (
	"fmt"
	"os"
)

// You could think of embedding the clientID and client_secret into the type
// so that you can build heper methods to fetch them from os.ENV instead of getting them directly
// this could help you validate the Switcket struct earlier on and decouples you from
// how you instantiate the type
type Swicket struct {
	clientID,
	clientSecret string
}

func NewSwicketFromEnv() *Swicket {
	return &Swicket{
		clientID:     os.Getenv("SWICKET_CLIENT_ID"),
		clientSecret: os.Getenv("SWICKET_CLIENT_SECRET"),
	}
}

func (s *Swicket) valid() bool {
	return s.clientID != "" && s.clientSecret != ""
}

func (s Swicket) IsActive() bool {
	return s.valid()
}

func (Swicket) GetLatestVersion() (string, error) {
	return "WIP", nil
}

func (Swicket) GetArtifactUrl(version string) string {
	return fmt.Sprintf("https://go.swicket.io/%s", version)
}

func (Swicket) Name() string {
	return "Swicket.io"
}
