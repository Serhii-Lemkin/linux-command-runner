//go:build windows
// +build windows

package helpers

import "syscall"

func getSysProcAttr() *syscall.SysProcAttr {
	return nil // or &syscall.SysProcAttr{} if needed
}
