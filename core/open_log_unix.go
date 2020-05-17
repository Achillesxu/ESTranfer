// +build !windows
// +build !arm64

package core

import (
	"syscall"
)

func internalDup2(from, to uintptr) error {
	return syscall.Dup2(int(from), int(to))
}
