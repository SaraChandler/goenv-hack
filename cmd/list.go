package cmd

import (
	//"fmt"

	"github.com/mitchellh/cli"
)

type ListCommand struct {
	version string
}

func (i *ListCommand) Help() string {
	return "List <version>\n"
}

func (i *ListCommand) Synopsis() string {
	return "List all installed versions of go"
}

func (i *ListCommand) Run(args []string) int {
	// goenv-hack install
	// if len(args) == 0 {
	// 	fmt.Print(i.Help())
	// 	return 0
	// }
  return cli.RunResultHelp
	//version := args[len(args)-1]
	// goenv-hack install 1.16
	//fmt.Printf("Listing go versions")
	// err := pkg.ValidateVersion(version)

	//err = pkg.Install(version)
	//return 0
}
