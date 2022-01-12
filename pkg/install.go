package pkg

import (
	"fmt"
	"log"
	"runtime"
)

// install a specified version of go
func Install(version string) error {
	fmt.Printf("Installing version %s\n", version)

	// validate version
	err := ValidateVersion(version)
	fmt.Printf("Printing validate version response= %s\n", err)
	if err != nil {
		return err
	}
	// list downloadable versions
  list, err := listDownloadableVersions();

	if err != nil {
		return err
	}

	fmt.Printf("list downloadable version response = %s\n %s", list, err)

	// find version in the list that matches our version
  downloadable, err := isVersionInList(version, list)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("IsVersionDownloadable returns= %s\n %s", downloadable, err)
	// download it
	os := runtime.GOOS
	arch := runtime.GOARCH
	url := makeDownloadURL(version, os, arch)
	// extract it
	return nil
}

func isVersionInList(version string, list []string) (bool, error) {
	for _, n := range list {
		if version == n {
				return true, nil
		}
  }
  return false, fmt.Errorf("Unable to find matching version: %v", version)
}

// List versions available to download from the internet
func listDownloadableVersions() ([]string, error) {
	return []string{}, nil
}

// Make a download url for a given version, os, and arch
func makeDownloadURL(version string, os string, architecture string) string {
	url := fmt.Sprintf("https://go.dev/dl/go%s.%s-%s.tar.gz", version, os, architecture)
	return url
}
