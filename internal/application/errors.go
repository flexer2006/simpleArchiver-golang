package application

import (
	"log"
	"os"
	"runtime"
)

func HandlePanic(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			logPanic(r)
		}
	}()
	fn()
}

func HandleError(fn func() error) {
	if err := fn(); err != nil {
		logError(err)
		os.Exit(1)
	}
}

func logPanic(recovered interface{}) {
	buf := make([]byte, 1024)
	runtime.Stack(buf, false)
	log.Printf("PANIC: %v\nSTACK TRACE:\n%s", recovered, buf)
}

func logError(err error) {
	log.Printf("ERROR: %v", err)
}
