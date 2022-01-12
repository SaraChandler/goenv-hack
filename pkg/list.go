package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ListInstalledVersions() ([]string, error) {
	versions := []string{}
	files, err := ioutil.ReadDir(getPath(InstallVersionsDir))
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
	linked, err := os.Readlink(getPath(InstallBin))
	if err != nil {
		fmt.Println(err)
		linked = ""
	}

	pathParts := strings.Split(linked, "/")
	if len(pathParts) > 3 {
		return pathParts[len(pathParts)-3], nil
	}

	return "", nil
}
