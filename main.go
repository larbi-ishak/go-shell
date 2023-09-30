package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	exec.Command("dir /B")
	for {
		fmt.Print("> ")
		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Handle the execution of the input.
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

	}
}

func execInput(input string) error {
	// Remove the newline character.
	input = strings.TrimSpace(input)

	// Prepare the command to execute.
	var cmd *exec.Cmd

	// Determine the operating system and construct the appropriate command.
	// /C and -c are used to run the command in a different terminal and close it when the command finishes
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd.exe", "/C", input)
	case "linux", "darwin": // Linux or macOS
		cmd = exec.Command("bash", "-c", input)
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error.
	return cmd.Run()
}
