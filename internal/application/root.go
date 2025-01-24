package application

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "simpleArchiver",
	Short: "A simple archiver program",
}

func Execute() {
	HandleError(RootCmd.Execute)
}
