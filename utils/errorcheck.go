package utils

import (
	"os"

	"github.com/charmbracelet/log"
)

// ErrorCheck ... utility function to help stream line error checking
func ErrorCheck(e error) {
	if e != nil {
		ErrorOut(e)
	}
}

// ErrorOut .. function to hold the actual erroring out
func ErrorOut(err interface{}) {
	log.Error(err)
	os.Exit(1)
}

// CustomErrorOut ... function to error using a custom string and exit code
func CustomErrorOut(customErr string, customExitCode int) {
	log.Error(customErr)
	os.Exit(customExitCode)
}
