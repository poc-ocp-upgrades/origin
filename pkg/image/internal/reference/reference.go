package reference

import (
	"errors"
	"fmt"
	goformat "fmt"
	"github.com/openshift/origin/pkg/image/internal/digest"
	goos "os"
	"path"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

const (
	NameTotalLengthMax = 255
)

var (
	ErrReferenceInvalidFormat = errors.New("invalid reference format")
	ErrTagInvalidFormat       = errors.New("invalid tag format")
	ErrDigestInvalidFormat    = errors.New("invalid digest format")
	ErrNameContainsUppercase  = errors.New("repository name must be lowercase")
	ErrNameEmpty              = errors.New("repository name must have at least one component")
	ErrNameTooLong            = fmt.Errorf("repository name must not be more than %v characters", NameTotalLengthMax)
)

type Reference interface{ String() string }
type Field struct{ reference Reference }

func AsField(reference Reference) Field {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return Field{reference}
}
func (f Field) Reference() Reference {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return f.reference
}
func (f Field) MarshalText() (p []byte, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return []byte(f.reference.String()), nil
}
func (f *Field) UnmarshalText(p []byte) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	r, err := Parse(string(p))
	if err != nil {
		return err
	}
	f.reference = r
	return nil
}

type Named interface {
	Reference
	Name() string
}
type Tagged interface {
	Reference
	Tag() string
}
type NamedTagged interface {
	Named
	Tag() string
}
type Digested interface {
	Reference
	Digest() digest.Digest
}
type Canonical interface {
	Named
	Digest() digest.Digest
}

func SplitHostname(named Named) (string, string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	name := named.Name()
	match := anchoredNameRegexp.FindStringSubmatch(name)
	if len(match) != 3 {
		return "", name
	}
	return match[1], match[2]
}
func Parse(s string) (Reference, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	matches := ReferenceRegexp.FindStringSubmatch(s)
	if matches == nil {
		if s == "" {
			return nil, ErrNameEmpty
		}
		if ReferenceRegexp.FindStringSubmatch(strings.ToLower(s)) != nil {
			return nil, ErrNameContainsUppercase
		}
		return nil, ErrReferenceInvalidFormat
	}
	if len(matches[1]) > NameTotalLengthMax {
		return nil, ErrNameTooLong
	}
	ref := reference{name: matches[1], tag: matches[2]}
	if matches[3] != "" {
		var err error
		ref.digest, err = digest.ParseDigest(matches[3])
		if err != nil {
			return nil, err
		}
	}
	r := getBestReferenceType(ref)
	if r == nil {
		return nil, ErrNameEmpty
	}
	return r, nil
}
func ParseNamed(s string) (Named, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ref, err := Parse(s)
	if err != nil {
		return nil, err
	}
	named, isNamed := ref.(Named)
	if !isNamed {
		return nil, fmt.Errorf("reference %s has no name", ref.String())
	}
	return named, nil
}
func WithName(name string) (Named, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(name) > NameTotalLengthMax {
		return nil, ErrNameTooLong
	}
	if !anchoredNameRegexp.MatchString(name) {
		return nil, ErrReferenceInvalidFormat
	}
	return repository(name), nil
}
func WithTag(name Named, tag string) (NamedTagged, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !anchoredTagRegexp.MatchString(tag) {
		return nil, ErrTagInvalidFormat
	}
	if canonical, ok := name.(Canonical); ok {
		return reference{name: name.Name(), tag: tag, digest: canonical.Digest()}, nil
	}
	return taggedReference{name: name.Name(), tag: tag}, nil
}
func WithDigest(name Named, digest digest.Digest) (Canonical, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !anchoredDigestRegexp.MatchString(digest.String()) {
		return nil, ErrDigestInvalidFormat
	}
	if tagged, ok := name.(Tagged); ok {
		return reference{name: name.Name(), tag: tagged.Tag(), digest: digest}, nil
	}
	return canonicalReference{name: name.Name(), digest: digest}, nil
}
func Match(pattern string, ref Reference) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	matched, err := path.Match(pattern, ref.String())
	if namedRef, isNamed := ref.(Named); isNamed && !matched {
		matched, _ = path.Match(pattern, namedRef.Name())
	}
	return matched, err
}
func TrimNamed(ref Named) Named {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return repository(ref.Name())
}
func getBestReferenceType(ref reference) Reference {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if ref.name == "" {
		if ref.digest != "" {
			return digestReference(ref.digest)
		}
		return nil
	}
	if ref.tag == "" {
		if ref.digest != "" {
			return canonicalReference{name: ref.name, digest: ref.digest}
		}
		return repository(ref.name)
	}
	if ref.digest == "" {
		return taggedReference{name: ref.name, tag: ref.tag}
	}
	return ref
}

type reference struct {
	name   string
	tag    string
	digest digest.Digest
}

func (r reference) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return r.name + ":" + r.tag + "@" + r.digest.String()
}
func (r reference) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return r.name
}
func (r reference) Tag() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return r.tag
}
func (r reference) Digest() digest.Digest {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return r.digest
}

type repository string

func (r repository) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return string(r)
}
func (r repository) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return string(r)
}

type digestReference digest.Digest

func (d digestReference) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return string(d)
}
func (d digestReference) Digest() digest.Digest {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return digest.Digest(d)
}

type taggedReference struct {
	name string
	tag  string
}

func (t taggedReference) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return t.name + ":" + t.tag
}
func (t taggedReference) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return t.name
}
func (t taggedReference) Tag() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return t.tag
}

type canonicalReference struct {
	name   string
	digest digest.Digest
}

func (c canonicalReference) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.name + "@" + c.digest.String()
}
func (c canonicalReference) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.name
}
func (c canonicalReference) Digest() digest.Digest {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.digest
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
