// +build solaris

package remote

import "github.com/golang/sys/unix"

func syscallDup(oldfd int, newfd int) (err error) {
	return unix.Dup2(oldfd, newfd)
}
