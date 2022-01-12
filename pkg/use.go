package pkg

import (
	"os"
	"path/filepath"
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
	return os.Symlink(installedBin, getPath(InstallBin))
}
