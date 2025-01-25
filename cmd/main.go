// Package main provides the entry point for the simpleArchiver application.
// It initializes and executes the command-line interface for archiving operations.
// The main function handles panics gracefully using the application's panic handler
// and sets up the necessary commands before executing the application logic.
package main

import (
	"github.com/flexer2006/simpleArchiver-golang/cmds"
	"github.com/flexer2006/simpleArchiver-golang/internal/application"
)

// main is the entry point of the application. It:
// 1. Sets up panic recovery using application.HandlePanic
// 2. Initializes the CLI commands using cmds.InitCommands()
// 3. Executes the root command using application.Execute()
// The panic handler ensures any unexpected errors are properly logged and handled.
func main() {
	application.HandlePanic(func() {
		cmds.InitCommands()
		application.Execute()
	})
}
