package cmd

import (
	"fmt"

	"github.com/SaraChandler/goenv-hack/pkg"
)

type UninstallCommand struct {
	version string
}

func (i *UninstallCommand) Help() string {
	return "Uninstall <version>\n"
}

func (i *UninstallCommand) Synopsis() string {
	return "Uninstall versions of go"
}

func (i *UninstallCommand) Run(args []string) int {

  if len(args) == 0 {
    fmt.Print(i.Help())
    return 0
  }

  version := args[len(args)-1]

  err := pkg.ValidateVersion(version)
  if err != nil {
    fmt.Println(err)
    return 1
  }

	installedVersions, err := pkg.ListInstalledVersions()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	activeVersion, err := pkg.GetActiveVersion()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	for _, v := range installedVersions {
		if v == version {
      if activeVersion != version {
        err := pkg.UninstallVersion(version)
        if err != nil {
          fmt.Println(err)
          return 1
        }
        return 0
      } else {
        fmt.Println("Requested version to uninstall is currently active. Deactivate to uninstall.\n")
        return 1
      }
		}
	}
  fmt.Printf("Version %v is not installed\n", version)
	return 0
}
