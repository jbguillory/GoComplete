package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Ensure a filename is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer f.Close() // Ensure the file is closed

	// Copy file content to stdout
	_, err = io.Copy(os.Stdout, f)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
