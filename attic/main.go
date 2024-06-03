package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	AtticRoot string `yaml:"attic_root"`
}

func readConfig(configPath string) (Config, error) {
	var config Config
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	return config, err
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: attic <file_to_store>")
		os.Exit(1)
	}

	fileToStore := os.Args[1]
	configPath := filepath.Join(os.Getenv("HOME"), ".config", "attic.yml")

	config, err := readConfig(configPath)
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		os.Exit(1)
	}

	atticRoot := config.AtticRoot
	os.Setenv("attic_root", atticRoot)

	absFileToStore, err := filepath.Abs(fileToStore)
	if err != nil {
		fmt.Printf("Error getting absolute path: %v\n", err)
		os.Exit(1)
	}

	// Create the destination path by replacing the root with the attic root
	destPath := filepath.Join(atticRoot, absFileToStore[1:]) // removing the leading '/'

	// Create necessary directories
	destDir := filepath.Dir(destPath)
	err = os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating directories: %v\n", err)
		os.Exit(1)
	}

	// Move the file
	err = os.Rename(absFileToStore, destPath)
	if err != nil {
		fmt.Printf("Error moving file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Moved %s to %s\n", absFileToStore, destPath)
}
