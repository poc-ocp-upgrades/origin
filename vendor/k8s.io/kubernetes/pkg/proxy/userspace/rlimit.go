package userspace

import "golang.org/x/sys/unix"

func setRLimit(limit uint64) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return unix.Setrlimit(unix.RLIMIT_NOFILE, &unix.Rlimit{Max: limit, Cur: limit})
}
