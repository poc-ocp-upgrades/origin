package util

import (
	"github.com/pkg/errors"
)

func Chroot(rootfs string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return errors.New("chroot is not implemented on Windows")
}
