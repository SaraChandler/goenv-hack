package pkg

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func Use(version string) error {
	installed, err := ListInstalledVersions()
	if err != nil {
		return err
	}

	isInstalled, err := isVersionInList(version, installed)
	if err != nil {
		return err
	}

	if !isInstalled {
		err = Install(version)
		if err != nil {
			return err
		}
	}

	// Remove link
	_ = os.Remove(getPath(InstallBin))

	// Create new symlink
	installedBin := filepath.Join(getPath(InstallVersionsDir), version, "go", "bin")
	fmt.Printf("Using go version %s\n", version)
	return os.Symlink(installedBin, getPath(InstallBin))
}

// Automatically select a go version to use
func UseAuto() error {
	gomod, err := os.Open("go.mod")
	if err != nil {
		return err
	}
	defer gomod.Close()
	modVersion, err := checkGoMod(gomod)
	if err != nil {
		return err
	}

	return Use(modVersion)
}

func checkGoMod(modfile *os.File) (string, error) {
	scanner := bufio.NewScanner(modfile)
	re := regexp.MustCompile("^go 1.*")

	for scanner.Scan() {
		line := scanner.Text()
		if re.MatchString(line) {
			version := strings.TrimPrefix(line, "go ") // Remove "go " from "go 1.xx"
			return version, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("Unable to determine version in go mod")
}
