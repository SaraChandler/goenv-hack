package pkg

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// install a specified version of go
func Install(version string) error {
	// download it
	os := runtime.GOOS
	arch := runtime.GOARCH
	url := makeDownloadURL(version, os, arch)

	downloadedFile, err := download(url, getPath(DownloadDir))
	if err != nil {
		return err
	}
	// extract it
	destinationDirectory := filepath.Join(getPath(InstallVersionsDir), version)
	err = extract(downloadedFile, destinationDirectory)
	return err
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
	return []string{"1.16.5"}, nil
}

// Make a download url for a given version, os, and arch
func makeDownloadURL(version string, os string, architecture string) string {
	url := fmt.Sprintf("https://go.dev/dl/go%s.%s-%s.tar.gz", version, os, architecture)
	return url
}
