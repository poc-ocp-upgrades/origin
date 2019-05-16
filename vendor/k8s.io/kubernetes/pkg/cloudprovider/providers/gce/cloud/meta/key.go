package meta

import (
	"fmt"
	goformat "fmt"
	goos "os"
	"regexp"
	godefaultruntime "runtime"
	gotime "time"
)

type Key struct {
	Name   string
	Zone   string
	Region string
}
type KeyType string

const (
	Zonal    = "zonal"
	Regional = "regional"
	Global   = "global"
)

var (
	locationRegexp = regexp.MustCompile("^[a-z](?:[-a-z0-9]+)?$")
)

func ZonalKey(name, zone string) *Key {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &Key{name, zone, ""}
}
func RegionalKey(name, region string) *Key {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &Key{name, "", region}
}
func GlobalKey(name string) *Key {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &Key{name, "", ""}
}
func (k *Key) Type() KeyType {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch {
	case k.Zone != "":
		return Zonal
	case k.Region != "":
		return Regional
	default:
		return Global
	}
}
func (k Key) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch k.Type() {
	case Zonal:
		return fmt.Sprintf("Key{%q, zone: %q}", k.Name, k.Zone)
	case Regional:
		return fmt.Sprintf("Key{%q, region: %q}", k.Name, k.Region)
	default:
		return fmt.Sprintf("Key{%q}", k.Name)
	}
}
func (k *Key) Valid() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if k.Zone != "" && k.Region != "" {
		return false
	}
	switch {
	case k.Region != "":
		return locationRegexp.Match([]byte(k.Region))
	case k.Zone != "":
		return locationRegexp.Match([]byte(k.Zone))
	}
	return true
}
func KeysToMap(keys ...Key) map[Key]bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ret := map[Key]bool{}
	for _, k := range keys {
		ret[k] = true
	}
	return ret
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
