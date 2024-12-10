package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/atotto/clipboard"
)

const (
	colorReset  = "\033[0m"  // Reset to default color
	colorGreen  = "\033[32m" // Green for success
	colorYellow = "\033[33m" // Yellow for warnings
	colorRed    = "\033[31m" // Red for errors
)

// Log functions with colors
func logSuccess(message string) {
	fmt.Printf("%s[+]%s %s\n", colorGreen, colorReset, message)
}

func logWarning(message string) {
	fmt.Printf("%s[!]%s %s\n", colorYellow, colorReset, message)
}

func logError(prefix string, err error) {
	fmt.Printf("%s[X]%s %s: %s\n", colorRed, colorReset, prefix, err)
}

// Main function
func main() {
	// Check for missing dependencies on Linux
	checkEnvironment()

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
			logError("The path does not exist", err)
		} else {
			logError("Error", err)
		}
		return
	}

	// It's a relative path
	if !filepath.IsAbs(path) {
		absolutePath, err := filepath.Abs(path)
		if err != nil {
			logError("Error converting to absolute path", err)
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

func checkEnvironment() {
	switch runtime.GOOS {
	case "linux":
		checkLinuxEnvironment()
	case "darwin":
		// mac os has by default pbcopy installed
	case "windows":
		// windows use their own system calls.
	default:
		logError("Unsupported operating system", fmt.Errorf(runtime.GOOS))
	}
}

func checkLinuxEnvironment() {
	// Check if xclip is installed
	if _, err := exec.LookPath("xclip"); err != nil {
		logWarning("xclip is not installed. Install it using your package manager.")
	}

	// Check if DISPLAY is set
	display := os.Getenv("DISPLAY")
	if display == "" {
		logWarning("DISPLAY is not set. Clipboard operations might fail.")
		fmt.Println("In case it fails, set DISPLAY to 0. Run: \033[0mexport DISPLAY=:0\033[0m")
	}
}

func copyPath2Clip(path string) {
	copy2Clip(path)
	logSuccess(fmt.Sprintf("Copied path of: %s", path))
}

func copyContent2Clip(path string) {
	// Check if the path is a directory
	info, err := os.Stat(path)
	if err != nil {
		logError("Failed to read file info", err)
		return
	}
	if info.IsDir() {
		logWarning("Cannot copy content of a directory.")
		return
	}

	content, err := os.ReadFile(path)
	if err != nil {
		logError("Failed to read file content", err)
		return
	}
	copy2Clip(string(content))
	logSuccess(fmt.Sprintf("Copied content of: %s", path))
}

func copy2Clip(data string) {
	err := clipboard.WriteAll(data)
	if err != nil {
		logError("Failed to write to clipboard", err)
		return
	}
}

func displayUsage() {
	fmt.Println("")
	fmt.Println("pathclip v0.3")
	fmt.Println("")
	fmt.Println("Copy the absolute location of a file to clipboard")
	fmt.Println("Usage: ptc {file}")
	fmt.Println("")
	fmt.Println("Copy the content of the file to clipboard")
	fmt.Println("Usage: ptc [-c] {file}")
}
