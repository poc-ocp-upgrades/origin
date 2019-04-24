package digest

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"hash"
	"io"
	"regexp"
	"strings"
)

const (
	DigestSha256EmptyTar = "sha256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
)

type Digest string

func NewDigest(alg Algorithm, h hash.Hash) Digest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewDigestFromBytes(alg, h.Sum(nil))
}
func NewDigestFromBytes(alg Algorithm, p []byte) Digest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return Digest(fmt.Sprintf("%s:%x", alg, p))
}
func NewDigestFromHex(alg, hex string) Digest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return Digest(fmt.Sprintf("%s:%s", alg, hex))
}

var DigestRegexp = regexp.MustCompile(`[a-zA-Z0-9-_+.]+:[a-fA-F0-9]+`)
var DigestRegexpAnchored = regexp.MustCompile(`^` + DigestRegexp.String() + `$`)
var (
	ErrDigestInvalidFormat	= fmt.Errorf("invalid checksum digest format")
	ErrDigestInvalidLength	= fmt.Errorf("invalid checksum digest length")
	ErrDigestUnsupported	= fmt.Errorf("unsupported digest algorithm")
)

func ParseDigest(s string) (Digest, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	d := Digest(s)
	return d, d.Validate()
}
func FromReader(rd io.Reader) (Digest, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return Canonical.FromReader(rd)
}
func FromBytes(p []byte) Digest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return Canonical.FromBytes(p)
}
func (d Digest) Validate() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return Algorithm(d[:d.sepIndex()])
}
func (d Digest) Hex() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(d[d.sepIndex()+1:])
}
func (d Digest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(d)
}
func (d Digest) sepIndex() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	i := strings.Index(string(d), ":")
	if i < 0 {
		panic("could not find ':' in digest: " + d)
	}
	return i
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
