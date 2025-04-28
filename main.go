package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Get the full path to our own executable
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to get executable path: %v", err)
	}
	exeDir := filepath.Dir(exePath)

	// Build the path to the real skse64_loader.exe
	skseLoaderPath := filepath.Join(exeDir, "skse64_loader.exe")

	// Collect all command-line arguments passed to this program
	args := os.Args[1:]

	// Create the command to run skse64_loader.exe with the same arguments
	cmd := exec.Command(skseLoaderPath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Run the command
	err = cmd.Run()

	// Check the exit code
	if exitError, ok := err.(*exec.ExitError); ok {
		// Program exited with error code
		os.Exit(exitError.ExitCode())
	} else if err != nil {
		// Some other error (e.g. couldn't start)
		log.Fatalf("Failed to run skse64_loader.exe: %v", err)
	}

	// Otherwise exit normally
	os.Exit(0)
}
