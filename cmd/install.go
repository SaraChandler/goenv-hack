package cmd

import "fmt"

type InstallCommand struct {
	version string
}

func (i *InstallCommand) Help() string {
	return "install <version>\n"
}

func (i *InstallCommand) Synopsis() string {
	return "Install a version of go"
}

func (i *InstallCommand) Run(args []string) int {
	// goenv-hack install
	if len(args) == 0 {
		fmt.Print(i.Help())
		return 0
	}

	version := args[len(args)-1]
	// goenv-hack install 1.16
	fmt.Printf("Installing go %s\n", version)
	// err := pkg.ValidateVersion(version)

	//err = pkg.Install(version)
	return 0
}
