package pubkeypin

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	goformat "fmt"
	"github.com/pkg/errors"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

const (
	formatSHA256 = "sha256"
)

type Set struct{ sha256Hashes map[string]bool }

func NewSet() *Set {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &Set{make(map[string]bool)}
}
func (s *Set) Allow(pubKeyHashes ...string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, pubKeyHash := range pubKeyHashes {
		parts := strings.Split(pubKeyHash, ":")
		if len(parts) != 2 {
			return errors.New("invalid public key hash, expected \"format:value\"")
		}
		format, value := parts[0], parts[1]
		switch strings.ToLower(format) {
		case "sha256":
			return s.allowSHA256(value)
		default:
			return errors.Errorf("unknown hash format %q", format)
		}
	}
	return nil
}
func (s *Set) Check(certificate *x509.Certificate) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if s.checkSHA256(certificate) {
		return nil
	}
	return errors.Errorf("public key %s not pinned", Hash(certificate))
}
func (s *Set) Empty() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return len(s.sha256Hashes) == 0
}
func Hash(certificate *x509.Certificate) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	spkiHash := sha256.Sum256(certificate.RawSubjectPublicKeyInfo)
	return formatSHA256 + ":" + strings.ToLower(hex.EncodeToString(spkiHash[:]))
}
func (s *Set) allowSHA256(hash string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	hashLength := hex.DecodedLen(len(hash))
	if hashLength != sha256.Size {
		return errors.Errorf("expected a %d byte SHA-256 hash, found %d bytes", sha256.Size, hashLength)
	}
	_, err := hex.DecodeString(hash)
	if err != nil {
		return err
	}
	s.sha256Hashes[strings.ToLower(hash)] = true
	return nil
}
func (s *Set) checkSHA256(certificate *x509.Certificate) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	actualHash := sha256.Sum256(certificate.RawSubjectPublicKeyInfo)
	actualHashHex := strings.ToLower(hex.EncodeToString(actualHash[:]))
	return s.sha256Hashes[actualHashHex]
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
