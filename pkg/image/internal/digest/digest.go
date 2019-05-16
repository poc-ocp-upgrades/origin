package digest

import (
	"fmt"
	goformat "fmt"
	"hash"
	"io"
	goos "os"
	"regexp"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

const (
	DigestSha256EmptyTar = "sha256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
)

type Digest string

func NewDigest(alg Algorithm, h hash.Hash) Digest {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return NewDigestFromBytes(alg, h.Sum(nil))
}
func NewDigestFromBytes(alg Algorithm, p []byte) Digest {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return Digest(fmt.Sprintf("%s:%x", alg, p))
}
func NewDigestFromHex(alg, hex string) Digest {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return Digest(fmt.Sprintf("%s:%s", alg, hex))
}

var DigestRegexp = regexp.MustCompile(`[a-zA-Z0-9-_+.]+:[a-fA-F0-9]+`)
var DigestRegexpAnchored = regexp.MustCompile(`^` + DigestRegexp.String() + `$`)
var (
	ErrDigestInvalidFormat = fmt.Errorf("invalid checksum digest format")
	ErrDigestInvalidLength = fmt.Errorf("invalid checksum digest length")
	ErrDigestUnsupported   = fmt.Errorf("unsupported digest algorithm")
)

func ParseDigest(s string) (Digest, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	d := Digest(s)
	return d, d.Validate()
}
func FromReader(rd io.Reader) (Digest, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return Canonical.FromReader(rd)
}
func FromBytes(p []byte) Digest {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return Canonical.FromBytes(p)
}
func (d Digest) Validate() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	s := string(d)
	if !DigestRegexpAnchored.MatchString(s) {
		return ErrDigestInvalidFormat
	}
	i := strings.Index(s, ":")
	if i < 0 {
		return ErrDigestInvalidFormat
	}
	if i+1 == len(s) {
		return ErrDigestInvalidFormat
	}
	switch algorithm := Algorithm(s[:i]); algorithm {
	case SHA256, SHA384, SHA512:
		if algorithm.Size()*2 != len(s[i+1:]) {
			return ErrDigestInvalidLength
		}
	default:
		return ErrDigestUnsupported
	}
	return nil
}
func (d Digest) Algorithm() Algorithm {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return Algorithm(d[:d.sepIndex()])
}
func (d Digest) Hex() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return string(d[d.sepIndex()+1:])
}
func (d Digest) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return string(d)
}
func (d Digest) sepIndex() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	i := strings.Index(string(d), ":")
	if i < 0 {
		panic("could not find ':' in digest: " + d)
	}
	return i
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
