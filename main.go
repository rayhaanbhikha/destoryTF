package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
	"path"
)

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
			fmt.Println("directory: ", directory)
			fmt.Println("workspace:", workspace)
			fmt.Println("auto-approve:", autoApprove)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		RED := "\033[0;31m"
		fmt.Printf("\n%s %s\n", RED, err)
	}
	// modules := getModules(workspace)

	// for _, module := range modules {
	// 	destroyResource(workspace, module, directory, autoApprove)
	// }
}

func destroyResource(workspace, moduleToDelete, directory string, autoApprove bool) {
	if moduleToDelete == "ecr" || moduleToDelete == "main" {
		handleErr(fmt.Errorf("Should not destroy %s", moduleToDelete))
	}

	err := os.Chdir(path.Join(directory, moduleToDelete))
	handleErr(err)

	fmt.Println("destroying component ", moduleToDelete)
	runTFCommand("init")
	runTFCommand("workspace", "select", workspace)

	vars := []string{
		"-var",
		"DRONE_BUILD_NUMBER=${DRONE_BUILD_NUMBER}",
		"-var",
		fmt.Sprintf("domain_prefix=bb-%s", workspace),
	}
	if autoApprove {
		vars = append(vars, "-auto-approve")
		args := append([]string{"destroy"}, vars...)
		runTFCommand(args...)
	} else {
		vars = append(vars, "-destroy")
		args := append([]string{"plan"}, vars...)
		runTFCommand(args...)
	}
}

func runTFCommand(args ...string) {
	cmd := exec.Command("terraform", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	fmt.Println(err)
}
