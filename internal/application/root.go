package application

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Short: "Simple Archiver",
}

func Execute() {
	HandleError(RootCmd.Execute)
}
