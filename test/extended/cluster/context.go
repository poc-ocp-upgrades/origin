package cluster

import (
	"path/filepath"
	"strings"
	"time"
	"github.com/spf13/viper"
)

const (
	IF_EXISTS_DELETE	= "delete"
	IF_EXISTS_REUSE		= "reuse"
)

type ContextType struct {
	ClusterLoader struct {
		Cleanup		bool
		Projects	[]ClusterLoaderType
		Sync		SyncObjectType	`yaml:",omitempty"`
		TuningSets	[]TuningSetType	`yaml:",omitempty"`
	}
}
type ClusterLoaderType struct {
	Number		int	`mapstructure:"num" yaml:"num"`
	Basename	string
	IfExists	string				`json:"ifexists"`
	Labels		map[string]string		`yaml:",omitempty"`
	NodeSelector	string				`yaml:",omitempty"`
	Tuning		string				`yaml:",omitempty"`
	Configmaps	map[string]interface{}		`yaml:",omitempty"`
	Secrets		map[string]interface{}		`yaml:",omitempty"`
	Pods		[]ClusterLoaderObjectType	`yaml:",omitempty"`
	Templates	[]ClusterLoaderObjectType	`yaml:",omitempty"`
}
type ClusterLoaderObjectType struct {
	Total		int			`yaml:",omitempty"`
	Number		int			`mapstructure:"num" yaml:"num"`
	Image		string			`yaml:",omitempty"`
	Basename	string			`yaml:",omitempty"`
	File		string			`yaml:",omitempty"`
	Sync		SyncObjectType		`yaml:",omitempty"`
	Parameters	map[string]interface{}	`yaml:",omitempty"`
}
type SyncObjectType struct {
	Server	struct {
		Enabled	bool
		Port	int
	}
	Running		bool
	Succeeded	bool
	Selectors	map[string]string
	Timeout		string
}
type TuningSetType struct {
	Name		string
	Pods		TuningSetObjectType
	Templates	TuningSetObjectType
}
type TuningSetObjectType struct {
	Stepping	struct {
		StepSize	int
		Pause		time.Duration
		Timeout		time.Duration
	}
	RateLimit	struct{ Delay time.Duration }
}

var ConfigContext ContextType

type PodCount struct {
	Started		int
	Stopped		int
	Shutdown	chan bool
}
type ServiceInfo struct {
	Name	string
	IP	string
	Port	int32
}

func ParseConfig(config string, isFixture bool) error {
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
	if isFixture {
		dir, file := filepath.Split(config)
		s := strings.Split(file, ".")
		viper.SetConfigName(s[0])
		viper.AddConfigPath(dir)
	} else {
		viper.SetConfigName(config)
		viper.AddConfigPath(".")
	}
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	viper.Unmarshal(&ConfigContext)
	return nil
}
