package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n') // \n is the delimiter, read until you find a new line, you can use \t
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// always trim the spaces from the input
	input = strings.TrimSpace(input)

	switch input {
	case "pwd":
		pwd()
	case "ls":
		ls()
	default:
		fmt.Println("Command not found")
	}
}

func pwd() {
	currentDir, _ := os.Getwd()
	fmt.Println(currentDir)
}

func ls() {
	currentDir, _ := os.Getwd()
	// get files and directories in the current directory
	files, err := os.ReadDir(currentDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println(file.Name(), "dir")
		} else {
			fmt.Println(file.Name(), "file")
		}
	}
}
