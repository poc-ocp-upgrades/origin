package cniserver

import (
	"encoding/json"
	"fmt"
	goformat "fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	utilwait "k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog"
	"net"
	"net/http"
	"os"
	goos "os"
	"path/filepath"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

const CNIServerRunDir string = "/var/run/openshift-sdn/cniserver"
const CNIServerSocketName string = "socket"
const CNIServerSocketPath string = CNIServerRunDir + "/" + CNIServerSocketName
const CNIServerConfigFileName string = "config.json"
const CNIServerConfigFilePath string = CNIServerRunDir + "/" + CNIServerConfigFileName

type Config struct {
	MTU                uint32 `json:"mtu"`
	ServiceNetworkCIDR string `json:"serviceNetworkCIDR"`
}
type CNICommand string

const CNI_ADD CNICommand = "ADD"
const CNI_UPDATE CNICommand = "UPDATE"
const CNI_DEL CNICommand = "DEL"

type CNIRequest struct {
	Env      map[string]string `json:"env,omitempty"`
	Config   []byte            `json:"config,omitempty"`
	HostVeth string            `json:"hostVeth,omitempty"`
}
type PodRequest struct {
	Command      CNICommand
	PodNamespace string
	PodName      string
	SandboxID    string
	Netns        string
	HostVeth     string
	AssignedIP   string
	Result       chan *PodResult
}
type PodResult struct {
	Response []byte
	Err      error
}
type cniRequestFunc func(request *PodRequest) ([]byte, error)
type CNIServer struct {
	http.Server
	requestFunc cniRequestFunc
	rundir      string
	config      *Config
}

func NewCNIServer(rundir string, config *Config) *CNIServer {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	router := mux.NewRouter()
	s := &CNIServer{Server: http.Server{Handler: router}, rundir: rundir, config: config}
	router.NotFoundHandler = http.HandlerFunc(http.NotFound)
	router.HandleFunc("/", s.handleCNIRequest).Methods("POST")
	return s
}
func (s *CNIServer) Start(requestFunc cniRequestFunc) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if requestFunc == nil {
		return fmt.Errorf("no pod request handler")
	}
	s.requestFunc = requestFunc
	configPath := filepath.Join(s.rundir, CNIServerConfigFileName)
	socketPath := filepath.Join(s.rundir, CNIServerSocketName)
	info, err := os.Stat(s.rundir)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("could not read CNIServer directory: %v", err)
		}
	} else if info.IsDir() && info.Mode().Perm() == 0700 {
		if err := os.Remove(socketPath); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("failed to remove old CNIServer socket: %v", err)
		}
		if err := os.Remove(configPath); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("failed to remove old CNIServer config: %v", err)
		}
	} else {
		if err := os.RemoveAll(s.rundir); err != nil {
			return fmt.Errorf("failed to remove old CNIServer directory: %v", err)
		}
	}
	if err := os.MkdirAll(s.rundir, 0700); err != nil {
		return fmt.Errorf("failed to create CNIServer directory: %v", err)
	}
	config, err := json.Marshal(s.config)
	if err != nil {
		return fmt.Errorf("could not marshal config data: %v", err)
	}
	err = ioutil.WriteFile(configPath, config, 0444)
	if err != nil {
		return fmt.Errorf("could not write config file %q: %v", configPath, err)
	}
	l, err := net.Listen("unix", socketPath)
	if err != nil {
		return fmt.Errorf("failed to listen on pod info socket: %v", err)
	}
	if err := os.Chmod(socketPath, 0600); err != nil {
		l.Close()
		return fmt.Errorf("failed to set pod info socket mode: %v", err)
	}
	s.SetKeepAlivesEnabled(false)
	go utilwait.Forever(func() {
		if err := s.Serve(l); err != nil {
			utilruntime.HandleError(fmt.Errorf("CNI server Serve() failed: %v", err))
		}
	}, 0)
	return nil
}
func ReadConfig(configPath string) (*Config, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("OpenShift SDN network process is not (yet?) available")
		} else {
			return nil, fmt.Errorf("could not read config file %q: %v", configPath, err)
		}
	}
	var config Config
	if err = json.Unmarshal(bytes, &config); err != nil {
		return nil, fmt.Errorf("could not parse config file %q: %v", configPath, err)
	}
	return &config, nil
}
func gatherCNIArgs(env map[string]string) (map[string]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cniArgs, ok := env["CNI_ARGS"]
	if !ok {
		return nil, fmt.Errorf("missing CNI_ARGS: '%s'", env)
	}
	mapArgs := make(map[string]string)
	for _, arg := range strings.Split(cniArgs, ";") {
		parts := strings.Split(arg, "=")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid CNI_ARG '%s'", arg)
		}
		mapArgs[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}
	return mapArgs, nil
}
func cniRequestToPodRequest(r *http.Request) (*PodRequest, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var cr CNIRequest
	b, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(b, &cr); err != nil {
		return nil, fmt.Errorf("JSON unmarshal error: %v", err)
	}
	cmd, ok := cr.Env["CNI_COMMAND"]
	if !ok {
		return nil, fmt.Errorf("unexpected or missing CNI_COMMAND")
	}
	req := &PodRequest{Command: CNICommand(cmd), Result: make(chan *PodResult)}
	req.SandboxID, ok = cr.Env["CNI_CONTAINERID"]
	if !ok {
		return nil, fmt.Errorf("missing CNI_CONTAINERID")
	}
	req.Netns, ok = cr.Env["CNI_NETNS"]
	if !ok {
		return nil, fmt.Errorf("missing CNI_NETNS")
	}
	req.HostVeth = cr.HostVeth
	if req.HostVeth == "" && req.Command == CNI_ADD {
		return nil, fmt.Errorf("missing HostVeth")
	}
	cniArgs, err := gatherCNIArgs(cr.Env)
	if err != nil {
		return nil, err
	}
	req.PodNamespace, ok = cniArgs["K8S_POD_NAMESPACE"]
	if err != nil {
		return nil, fmt.Errorf("missing K8S_POD_NAMESPACE")
	}
	req.PodName, ok = cniArgs["K8S_POD_NAME"]
	if err != nil {
		return nil, fmt.Errorf("missing K8S_POD_NAME")
	}
	return req, nil
}
func (s *CNIServer) handleCNIRequest(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	req, err := cniRequestToPodRequest(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}
	klog.V(5).Infof("Waiting for %s result for pod %s/%s", req.Command, req.PodNamespace, req.PodName)
	result, err := s.requestFunc(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(result); err != nil {
			klog.Warningf("Error writing %s HTTP response: %v", req.Command, err)
		}
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
