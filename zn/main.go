package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("ZN: wrapper for ln -s")
		fmt.Println("Usage: <source> <target> - no need to write the absolute path, it it automatically resolved")
		os.Exit(1)
	}

	source, err := filepath.Abs(os.Args[1])
	if err != nil {
		fmt.Printf("Error resolving absolute path for source: %s\n", err)
		os.Exit(1)
	}

	target, err := filepath.Abs(os.Args[2])
	if err != nil {
		fmt.Printf("Error resolving absolute path for target: %s\n", err)
		os.Exit(1)
	}

	cmd := exec.Command("ln", "-s", source, target)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing ln -s command: %s\n", err)
		os.Exit(1)
	}
}
