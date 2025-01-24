package application

import "github.com/spf13/cobra"

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack files",
}

func init() {
	HandlePanic(func() {
		RootCmd.AddCommand(packCmd)
	})
}
