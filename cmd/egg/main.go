package main

import (
	"os"

	"github.com/course-go/egg/internal/egg"
)

func main() {
	exitCode := egg.Run()
	os.Exit(exitCode)
}
