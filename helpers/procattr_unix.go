//go:build !windows
// +build !windows

package helpers

import "syscall"

func getSysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{Setpgid: true}
}
