// +build !windows
// +build !arm64

package core

import (
	"syscall"
)

// linux_arm64 doesn't have syscall.Dup2, so use
// the nearly identical syscall.Dup3 instead
func internalDup2(from, to uintptr) error {
	return syscall.Dup3(int(oldfd), int(newfd), 0)
}
