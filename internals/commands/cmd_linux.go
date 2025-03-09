//go:build linux
package commands

import (
	"syscall"
)

func WithNewNamespace() CmdOptions {
	return func(c *Command) {
		c.SysProcAttr = &syscall.SysProcAttr{
			Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
		}
	}
}