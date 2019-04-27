package rsync

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	krand "k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/klog"
)

const (
	startDaemonScript	= `set -e
TMPDIR=${TMPDIR:-/tmp}
CONFIGFILE=$(echo -n "" > ${TMPDIR}/%[1]s && echo ${TMPDIR}/%[1]s)
PIDFILE=$(echo -n "" > ${TMPDIR}/%[2]s && echo ${TMPDIR}/%[2]s)
rm $PIDFILE
printf "pid file = ${PIDFILE}\n[root]\n  path = /\n  use chroot = no\n  read only = no" > $CONFIGFILE
rsync --no-detach --daemon --config=${CONFIGFILE} --port=%[3]d
`
	killDaemonScript	= `set -e
TMPDIR=${TMPDIR:-/tmp}
PIDFILE=${TMPDIR}/%[1]s
kill $(cat ${PIDFILE})
`
	checkDaemonScript	= `set -e
TMPDIR=${TMPDIR:-/tmp}
PIDFILE=${TMPDIR}/%[1]s
ls ${PIDFILE}
`
	portRangeFrom		= 30000
	portRangeTo		= 60000
	remoteLabel		= "root"
	RsyncDaemonStartTimeOut	= 10 * time.Second
)

var (
	random = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
)

type rsyncDaemonStrategy struct {
	Flags		[]string
	RemoteExecutor	executor
	PortForwarder	forwarder
	LocalExecutor	executor
	daemonPIDFile	string
	daemonPort	int
	localPort	int
	portForwardChan	chan struct{}
}

func localRsyncURL(port int, label string, path string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("rsync://127.0.0.1:%d/%s/%s", port, label, strings.TrimPrefix(path, "/"))
}
func getUsedPorts(data string) map[int]struct{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ports := map[int]struct{}{}
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		if !strings.Contains(parts[0], ":") {
			continue
		}
		klog.V(5).Infof("Determining port in use from: %s", line)
		localAddress := strings.Split(parts[1], ":")
		if len(localAddress) < 2 {
			continue
		}
		port, err := strconv.ParseInt(localAddress[1], 16, 0)
		if err == nil {
			ports[int(port)] = struct{}{}
		}
	}
	klog.V(2).Infof("Used ports in container: %#v", ports)
	return ports
}
func randomPort() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return portRangeFrom + random.Intn(portRangeTo-portRangeFrom)
}
func localPort() (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		klog.V(1).Infof("Could not determine local free port: %v", err)
		return 0, err
	}
	defer l.Close()
	klog.V(1).Infof("Found listener port at: %s", l.Addr().String())
	addr := strings.Split(l.Addr().String(), ":")
	port, err := strconv.Atoi(addr[len(addr)-1])
	if err != nil {
		klog.V(1).Infof("Could not parse listener address %#v: %v", addr, err)
		return 0, err
	}
	return port, nil
}
func (s *rsyncDaemonStrategy) getFreePort() (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := []string{"bash", "-c", "cat /proc/net/tcp && ( [ -e /proc/net/tcp6 ] && cat /proc/net/tcp6 || true)"}
	tcpData := &bytes.Buffer{}
	cmdErr := &bytes.Buffer{}
	usedPorts := map[int]struct{}{}
	err := s.RemoteExecutor.Execute(cmd, nil, tcpData, cmdErr)
	if err == nil {
		usedPorts = getUsedPorts(tcpData.String())
	} else {
		klog.V(4).Infof("Error getting free port data: %v, Err: %s", err, cmdErr.String())
	}
	tries := 0
	for {
		tries++
		if tries > 20 {
			klog.V(4).Infof("Too many attempts trying to find free port")
			break
		}
		port := randomPort()
		if _, used := usedPorts[port]; !used {
			klog.V(4).Infof("Found free container port: %d", port)
			return port, nil
		}
	}
	return 0, fmt.Errorf("could not find a free port")
}
func (s *rsyncDaemonStrategy) startRemoteDaemon() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	port, err := s.getFreePort()
	if err != nil {
		return err
	}
	cmdOut := &bytes.Buffer{}
	cmdErr := &bytes.Buffer{}
	pidFile := krand.String(32)
	configFile := krand.String(32)
	cmdIn := bytes.NewBufferString(fmt.Sprintf(startDaemonScript, configFile, pidFile, port))
	daemonErr := make(chan error, 1)
	go func() {
		err = s.RemoteExecutor.Execute([]string{"sh"}, cmdIn, cmdOut, cmdErr)
		if err != nil {
			daemonErr <- fmt.Errorf("error starting rsync daemon: %v\nOut: %s\nErr: %s\n", err, cmdOut.String(), cmdErr.String())
		}
	}()
	checkScript := bytes.NewBufferString(fmt.Sprintf(checkDaemonScript, pidFile))
	startTime := time.Now()
	for {
		if time.Since(startTime) > RsyncDaemonStartTimeOut {
			return fmt.Errorf("timed out waiting for rsync daemon to start")
		}
		checkScript.Reset()
		err = s.RemoteExecutor.Execute([]string{"sh"}, checkScript, ioutil.Discard, ioutil.Discard)
		if err == nil {
			break
		}
		if len(daemonErr) > 0 {
			return <-daemonErr
		}
		time.Sleep(100 * time.Millisecond)
	}
	s.daemonPort = port
	s.daemonPIDFile = pidFile
	return nil
}
func (s *rsyncDaemonStrategy) killRemoteDaemon() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := []string{"sh"}
	cmdIn := bytes.NewBufferString(fmt.Sprintf(killDaemonScript, s.daemonPIDFile))
	cmdOut := &bytes.Buffer{}
	cmdErr := &bytes.Buffer{}
	err := s.RemoteExecutor.Execute(cmd, cmdIn, cmdOut, cmdErr)
	if err != nil {
		klog.V(1).Infof("Error killing rsync daemon: %v. Out: %s, Err: %s\n", err, cmdOut.String(), cmdErr.String())
	}
	return err
}
func (s *rsyncDaemonStrategy) startPortForward() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	s.localPort, err = localPort()
	if err != nil {
		s.localPort = randomPort()
	}
	s.portForwardChan = make(chan struct{})
	return s.PortForwarder.ForwardPorts([]string{fmt.Sprintf("%d:%d", s.localPort, s.daemonPort)}, s.portForwardChan)
}
func (s *rsyncDaemonStrategy) stopPortForward() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	close(s.portForwardChan)
}
func (s *rsyncDaemonStrategy) copyUsingDaemon(source, destination *PathSpec, out, errOut io.Writer) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(3).Infof("Copying files with rsync daemon")
	cmd := append([]string{"rsync"}, s.Flags...)
	var sourceArg, destinationArg string
	if source.Local() {
		sourceArg = source.RsyncPath()
	} else {
		sourceArg = localRsyncURL(s.localPort, remoteLabel, source.Path)
	}
	if destination.Local() {
		destinationArg = destination.RsyncPath()
	} else {
		destinationArg = localRsyncURL(s.localPort, remoteLabel, destination.Path)
	}
	cmd = append(cmd, sourceArg, destinationArg)
	err := s.LocalExecutor.Execute(cmd, nil, out, errOut)
	if err != nil {
		testRsyncErr := executeWithLogging(s.RemoteExecutor, testRsyncCommand)
		if testRsyncErr != nil {
			return strategySetupError("rsync not available in container")
		}
	}
	return err
}
func (s *rsyncDaemonStrategy) Copy(source, destination *PathSpec, out, errOut io.Writer) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := s.startRemoteDaemon()
	if err != nil {
		if isExitError(err) {
			return strategySetupError(fmt.Sprintf("cannot start remote rsync daemon: %v", err))
		}
		return err
	}
	defer s.killRemoteDaemon()
	err = s.startPortForward()
	if err != nil {
		if isExitError(err) {
			return strategySetupError(fmt.Sprintf("cannot start port-forward: %v", err))
		}
		return err
	}
	defer s.stopPortForward()
	err = s.copyUsingDaemon(source, destination, out, errOut)
	return err
}
func (s *rsyncDaemonStrategy) Validate() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	errs := []error{}
	if s.PortForwarder == nil {
		errs = append(errs, errors.New("port forwarder must be provided"))
	}
	if s.LocalExecutor == nil {
		errs = append(errs, errors.New("local executor must be provided"))
	}
	if s.RemoteExecutor == nil {
		errs = append(errs, errors.New("remote executor must be provided"))
	}
	if len(errs) > 0 {
		return kerrors.NewAggregate(errs)
	}
	return nil
}
func NewRsyncDaemonStrategy(o *RsyncOptions) CopyStrategy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	flags := rsyncDefaultFlags
	flags = append(flags, rsyncFlagsFromOptions(o)...)
	remoteExec := newRemoteExecutor(o)
	forwarder := newPortForwarder(o)
	return &rsyncDaemonStrategy{Flags: flags, RemoteExecutor: remoteExec, LocalExecutor: newLocalExecutor(), PortForwarder: forwarder}
}
func (s *rsyncDaemonStrategy) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "rsync-daemon"
}
