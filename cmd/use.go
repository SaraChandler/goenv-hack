package cmd

import "fmt"

type UseCommand struct {
	use string
}

func (i *UseCommand) Help() string {
	return "Use <version>\n"
}

func (i *UseCommand) Synopsis() string {
	return "Use a version of go"
}

func (i *UseCommand) Run(args []string) int {
	// goenv-hack install
	if len(args) == 0 {
		fmt.Print(i.Help())
		return 0
	}

	version := args[len(args)-1]
	// goenv-hack use 1.16
	fmt.Printf("Using go %s\n", version)
	// err := pkg.ValidateVersion(version)

	//err = pkg.Install(version)
	return 0
}
