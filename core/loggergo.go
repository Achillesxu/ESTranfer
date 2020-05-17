package core

import (
	"fmt"
	"os"
	"time"
)

// OpenOutLog takes a single argument, which is the path to a log file. This process's stdout and stderr are redirected
// to this log file. The os.File object is returned so that it can be managed.
func OpenOutLog(filename string) *os.File {
	// Move existing out file to a dated file if it exists
	if _, err := os.Stat(filename); err == nil {
		if err = os.Rename(filename, filename+"."+time.Now().Format("2006-01-02_15:04:05")); err != nil {
			fmt.Printf("Cannot move old out file: %v", err)
			os.Exit(1)
		}
	}

	// Redirect stdout and stderr to out file
	logFile, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0644)
	internalDup2(logFile.Fd(), 1)
	internalDup2(logFile.Fd(), 2)
	return logFile
}
