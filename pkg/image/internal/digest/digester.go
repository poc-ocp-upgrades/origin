package digest

import (
	"crypto"
	"fmt"
	"hash"
	"io"
)

type Algorithm string

const (
	SHA256		Algorithm	= "sha256"
	SHA384		Algorithm	= "sha384"
	SHA512		Algorithm	= "sha512"
	Canonical			= SHA256
)

var (
	algorithms = map[Algorithm]crypto.Hash{SHA256: crypto.SHA256, SHA384: crypto.SHA384, SHA512: crypto.SHA512}
)

func (a Algorithm) Available() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	h, ok := algorithms[a]
	if !ok {
		return false
	}
	return h.Available()
}
func (a Algorithm) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(a)
}
func (a Algorithm) Size() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	h, ok := algorithms[a]
	if !ok {
		return 0
	}
	return h.Size()
}
func (a *Algorithm) Set(value string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if value == "" {
		*a = Canonical
	} else {
		*a = Algorithm(value)
	}
	return nil
}
func (a Algorithm) New() Digester {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &digester{alg: a, hash: a.Hash()}
}
func (a Algorithm) Hash() hash.Hash {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !a.Available() {
		panic(fmt.Sprintf("%v not available (make sure it is imported)", a))
	}
	return algorithms[a].New()
}
func (a Algorithm) FromReader(rd io.Reader) (Digest, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	digester := a.New()
	if _, err := io.Copy(digester.Hash(), rd); err != nil {
		return "", err
	}
	return digester.Digest(), nil
}
func (a Algorithm) FromBytes(p []byte) Digest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	digester := a.New()
	if _, err := digester.Hash().Write(p); err != nil {
		panic("write to hash function returned error: " + err.Error())
	}
	return digester.Digest()
}

type Digester interface {
	Hash() hash.Hash
	Digest() Digest
}
type digester struct {
	alg	Algorithm
	hash	hash.Hash
}

func (d *digester) Hash() hash.Hash {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return d.hash
}
func (d *digester) Digest() Digest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewDigest(d.alg, d.hash)
}
