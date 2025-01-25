// Package application provides core command structure and initialization logic
// for CLI operations. This file specifically handles the setup of the 'pack' command.
package application

import (
	"github.com/spf13/cobra"
)

// vlcPackCmd represents the 'pack' subcommand for file compression operations.
// It is automatically registered to the root command during package initialization.
//
// Usage: pack [flags]
// Short: Pack files using VLC encoding
var vlcPackCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack files",
}

// init performs package-level initialization for command registration.
// Registers vlcPackCmd with the root command using panic-safe handling.
// Called automatically when the package is imported.
func init() {
	HandlePanic(func() {
		RootCmd.AddCommand(vlcPackCmd)
	})
}
