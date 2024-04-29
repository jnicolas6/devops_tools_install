package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// Array of commands to be executed
	commands := []string{
		"jenkins.sh",
		"maven.sh",
		"trivy.sh",
		"docker.sh",
		"sonarqube.sh",
	}

	// Loop through each command and execute it
	for _, cmd := range commands {
		// Split the command string into command and arguments
		parts := strings.Fields(cmd)
		command := parts[0]
		args := parts[1:]

		// Executing the command
		out, err := exec.Command(command, args...).CombinedOutput()
		if err != nil {
			fmt.Printf("Error executing command '%s': %s\n", cmd, err)
			continue
		}

		// Print the output of the command
		fmt.Printf("Output of '%s':\n%s\n", cmd, out)
	}
}

