package pkg

import "fmt"

// install a specified version of go
func Install(version string) error {
	fmt.Printf("Installing version %s\n", version)
	// validate version
	// list downloadable versions
	// find version in the list that matches our version
	// download it
	// extract it
	return nil
}

// List versions available to download from the internet
func ListDownloadableVersions() ([]string, error) {
	return []string{}, nil
}

// Make a download url for a given version, os, and arch
func MakeDownloadURL(version string, os string, architecture string) string {
	url := fmt.Sprintf("https://go.dev/dl/go%s.%s-%s.tar.gz", version, os, architecture)
	return url
}
