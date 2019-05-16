package verflag

import (
	"fmt"
	goformat "fmt"
	flag "github.com/spf13/pflag"
	"k8s.io/kubernetes/pkg/version"
	"os"
	goos "os"
	godefaultruntime "runtime"
	"strconv"
	gotime "time"
)

type versionValue int

const (
	VersionFalse versionValue = 0
	VersionTrue  versionValue = 1
	VersionRaw   versionValue = 2
)
const strRawVersion string = "raw"

func (v *versionValue) IsBoolFlag() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func (v *versionValue) Get() interface{} {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return versionValue(*v)
}
func (v *versionValue) Set(s string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if s == strRawVersion {
		*v = VersionRaw
		return nil
	}
	boolVal, err := strconv.ParseBool(s)
	if boolVal {
		*v = VersionTrue
	} else {
		*v = VersionFalse
	}
	return err
}
func (v *versionValue) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if *v == VersionRaw {
		return strRawVersion
	}
	return fmt.Sprintf("%v", bool(*v == VersionTrue))
}
func (v *versionValue) Type() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "version"
}
func VersionVar(p *versionValue, name string, value versionValue, usage string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	*p = value
	flag.Var(p, name, usage)
	flag.Lookup(name).NoOptDefVal = "true"
}
func Version(name string, value versionValue, usage string) *versionValue {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	p := new(versionValue)
	VersionVar(p, name, value, usage)
	return p
}

const versionFlagName = "version"

var (
	versionFlag = Version(versionFlagName, VersionFalse, "Print version information and quit")
)

func AddFlags(fs *flag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fs.AddFlag(flag.Lookup(versionFlagName))
}
func PrintAndExitIfRequested() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if *versionFlag == VersionRaw {
		fmt.Printf("%#v\n", version.Get())
		os.Exit(0)
	} else if *versionFlag == VersionTrue {
		fmt.Printf("Kubernetes %s\n", version.Get())
		os.Exit(0)
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
