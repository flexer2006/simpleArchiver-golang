// Package application provides CLI command configuration and error handling
// for archive operations. This file implements the 'unpack' command setup.
package application

import (
	"errors"
	"github.com/spf13/cobra"
	"log"
)

// vlcUnpackCmd represents the 'unpack' subcommand for file extraction operations.
// Registered automatically during package initialization to the root command.
//
// Usage: unpack [flags]
// Short: Unpack files using VLC decoding
var vlcUnpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack files",
}

// init performs package-level initialization for the unpack command:
// 1. Wraps registration logic in panic recovery using HandlePanic
// 2. Validates RootCommand availability with nil-check
// 3. Adds unpack command to root command hierarchy
// 4. Logs fatal error and terminates if registration fails
//
// Automatically called when package is imported. Uses layered error handling:
// - HandlePanic for runtime panic recovery
// - HandleError for error propagation
// - log.Fatalf for critical failure reporting
func init() {
	HandlePanic(func() {
		err := HandleError(func() error {
			if RootCmd == nil {
				return errors.New("RootCmd is nil")
			}
			RootCmd.AddCommand(vlcUnpackCmd)
			return nil
		})

		if err != nil {
			log.Fatalf("Error occurred while adding unpack command: %v", err)
		}
	})
}
