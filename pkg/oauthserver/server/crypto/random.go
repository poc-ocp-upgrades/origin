package crypto

import (
	"crypto/rand"
	"encoding/base64"
)

func RandomBits(bits int) []byte {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	size := bits / 8
	if bits%8 != 0 {
		size++
	}
	b := make([]byte, size)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return b
}
func RandomBitsString(bits int) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return base64.RawURLEncoding.EncodeToString(RandomBits(bits))
}
func Random256BitsString() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return RandomBitsString(256)
}
