package userspace

import "golang.org/x/sys/unix"

func setRLimit(limit uint64) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &unix.Rlimit{Max: limit, Cur: limit})
}
