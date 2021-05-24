package main

import (
	"log"
	"os/exec"
	"strings"
)

func RunSubprocess(cmd string) string {
	// bash -c should be passed for shell commands
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}

	// the output of exec is bytes and need to be converted to string for meaningful usuage
	data := string(out)

	// remove \n from the end of the line if present
	data = strings.TrimSuffix(data, "\n")

	return data
}
