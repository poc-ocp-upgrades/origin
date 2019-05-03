package bootstrap

import (
 "fmt"
 "strings"
 jose "gopkg.in/square/go-jose.v2"
)

func computeDetachedSig(content, tokenID, tokenSecret string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 jwk := &jose.JSONWebKey{Key: []byte(tokenSecret), KeyID: tokenID}
 opts := &jose.SignerOptions{ExtraHeaders: map[jose.HeaderKey]interface{}{"kid": tokenID}}
 signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: jwk}, opts)
 if err != nil {
  return "", fmt.Errorf("can't make a HS256 signer from the given token: %v", err)
 }
 jws, err := signer.Sign([]byte(content))
 if err != nil {
  return "", fmt.Errorf("can't HS256-sign the given token: %v", err)
 }
 fullSig, err := jws.CompactSerialize()
 if err != nil {
  return "", fmt.Errorf("can't serialize the given token: %v", err)
 }
 return stripContent(fullSig)
}
func stripContent(fullSig string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 parts := strings.Split(fullSig, ".")
 if len(parts) != 3 {
  return "", fmt.Errorf("compact JWS format must have three parts")
 }
 return parts[0] + ".." + parts[2], nil
}
func DetachedTokenIsValid(detachedToken, content, tokenID, tokenSecret string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newToken, err := computeDetachedSig(content, tokenID, tokenSecret)
 if err != nil {
  return false
 }
 return detachedToken == newToken
}
