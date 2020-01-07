package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

const red string = "\033[01;31m"
const green = "\033[01;32m"

func main() {
	var (
		directory   string
		workspace   string
		autoApprove bool
	)
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "directory",
			Aliases:     []string{"d"},
			Usage:       "specify `directory` of top level branch-builds tf plan",
			Required:    true,
			Destination: &directory,
		},
		&cli.StringFlag{
			Name:        "workspace",
			Aliases:     []string{"w"},
			Usage:       "specify `workspace` to run tf plan",
			Required:    true,
			Destination: &workspace,
		},
		&cli.BoolFlag{
			Name:        "auto-approve",
			Aliases:     []string{"a"},
			Value:       false,
			Usage:       "destroy terraform resources without confirmation",
			Destination: &autoApprove,
		},
	}

	app := &cli.App{
		Usage: "Destroy terraform resources in a given workspace",
		Flags: flags,
		Action: func(c *cli.Context) error {
			fmt.Printf("\n%s Directory: %s", green, directory)
			fmt.Printf("\n%s Workspace: %s", green, workspace)
			fmt.Printf("\n%s AutoApprove: %t\n", green, autoApprove)
			return destroyResources(directory, workspace, autoApprove)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("\n%s %s\n", red, err)
	}
}
