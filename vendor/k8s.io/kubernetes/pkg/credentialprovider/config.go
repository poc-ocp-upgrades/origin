package credentialprovider

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	goformat "fmt"
	"io/ioutil"
	"k8s.io/klog"
	"net/http"
	"os"
	goos "os"
	"path/filepath"
	godefaultruntime "runtime"
	"strings"
	"sync"
	gotime "time"
)

type DockerConfigJson struct {
	Auths       DockerConfig      `json:"auths"`
	HttpHeaders map[string]string `json:"HttpHeaders,omitempty"`
}
type DockerConfig map[string]DockerConfigEntry
type DockerConfigEntry struct {
	Username string
	Password string
	Email    string
	Provider DockerConfigProvider
}

var (
	preferredPathLock  sync.Mutex
	preferredPath      = ""
	workingDirPath     = ""
	homeDirPath        = os.Getenv("HOME")
	rootDirPath        = "/"
	homeJsonDirPath    = filepath.Join(homeDirPath, ".docker")
	rootJsonDirPath    = filepath.Join(rootDirPath, ".docker")
	configFileName     = ".dockercfg"
	configJsonFileName = "config.json"
)

func SetPreferredDockercfgPath(path string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	preferredPathLock.Lock()
	defer preferredPathLock.Unlock()
	preferredPath = path
}
func GetPreferredDockercfgPath() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	preferredPathLock.Lock()
	defer preferredPathLock.Unlock()
	return preferredPath
}
func DefaultDockercfgPaths() []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return []string{GetPreferredDockercfgPath(), workingDirPath, homeDirPath, rootDirPath}
}
func DefaultDockerConfigJSONPaths() []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return []string{GetPreferredDockercfgPath(), workingDirPath, homeJsonDirPath, rootJsonDirPath}
}
func ReadDockercfgFile(searchPaths []string) (cfg DockerConfig, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(searchPaths) == 0 {
		searchPaths = DefaultDockercfgPaths()
	}
	for _, configPath := range searchPaths {
		absDockerConfigFileLocation, err := filepath.Abs(filepath.Join(configPath, configFileName))
		if err != nil {
			klog.Errorf("while trying to canonicalize %s: %v", configPath, err)
			continue
		}
		klog.V(4).Infof("looking for .dockercfg at %s", absDockerConfigFileLocation)
		contents, err := ioutil.ReadFile(absDockerConfigFileLocation)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			klog.V(4).Infof("while trying to read %s: %v", absDockerConfigFileLocation, err)
			continue
		}
		cfg, err := readDockerConfigFileFromBytes(contents)
		if err == nil {
			klog.V(4).Infof("found .dockercfg at %s", absDockerConfigFileLocation)
			return cfg, nil
		}
	}
	return nil, fmt.Errorf("couldn't find valid .dockercfg after checking in %v", searchPaths)
}
func ReadDockerConfigJSONFile(searchPaths []string) (cfg DockerConfig, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(searchPaths) == 0 {
		searchPaths = DefaultDockerConfigJSONPaths()
	}
	for _, configPath := range searchPaths {
		absDockerConfigFileLocation, err := filepath.Abs(filepath.Join(configPath, configJsonFileName))
		if err != nil {
			klog.Errorf("while trying to canonicalize %s: %v", configPath, err)
			continue
		}
		klog.V(4).Infof("looking for %s at %s", configJsonFileName, absDockerConfigFileLocation)
		cfg, err = ReadSpecificDockerConfigJsonFile(absDockerConfigFileLocation)
		if err != nil {
			if !os.IsNotExist(err) {
				klog.V(4).Infof("while trying to read %s: %v", absDockerConfigFileLocation, err)
			}
			continue
		}
		klog.V(4).Infof("found valid %s at %s", configJsonFileName, absDockerConfigFileLocation)
		return cfg, nil
	}
	return nil, fmt.Errorf("couldn't find valid %s after checking in %v", configJsonFileName, searchPaths)
}
func ReadSpecificDockerConfigJsonFile(filePath string) (cfg DockerConfig, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var contents []byte
	if contents, err = ioutil.ReadFile(filePath); err != nil {
		return nil, err
	}
	return readDockerConfigJsonFileFromBytes(contents)
}
func ReadDockerConfigFile() (cfg DockerConfig, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if cfg, err := ReadDockerConfigJSONFile(nil); err == nil {
		return cfg, nil
	}
	return ReadDockercfgFile(nil)
}

type HttpError struct {
	StatusCode int
	Url        string
}

func (he *HttpError) Error() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("http status code: %d while fetching url %s", he.StatusCode, he.Url)
}
func ReadUrl(url string, client *http.Client, header *http.Header) (body []byte, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if header != nil {
		req.Header = *header
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		klog.V(2).Infof("body of failing http response: %v", resp.Body)
		return nil, &HttpError{StatusCode: resp.StatusCode, Url: url}
	}
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return contents, nil
}
func ReadDockerConfigFileFromUrl(url string, client *http.Client, header *http.Header) (cfg DockerConfig, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if contents, err := ReadUrl(url, client, header); err != nil {
		return nil, err
	} else {
		return readDockerConfigFileFromBytes(contents)
	}
}
func readDockerConfigFileFromBytes(contents []byte) (cfg DockerConfig, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err = json.Unmarshal(contents, &cfg); err != nil {
		klog.Errorf("while trying to parse blob %q: %v", contents, err)
		return nil, err
	}
	return
}
func readDockerConfigJsonFileFromBytes(contents []byte) (cfg DockerConfig, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var cfgJson DockerConfigJson
	if err = json.Unmarshal(contents, &cfgJson); err != nil {
		klog.Errorf("while trying to parse blob %q: %v", contents, err)
		return nil, err
	}
	cfg = cfgJson.Auths
	return
}

type dockerConfigEntryWithAuth struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	Auth     string `json:"auth,omitempty"`
}

func (ident *DockerConfigEntry) UnmarshalJSON(data []byte) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var tmp dockerConfigEntryWithAuth
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	ident.Username = tmp.Username
	ident.Password = tmp.Password
	ident.Email = tmp.Email
	if len(tmp.Auth) == 0 {
		return nil
	}
	ident.Username, ident.Password, err = decodeDockerConfigFieldAuth(tmp.Auth)
	return err
}
func (ident DockerConfigEntry) MarshalJSON() ([]byte, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	toEncode := dockerConfigEntryWithAuth{ident.Username, ident.Password, ident.Email, ""}
	toEncode.Auth = encodeDockerConfigFieldAuth(ident.Username, ident.Password)
	return json.Marshal(toEncode)
}
func decodeDockerConfigFieldAuth(field string) (username, password string, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	decoded, err := base64.StdEncoding.DecodeString(field)
	if err != nil {
		return
	}
	parts := strings.SplitN(string(decoded), ":", 2)
	if len(parts) != 2 {
		err = fmt.Errorf("unable to parse auth field")
		return
	}
	username = parts[0]
	password = parts[1]
	return
}
func encodeDockerConfigFieldAuth(username, password string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fieldValue := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(fieldValue))
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
