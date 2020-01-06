package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func main() {

	directory := "/Users/rayhaan.bhikha/projects/acc-audit/terraform/branch-builds"
	workspace := "262"
	modules := getModules(workspace)
	autoApprove := true

	for _, module := range modules {
		destroyResource(workspace, module, directory, autoApprove)
	}
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
