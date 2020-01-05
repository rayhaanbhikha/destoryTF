package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	directory := "/Users/rayhaan.bhikha/projects/acc-audit/scripts"
	cmd := exec.Command("ls", "-l", directory)
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
