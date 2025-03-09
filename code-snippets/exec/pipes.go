package main

import (
	"fmt"
	"os"
	"os/exec"
	"io"
)

func main() {
	command := os.Args[1]
	fmt.Println("Running command: ", command)
	args := os.Args[2:len(os.Args)]
	fmt.Println("Args: ", args)
	
	cmd := exec.Command(command, args...)
	
	stdout_pipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Err: %v", err)
		os.Exit(1)
	}

	stderr_pipe, err := cmd.StderrPipe()
	if err != nil {
		fmt.Printf("Err: %v", err)
		os.Exit(1)
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("Err: %v", err)
		os.Exit(1)
	}

	out, err := io.ReadAll(stdout_pipe)
	if err != nil {
		fmt.Printf("Err: %v", err)
		os.Exit(1)
	}

	err_out, err := io.ReadAll(stderr_pipe)
	if err != nil {
		fmt.Printf("Err: %v", err)
		os.Exit(1)
	}
	cmd.Wait()

	fmt.Println("Error: ", string(err_out))
	fmt.Println("Output: ", string(out))

}