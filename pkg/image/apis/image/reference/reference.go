package reference

import (
	goformat "fmt"
	"github.com/openshift/origin/pkg/image/internal/digest"
	"github.com/openshift/origin/pkg/image/internal/reference"
	"net"
	"net/url"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

type DockerImageReference struct {
	Registry  string
	Namespace string
	Name      string
	Tag       string
	ID        string
}

const (
	DockerDefaultRegistry   = "docker.io"
	DockerDefaultV1Registry = "index." + DockerDefaultRegistry
	DockerDefaultV2Registry = "registry-1." + DockerDefaultRegistry
)

func Parse(spec string) (DockerImageReference, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var ref DockerImageReference
	namedRef, err := reference.ParseNamed(spec)
	if err != nil {
		return ref, err
	}
	name := namedRef.Name()
	i := strings.IndexRune(name, '/')
	if i == -1 || (!strings.ContainsAny(name[:i], ":.") && name[:i] != "localhost") {
		ref.Name = name
	} else {
		ref.Registry, ref.Name = name[:i], name[i+1:]
	}
	if named, ok := namedRef.(reference.NamedTagged); ok {
		ref.Tag = named.Tag()
	}
	if named, ok := namedRef.(reference.Canonical); ok {
		ref.ID = named.Digest().String()
	}
	if i := strings.IndexRune(ref.Name, '/'); i != -1 {
		ref.Namespace, ref.Name = ref.Name[:i], ref.Name[i+1:]
	}
	return ref, nil
}
func (r DockerImageReference) Equal(other DockerImageReference) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defaultedRef := r.DockerClientDefaults()
	otherDefaultedRef := other.DockerClientDefaults()
	return defaultedRef == otherDefaultedRef
}
func (r DockerImageReference) DockerClientDefaults() DockerImageReference {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(r.Registry) == 0 {
		r.Registry = DockerDefaultRegistry
	}
	if len(r.Namespace) == 0 && IsRegistryDockerHub(r.Registry) {
		r.Namespace = "library"
	}
	if len(r.Tag) == 0 {
		r.Tag = "latest"
	}
	return r
}
func (r DockerImageReference) Minimal() DockerImageReference {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if r.Tag == "latest" {
		r.Tag = ""
	}
	return r
}
func (r DockerImageReference) AsRepository() DockerImageReference {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	r.Tag = ""
	r.ID = ""
	return r
}
func (r DockerImageReference) RepositoryName() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	r.Tag = ""
	r.ID = ""
	r.Registry = ""
	return r.Exact()
}
func (r DockerImageReference) RegistryHostPort(insecure bool) (string, string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	registryHost := r.AsV2().DockerClientDefaults().Registry
	if strings.Contains(registryHost, ":") {
		hostname, port, _ := net.SplitHostPort(registryHost)
		return hostname, port
	}
	if insecure {
		return registryHost, "80"
	}
	return registryHost, "443"
}
func (r DockerImageReference) RegistryURL() *url.URL {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &url.URL{Scheme: "https", Host: r.AsV2().Registry}
}
func (r DockerImageReference) DaemonMinimal() DockerImageReference {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch r.Registry {
	case DockerDefaultV1Registry, DockerDefaultV2Registry:
		r.Registry = DockerDefaultRegistry
	}
	if IsRegistryDockerHub(r.Registry) && r.Namespace == "library" {
		r.Namespace = ""
	}
	return r.Minimal()
}
func (r DockerImageReference) AsV2() DockerImageReference {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch r.Registry {
	case DockerDefaultV1Registry, DockerDefaultRegistry:
		r.Registry = DockerDefaultV2Registry
	}
	return r
}
func (r DockerImageReference) MostSpecific() DockerImageReference {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(r.ID) == 0 {
		return r
	}
	if _, err := digest.ParseDigest(r.ID); err == nil {
		r.Tag = ""
		return r
	}
	if len(r.Tag) == 0 {
		r.Tag, r.ID = r.ID, ""
		return r
	}
	return r
}
func (r DockerImageReference) NameString() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch {
	case len(r.Name) == 0:
		return ""
	case len(r.Tag) > 0:
		return r.Name + ":" + r.Tag
	case len(r.ID) > 0:
		var ref string
		if _, err := digest.ParseDigest(r.ID); err == nil {
			ref = "@" + r.ID
		} else {
			ref = ":" + r.ID
		}
		return r.Name + ref
	default:
		return r.Name
	}
}
func (r DockerImageReference) Exact() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	name := r.NameString()
	if len(name) == 0 {
		return name
	}
	s := r.Registry
	if len(s) > 0 {
		s += "/"
	}
	if len(r.Namespace) != 0 {
		s += r.Namespace + "/"
	}
	return s + name
}
func (r DockerImageReference) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(r.Namespace) == 0 && IsRegistryDockerHub(r.Registry) {
		r.Namespace = "library"
	}
	return r.Exact()
}
func IsRegistryDockerHub(registry string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch registry {
	case DockerDefaultRegistry, DockerDefaultV1Registry, DockerDefaultV2Registry:
		return true
	default:
		return false
	}
}
func (in *DockerImageReference) DeepCopyInto(out *DockerImageReference) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	*out = *in
	return
}
func (in *DockerImageReference) DeepCopy() *DockerImageReference {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if in == nil {
		return nil
	}
	out := new(DockerImageReference)
	in.DeepCopyInto(out)
	return out
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
