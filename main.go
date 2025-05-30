package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		//get current directory
		mainDir, err := getFinalFolderName()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		//get the host name
		host, err := os.Hostname()
		if err != nil {
			fmt.Println(os.Stderr, err)
			return
		}

		// get the current user of the machine
		user_name, err := user.Current()
		if err != nil {
			fmt.Println(os.Stderr, err)
			return
		}

		fmt.Printf("[%s@%s %s]$ ", user_name.Name, host, mainDir)
		//Read Keyboard input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(os.Stderr, err)

		}

		// call execInput function
		if execInput(input); err != nil {
			fmt.Println(os.Stderr, err)
		}
	}
}
func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")

	//split between input commands and arguments related to the commands
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		// "cd" to main directory not supported

		if len(args) < 2 {
			return errors.New("Please specify the path")
		}

		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	//prepare command to execute
	cmd := exec.Command(args[0], args[1:]...)

	//set the correct output
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	//execute a command and return the error.
	return cmd.Run()
}

func getFinalFolderName() (string, error) {

	//get current directory
	dir, err := os.Getwd()
	if err != nil {
		return " ", errors.New("error getting directory")
	}
	// get the final folder name
	mainDir := filepath.Base(dir)

	return mainDir, nil

}
