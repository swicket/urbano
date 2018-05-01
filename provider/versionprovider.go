package provider

type VersionProvider interface {
	IsActive() bool
	GetLatestVersion() string
	GetArtifactUrl(version string) string
	Name() string
}

const Error = "ERROR"
