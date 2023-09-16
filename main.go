package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/atotto/clipboard"
)

func main() {

	if len(os.Args) < 2 {
		displayUsage()
		return
	}

	path := os.Args[1]

	// Check if the path exists
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("The path does not exist.")
		} else {
			fmt.Println("Error:", err.Error())
		}
		return
	}

	// It's a relative path
	if !filepath.IsAbs(path) {
		absolutePath, err := filepath.Abs(path)
		if err != nil {
			fmt.Println("Error converting to absolute path:", err)
			return
		}
		path = absolutePath
	}

	// Check if it's a file or a directory
	if info.IsDir() {
		path += string(filepath.Separator)
	}

	copy2Clip(path)
	fmt.Println("copied: " + path)
}

func copy2Clip(path string) {
	err := clipboard.WriteAll(path)
	if err != nil {
		fmt.Println("Failed to write to clipboard." + err.Error())
		return
	}
}

func displayUsage() {
	fmt.Println("")
	fmt.Println("pathclip v.0.1")
	fmt.Println("")
	fmt.Println("Copy the absolute location of a file to clipboard")
	fmt.Println("")
	fmt.Println("Usage: pathclip")
	fmt.Println(os.Args[0] + " {selected file or folder}")
}
