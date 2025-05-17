package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Set up logging to file
	logPath := filepath.Join(filepath.Dir(os.Args[0]), "linux_loader.log")
	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	// Configure the default logger to write to the log file
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	const skse = "skse64_loader.exe"
	const fose = "fose_loader.exe"
	var loaderPath = ""

	// Get the full path to our own executable
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to get executable path: %v", err)
	}
	exeDir := filepath.Dir(exePath)

	// Use full paths for the loaders
	sksePath := filepath.Join(exeDir, skse)
	fosePath := filepath.Join(exeDir, fose)

	if _, err := os.Stat(sksePath); err == nil {
		loaderPath = sksePath
		log.Printf("Found loader: %s", sksePath)
	} else if _, err := os.Stat(fosePath); err == nil {
		loaderPath = fosePath
		log.Printf("Found loader: %s", fosePath)
	} else {
		log.Fatalf("Failed to find skse64_loader.exe or fose_loader.exe in %s", exeDir)
		fmt.Scanln()
		os.Exit(1)
	}

	// Set working directory to the executable directory
	if err := os.Chdir(exeDir); err != nil {
		log.Fatalf("Failed to change working directory: %v", err)
	}
	log.Printf("Changed working directory to: %s", exeDir)

	// Collect all command-line arguments passed to this program
	args := os.Args[1:]
	log.Printf("Launching %s with arguments: %v", loaderPath, args)

	// Create the command to run loader with the same arguments
	cmd := exec.Command(loaderPath, args...)

	// Create pipes for stdout and stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Failed to create stdout pipe: %v", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalf("Failed to create stderr pipe: %v", err)
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		log.Fatalf("Failed to start command: %v", err)
	}

	// Create a channel to signal when we're done reading output
	done := make(chan bool)

	// Read stdout and stderr in goroutines
	go func() {
		scanner := make([]byte, 1024)
		for {
			n, err := stdout.Read(scanner)
			if err != nil {
				if err != io.EOF {
					log.Printf("Error reading stdout: %v", err)
				}
				break
			}
			output := string(scanner[:n])
			log.Printf("STDOUT: %s", output)
			fmt.Print(output) // Still show in console
		}
		done <- true
	}()

	go func() {
		scanner := make([]byte, 1024)
		for {
			n, err := stderr.Read(scanner)
			if err != nil {
				if err != io.EOF {
					log.Printf("Error reading stderr: %v", err)
				}
				break
			}
			output := string(scanner[:n])
			log.Printf("STDERR: %s", output)
			fmt.Fprintf(os.Stderr, "%s", output) // Still show in console
		}
		done <- true
	}()

	// Wait for both stdout and stderr to be fully read
	<-done
	<-done

	// Wait for the command to complete
	err = cmd.Wait()

	// Check the exit code
	var exitError *exec.ExitError
	if errors.As(err, &exitError) {
		// Program exited with error code
		log.Printf("Program exited with error code: %d", exitError.ExitCode())
		os.Exit(exitError.ExitCode())
	}

	log.Printf("Program completed successfully")
	// Otherwise exit normally
	os.Exit(0)
}
