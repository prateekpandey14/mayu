package main

import "syscall"

func genPlatformSysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{Pdeathsig: 9}
}
