package main

import (
	"github.com/flexer2006/simpleArchiver-golang/internal/application"
	"github.com/flexer2006/simpleArchiver-golang/pkg/vlc"
)

func main() {

	application.HandlePanic(func() {

		application.RootCmd.AddCommand(vlc.Vlc)
		application.Execute()
	})
}
