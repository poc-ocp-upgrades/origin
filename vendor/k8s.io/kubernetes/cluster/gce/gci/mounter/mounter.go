package main

import (
	"fmt"
	goformat "fmt"
	"os"
	goos "os"
	"os/exec"
	"path/filepath"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

const (
	chrootCmd        = "chroot"
	mountCmd         = "mount"
	rootfs           = "rootfs"
	nfsRPCBindErrMsg = "mount.nfs: rpc.statd is not running but is required for remote locking.\nmount.nfs: Either use '-o nolock' to keep locks local, or start statd.\nmount.nfs: an incorrect mount option was specified\n"
	rpcBindCmd       = "/sbin/rpcbind"
	defaultRootfs    = "/home/kubernetes/containerized_mounter/rootfs"
)

func main() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Command failed: must provide a command to run.\n")
		return
	}
	path, _ := filepath.Split(os.Args[0])
	rootfsPath := filepath.Join(path, rootfs)
	if _, err := os.Stat(rootfsPath); os.IsNotExist(err) {
		rootfsPath = defaultRootfs
	}
	command := os.Args[1]
	switch command {
	case mountCmd:
		mountErr := mountInChroot(rootfsPath, os.Args[2:])
		if mountErr != nil {
			fmt.Fprintf(os.Stderr, "Mount failed: %v", mountErr)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "Unknown command, must be %s", mountCmd)
		os.Exit(1)
	}
}
func mountInChroot(rootfsPath string, args []string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if _, err := os.Stat(rootfsPath); os.IsNotExist(err) {
		return fmt.Errorf("path <%s> does not exist", rootfsPath)
	}
	args = append([]string{rootfsPath, mountCmd}, args...)
	output, err := exec.Command(chrootCmd, args...).CombinedOutput()
	if err == nil {
		return nil
	}
	if !strings.EqualFold(string(output), nfsRPCBindErrMsg) {
		return fmt.Errorf("mount failed: %v\nMounting command: %s\nMounting arguments: %v\nOutput: %s", err, chrootCmd, args, string(output))
	}
	output, err = exec.Command(chrootCmd, rootfsPath, rpcBindCmd, "-w").CombinedOutput()
	if err != nil {
		return fmt.Errorf("Mount issued for NFS V3 but unable to run rpcbind:\n Output: %s\n Error: %v", string(output), err)
	}
	output, err = exec.Command(chrootCmd, args...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("Mount failed for NFS V3 even after running rpcBind %s, %v", string(output), err)
	}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
