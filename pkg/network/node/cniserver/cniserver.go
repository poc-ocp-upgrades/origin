package cniserver

import (
	"encoding/json"
	"bytes"
	"runtime"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"github.com/gorilla/mux"
	"k8s.io/klog"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	utilwait "k8s.io/apimachinery/pkg/util/wait"
)

const CNIServerRunDir string = "/var/run/openshift-sdn"
const CNIServerSocketName string = "cni-server.sock"
const CNIServerSocketPath string = CNIServerRunDir + "/" + CNIServerSocketName
const CNIServerConfigFileName string = "config.json"
const CNIServerConfigFilePath string = CNIServerRunDir + "/" + CNIServerConfigFileName

type Config struct {
	MTU			uint32	`json:"mtu"`
	ServiceNetworkCIDR	string	`json:"serviceNetworkCIDR"`
}
type CNICommand string

const CNI_ADD CNICommand = "ADD"
const CNI_UPDATE CNICommand = "UPDATE"
const CNI_DEL CNICommand = "DEL"

type CNIRequest struct {
	Env		map[string]string	`json:"env,omitempty"`
	Config		[]byte			`json:"config,omitempty"`
	HostVeth	string			`json:"hostVeth,omitempty"`
}
type PodRequest struct {
	Command		CNICommand
	PodNamespace	string
	PodName		string
	SandboxID	string
	Netns		string
	HostVeth	string
	AssignedIP	string
	Result		chan *PodResult
}
type PodResult struct {
	Response	[]byte
	Err		error
}
type cniRequestFunc func(request *PodRequest) ([]byte, error)
type CNIServer struct {
	http.Server
	requestFunc	cniRequestFunc
	rundir		string
	config		*Config
}

func NewCNIServer(rundir string, config *Config) *CNIServer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	router := mux.NewRouter()
	s := &CNIServer{Server: http.Server{Handler: router}, rundir: rundir, config: config}
	router.NotFoundHandler = http.HandlerFunc(http.NotFound)
	router.HandleFunc("/", s.handleCNIRequest).Methods("POST")
	return s
}
func (s *CNIServer) Start(requestFunc cniRequestFunc) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if requestFunc == nil {
		return fmt.Errorf("no pod request handler")
	}
	s.requestFunc = requestFunc
	if err := os.RemoveAll(s.rundir); err != nil && !os.IsNotExist(err) {
		utilruntime.HandleError(fmt.Errorf("failed to remove old pod info socket: %v", err))
	}
	if err := os.RemoveAll(s.rundir); err != nil && !os.IsNotExist(err) {
		utilruntime.HandleError(fmt.Errorf("failed to remove contents of socket directory: %v", err))
	}
	if err := os.MkdirAll(s.rundir, 0700); err != nil {
		return fmt.Errorf("failed to create pod info socket directory: %v", err)
	}
	config, err := json.Marshal(s.config)
	if err != nil {
		return fmt.Errorf("could not marshal config data: %v", err)
	}
	configPath := filepath.Join(s.rundir, CNIServerConfigFileName)
	err = ioutil.WriteFile(configPath, config, os.FileMode(0444))
	if err != nil {
		return fmt.Errorf("could not write config file %q: %v", configPath, err)
	}
	socketPath := filepath.Join(s.rundir, CNIServerSocketName)
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
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
