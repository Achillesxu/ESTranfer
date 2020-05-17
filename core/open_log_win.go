// +build windows

package core

import (
	"syscall"
)

var (
	kernel32         = syscall.MustLoadDLL("kernel32.dll")
	procSetStdHandle = kernel32.MustFindProc("SetStdHandle")
)

func internalDup2(from, to uintptr) error {
	r0, _, e1 := syscall.Syscall(procSetStdHandle.Addr(), 2, from, to, 0)
	if r0 == 0 {
		if e1 != 0 {
			return error(e1)
		}
		return syscall.EINVAL
	}
	return nil
}
