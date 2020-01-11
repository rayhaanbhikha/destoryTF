package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"sync"
)

func destroyResources(directory, workspace string, autoApprove bool) error {
	modules := getModules(workspace)
	if len(modules) > 0 {
		fmt.Println("deleting modules: ", modules)
	}

	possibleErrors := make(chan error, len(modules))

	var wg sync.WaitGroup

	for _, module := range modules {
		wg.Add(1)
		go func(module string) {
			defer wg.Done()
			err := destroyResource(workspace, module, directory, autoApprove)
			if err != nil {
				possibleErrors <- err
				return
			}
			fmt.Printf("\n%s successfully destroyed => %s\n", green, module)
		}(module)
	}
	wg.Wait()
	close(possibleErrors)

	return generateErrors(possibleErrors)
}

func generateErrors(possibleErrors <-chan error) error {
	foundErrors := ""
	for err := range possibleErrors {
		foundErrors += fmt.Sprintf(" %s \n", err.Error())
	}

	if foundErrors != "" {
		return errors.New(foundErrors)
	}
	return nil
}

func destroyResource(workspace, moduleToDelete, directory string, autoApprove bool) error {
	if moduleToDelete == "ecr" || moduleToDelete == "main" {
		return fmt.Errorf("Should not destroy %s", moduleToDelete)
	}

	// get original dir and defer to change back to original dir.
	originalDir, err := os.Getwd()
	if err != nil {
		return err
	}
	defer func() {
		os.Chdir(originalDir)
	}()

	// change to dir with terraform plan.
	err = os.Chdir(path.Join(directory, moduleToDelete))
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
	args = append(args, "-lock=false")

	return runTFCommand(args...)
}

func runTFCommand(args ...string) error {
	cmd := exec.Command("terraform", args...)
	// cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
