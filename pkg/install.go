package pkg

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/coreos/go-semver/semver"
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

func isVersionInList(version string, list []string) (bool, string) {
	versionParts := strings.Split(version, ".")
	twoPartVersion := len(versionParts) == 2
	highestVersion := semver.New("0.0.0")
	for _, n := range list {
		if version == n {
			return true, n
		}

		if twoPartVersion && fuzzyCompare(version, n) {
			highestVersion = semver.New(n)
		}

	}

	if twoPartVersion && highestVersion.String() != "0.0.0" {
		return true, highestVersion.String()
	}
	return false, ""
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
