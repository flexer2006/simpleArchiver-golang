package main

import "github.com/flexer2006/simpleArchiver-golang/internal/application"

func main() {
	application.HandlePanic(func() {
		application.Execute()
	})
}
