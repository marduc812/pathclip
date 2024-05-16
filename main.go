package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/atotto/clipboard"
)

func main() {
	var copyContent bool
	flag.BoolVar(&copyContent, "c", false, "Copy the content of the file to clipboard")

	flag.Parse()

	if flag.NArg() < 1 {
		displayUsage()
		return
	}

	path := flag.Arg(0)

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

	// Check if the -c flag is provided
	if copyContent {
		copyContent2Clip(path)
	} else {
		copyPath2Clip(path)
	}
}

func copyPath2Clip(path string) {
	copy2Clip(path)
	fmt.Println("copied path of:", path)
}

func copyContent2Clip(path string) {
	// Check if the path is a directory
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("Failed to read file info:", err)
		return
	}
	if info.IsDir() {
		fmt.Println("Cannot copy content of a directory.")
		return
	}

	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Failed to read file content: ", err)
		return
	}
	copy2Clip(string(content))
	fmt.Println("copied content of:", path)
}

func copy2Clip(data string) {
	err := clipboard.WriteAll(data)
	if err != nil {
		fmt.Println("Failed to write to clipboard.", err)
		return
	}
}

func displayUsage() {
	fmt.Println("")
	fmt.Println("pathclip v0.2")
	fmt.Println("")
	fmt.Println("Copy the absolute location of a file to clipboard")
	fmt.Println("Usage: ptc {file}")
	fmt.Println("")
	fmt.Println("Copy the content of the file to clipboard")
	fmt.Println("Usage: ptc [-c] {file}")
}
