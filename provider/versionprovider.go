package provider

// VersionProvider represents a provider for retrieving both the latest alf.io stable version and
// the binary artifact
type VersionProvider interface {
	IsActive() bool
	GetLatestVersion() (string, error)
	GetArtifactURL(version string) string
	Name() string
}

// Error implements the error interface and can be used as en error in function returns
type Error struct {
	detail string
}

func (p Error) Error() string {
	return p.detail
}
