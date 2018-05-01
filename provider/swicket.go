package provider

import (
	"fmt"
	"os"
)

type Swicket struct{}

func (Swicket) IsActive() bool {
	return len(os.Getenv("SWICKET_CLIENT_ID")) > 0 && len(os.Getenv("SWICKET_CLIENT_SECRET")) > 0
}

func (Swicket) GetLatestVersion() string {
	return "WIP"
}

func (Swicket) GetArtifactUrl(version string) string {
	return fmt.Sprintf("https://go.swicket.io/%s", version)
}

func (Swicket) Name() string {
	return "Swicket.io"
}
