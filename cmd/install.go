package cmd

type InstallCommand struct {
	Version string
}

func (i *InstallCommand) Help() string {
	return "install <version>"
}

func (i *InstallCommand) Synopsis() string {
	return "Install a version of go"
}

func (i *InstallCommand) Run(args []string) int {
	return 0
}
