package pkg

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func ListInstalledVersions() ([]string, error) {
	versions := []string{}
	files, err := ioutil.ReadDir(InstallVersionsDir)
	if err != nil {
		return []string{}, err
	}

	for _, file := range files {
		if file.IsDir() {
			versions = append(versions, file.Name())
		}
	}
	return versions, nil
}

func GetActiveVersion() (string, error) {
	linked, err := os.Readlink(InstallBin)
	if err != nil {
		return "", err
	}

	pathParts := filepath.SplitList(linked)
	if len(pathParts) > 2 {
		return pathParts[len(pathParts)-2], nil
	}

	return "", nil
}
