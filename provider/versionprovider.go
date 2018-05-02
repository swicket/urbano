package provider

type VersionProvider interface {
	IsActive() bool
	// always good to return error in its type instead of strings
	GetLatestVersion() (string, error)
	GetArtifactUrl(version string) string
	Name() string
}

// if you want custom errors, you can implement your own imeplemting the error interface

// ProviderError implements the error interface and can be used as en error in function returns
type ProviderError struct {
	something string
}

// just an example to show how inner atteributes can contribute in building an error
func (p ProviderError) Error() string {
	return p.something
}
