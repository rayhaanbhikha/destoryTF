package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
	"path"
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

			modules := getModules(workspace)
			if len(modules) > 0 {
				fmt.Println("deleting modules: ", modules)
			}
			for _, module := range modules {
				err := destroyResource(workspace, module, directory, autoApprove)
				if err != nil {
					return err
				}
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("\n%s %s\n", red, err)
	}
}

func destroyResource(workspace, moduleToDelete, directory string, autoApprove bool) error {
	if moduleToDelete == "ecr" || moduleToDelete == "main" {
		return fmt.Errorf("Should not destroy %s", moduleToDelete)
	}

	// change to dir with terraform plan.
	err := os.Chdir(path.Join(directory, moduleToDelete))

	if err != nil {
		return err
	}

	fmt.Println("destroying component ", moduleToDelete)

	// initialise terraform
	err = runTFCommand("init")
	if err != nil {
		return err
	}

	// select terraform workspace
	err = runTFCommand("workspace", "select", workspace)
	if err != nil {
		return err
	}

	// generate terraform vars and destroy/plan
	vars := []string{
		"-var",
		"DRONE_BUILD_NUMBER=${DRONE_BUILD_NUMBER}",
		"-var",
		fmt.Sprintf("domain_prefix=bb-%s", workspace),
	}
	var args []string
	if autoApprove {
		vars = append(vars, "-auto-approve")
		args = append([]string{"destroy"}, vars...)
	} else {
		vars = append(vars, "-destroy")
		args = append([]string{"plan"}, vars...)
	}

	return runTFCommand(args...)
}

func runTFCommand(args ...string) error {
	cmd := exec.Command("terraform", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
