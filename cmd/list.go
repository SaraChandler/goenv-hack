package cmd

import (
	//"fmt"

	"fmt"

	"github.com/SaraChandler/goenv-hack/pkg"
)

type ListCommand struct {
	version string
}

func (i *ListCommand) Help() string {
	return "list\n"
}

func (i *ListCommand) Synopsis() string {
	return "List all installed versions of go"
}

func (i *ListCommand) Run(args []string) int {
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

	fmt.Println("Installed Go Versions:")
	for _, v := range installedVersions {
		isActive := ""
		if activeVersion == v {
			isActive = " * active version"
		}

		fmt.Printf("%s%s\n", v, isActive)
	}

	return 0
}
