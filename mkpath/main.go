package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mkdwrap <directory>")
		return
	}

	dir := os.Args[1]
	parts := strings.Split(dir, string(filepath.Separator))
	var failedPath string
	var errorMsg string

	red := color.New(color.FgRed, color.Underline).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	divider := color.New(color.Faint).SprintFunc()("/")

	for i := range parts {
		currentPath := filepath.Join(parts[:i+1]...)
		if _, err := os.Stat(currentPath); os.IsNotExist(err) {
			err := os.Mkdir(currentPath, os.ModePerm)
			if err != nil {
				failedPath = currentPath
				errorMsg = fmt.Sprintf("Error creating directory %s: %s\n", currentPath, err)
				break
			}
		} else if i == len(parts)-1 {
			failedPath = currentPath
			errorMsg = fmt.Sprintf("Error: Directory %s already exists\n", parts[i])
			break
		}
	}

	// Generate the final output with colors
	var output strings.Builder
	for i, part := range parts {
		if i > 0 {
			output.WriteString(divider)
		}
		if filepath.Join(parts[:i+1]...) == failedPath {
			output.WriteString(red(part))
		} else {
			output.WriteString(blue(part))
		}
	}

	if failedPath == "" {
		output.WriteString(".. Success")
		fmt.Println(output.String())
	} else {
		output.WriteString(".. Failed")
		fmt.Println(output.String())
		fmt.Print(errorMsg)
	}
}
