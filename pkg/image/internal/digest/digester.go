package digest

import (
	"crypto"
	"fmt"
	"hash"
	"io"
)

type Algorithm string

const (
	SHA256    Algorithm = "sha256"
	SHA384    Algorithm = "sha384"
	SHA512    Algorithm = "sha512"
	Canonical           = SHA256
)

var (
	algorithms = map[Algorithm]crypto.Hash{SHA256: crypto.SHA256, SHA384: crypto.SHA384, SHA512: crypto.SHA512}
)

func (a Algorithm) Available() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	h, ok := algorithms[a]
	if !ok {
		return false
	}
	return h.Available()
}
func (a Algorithm) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return string(a)
}
func (a Algorithm) Size() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	h, ok := algorithms[a]
	if !ok {
		return 0
	}
	return h.Size()
}
func (a *Algorithm) Set(value string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if value == "" {
		*a = Canonical
	} else {
		*a = Algorithm(value)
	}
	return nil
}
func (a Algorithm) New() Digester {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &digester{alg: a, hash: a.Hash()}
}
func (a Algorithm) Hash() hash.Hash {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !a.Available() {
		panic(fmt.Sprintf("%v not available (make sure it is imported)", a))
	}
	return algorithms[a].New()
}
func (a Algorithm) FromReader(rd io.Reader) (Digest, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	digester := a.New()
	if _, err := io.Copy(digester.Hash(), rd); err != nil {
		return "", err
	}
	return digester.Digest(), nil
}
func (a Algorithm) FromBytes(p []byte) Digest {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
	alg  Algorithm
	hash hash.Hash
}

func (d *digester) Hash() hash.Hash {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return d.hash
}
func (d *digester) Digest() Digest {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return NewDigest(d.alg, d.hash)
}
