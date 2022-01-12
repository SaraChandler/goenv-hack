package main

import (
	"fmt"
	"os"

	"github.com/SaraChandler/goenv-hack/cmd"
	"github.com/mitchellh/cli"
)

var version = "v0.0.0"

func main() {
	c := cli.NewCLI("goenv-hack", version)
	c.HelpWriter = os.Stdout
	c.ErrorWriter = os.Stderr

	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"install": func() (cli.Command, error) {
			return &cmd.InstallCommand{}, nil
		},
		"use": func() (cli.Command, error) {
			return &cmd.UseCommand{}, nil
		},
		"list": func() (cli.Command, error) {
			return &cmd.ListCommand{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	os.Exit(exitStatus)
}
