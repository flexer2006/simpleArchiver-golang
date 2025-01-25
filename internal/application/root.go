// Package application defines the root command and execution logic
// for the simpleArchiver CLI tool. It provides the entry point
// for command execution with integrated error handling.
package application

import (
	"github.com/spf13/cobra"
	"log"
)

// RootCmd represents the base command for the archiver CLI.
// It serves as the parent command for all subcommands (pack/unpack/etc).
//
// Usage: simpleArchiver [command]
// Short: A simple archiver program with VLC encoding support
var RootCmd = &cobra.Command{
	Use:   "simpleArchiver",
	Short: "A simple archiver program",
}

// Execute runs the root command and handles execution errors.
// Wraps the cobra command execution in the application's error handling,
// providing consistent error logging and process exit codes.
//
// Should be called after all subcommands are registered (typically from main).
// Terminates the process with status code 1 if any error occurs.
func Execute() {
	err := HandleError(func() error {
		return RootCmd.Execute()
	})

	if err != nil {
		log.Fatalf("Execution failed with error: %v", err)
	}
}
