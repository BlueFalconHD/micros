package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// fileSizeTable prints the file size in a table format up to the largest relevant unit.
func fileSizeTable(filename string) error {
	file, err := os.Stat(filename)
	if err != nil {
		return err
	}

	size := file.Size()
	boldCyan := color.New(color.FgCyan).Add(color.Bold)
	boldCyan.Printf("Size of %s:\n", filename)
	fmt.Println("----------------------------")

	printSize := func(label string, size int64) {
		if size > 0 {
			boldCyan.Printf("%-12s : %d\n", label, size)
		}
	}

	bytes := size % 1024
	kilobytes := (size / 1024) % 1024
	megabytes := (size / (1024 * 1024)) % 1024
	gigabytes := size / (1024 * 1024 * 1024)

	if gigabytes > 0 {
		printSize("Gigabytes", gigabytes)
		printSize("Megabytes", megabytes)
		printSize("Kilobytes", kilobytes)
	} else if megabytes > 0 {
		printSize("Megabytes", megabytes)
		printSize("Kilobytes", kilobytes)
	} else if kilobytes > 0 {
		printSize("Kilobytes", kilobytes)
	}

	printSize("Bytes", bytes)
	fmt.Println("----------------------------")

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	filename := os.Args[1]
	err := fileSizeTable(filename)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
