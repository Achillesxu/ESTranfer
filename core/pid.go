package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"syscall"
)

// CheckAndCreatePidFile takes a single argument, which is the path to a PID file (a file that contains a single
// integer, which is the process ID of a running process). If this file exists, and if the PID is that of a running
// process, return false as that indicates another copy of this process is already running. Otherwise, create the
// file and write this process's PID to the file and return true. Any error doing this (such as not having permissions
// to write the file) will return false.
//
// This func should be called when ESTransfer starts to prevent multiple copies from running.
func CheckAndCreatePidFile(filename string) bool {
	// Check if the PID file exists
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		// The file exists, so read it and check if the PID specified is running
		pidString, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Printf("Cannot read PID file: %v", err)
			return false
		}
		pid, err := strconv.Atoi(string(pidString))
		if err != nil {
			fmt.Printf("Cannot interpret contents of PID file: %v", err)
			return false
		}

		if pid == os.Getpid() {
			// This could happen inside a docker
			// container, e.g. the pid of ESTransfer could be
			// equal to 1 each time the container is
			// restarted.
			fmt.Println("Found existing pidfile matching current pid")
			return true
		}

		// Try sending a signal to the process to see if it is still running
		process, err := os.FindProcess(pid)
		if err == nil {
			err = process.Signal(syscall.Signal(0))
			if (err == nil) || (err == syscall.EPERM) {
				// The process exists, so we're going to assume it's an old ESTransfer and we shouldn't start
				fmt.Printf("Existing process running on PID %d. Exiting (my pid = %d)", pid, os.Getpid())
				return false
			}
		}
	}

	fmt.Println("write PID file")
	// Create a PID file, replacing any existing one (as we already checked it)
	pidFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Cannot write PID file: %v", err)
		return false
	}
	_, _ = fmt.Fprintf(pidFile, "%v", os.Getpid())
	_ = pidFile.Close()
	return true
}

// RemovePidFile takes a single argument, which is the path to a PID file. That file is deleted. This func should be
// called when ESTransfer exits.
func RemovePidFile(filename string) {
	err := os.Remove(filename)
	if err != nil {
		fmt.Printf("Failed to remove PID file: %v\n", err)
	}
}
