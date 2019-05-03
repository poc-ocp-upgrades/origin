package serviceaccount

import (
 "context"
 "crypto/ecdsa"
 "crypto/elliptic"
 "crypto/rsa"
 "encoding/base64"
 "encoding/json"
 "fmt"
 "strings"
 jose "gopkg.in/square/go-jose.v2"
 "gopkg.in/square/go-jose.v2/jwt"
 "k8s.io/api/core/v1"
 utilerrors "k8s.io/apimachinery/pkg/util/errors"
 "k8s.io/apiserver/pkg/authentication/authenticator"
)

type ServiceAccountTokenGetter interface {
 GetServiceAccount(namespace, name string) (*v1.ServiceAccount, error)
 GetPod(namespace, name string) (*v1.Pod, error)
 GetSecret(namespace, name string) (*v1.Secret, error)
}
type TokenGenerator interface {
 GenerateToken(claims *jwt.Claims, privateClaims interface{}) (string, error)
}

func JWTTokenGenerator(iss string, privateKey interface{}) (TokenGenerator, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var alg jose.SignatureAlgorithm
 switch pk := privateKey.(type) {
 case *rsa.PrivateKey:
  alg = jose.RS256
 case *ecdsa.PrivateKey:
  switch pk.Curve {
  case elliptic.P256():
   alg = jose.ES256
  case elliptic.P384():
   alg = jose.ES384
  case elliptic.P521():
   alg = jose.ES512
  default:
   return nil, fmt.Errorf("unknown private key curve, must be 256, 384, or 521")
  }
 case jose.OpaqueSigner:
  alg = jose.SignatureAlgorithm(pk.Public().Algorithm)
 default:
  return nil, fmt.Errorf("unknown private key type %T, must be *rsa.PrivateKey, *ecdsa.PrivateKey, or jose.OpaqueSigner", privateKey)
 }
 signer, err := jose.NewSigner(jose.SigningKey{Algorithm: alg, Key: privateKey}, nil)
 if err != nil {
  return nil, err
 }
 return &jwtTokenGenerator{iss: iss, signer: signer}, nil
}

type jwtTokenGenerator struct {
 iss    string
 signer jose.Signer
}

func (j *jwtTokenGenerator) GenerateToken(claims *jwt.Claims, privateClaims interface{}) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return jwt.Signed(j.signer).Claims(privateClaims).Claims(claims).Claims(&jwt.Claims{Issuer: j.iss}).CompactSerialize()
}
func JWTTokenAuthenticator(iss string, keys []interface{}, implicitAuds authenticator.Audiences, validator Validator) authenticator.Token {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &jwtTokenAuthenticator{iss: iss, keys: keys, implicitAuds: implicitAuds, validator: validator}
}

type jwtTokenAuthenticator struct {
 iss          string
 keys         []interface{}
 validator    Validator
 implicitAuds authenticator.Audiences
}
type Validator interface {
 Validate(tokenData string, public *jwt.Claims, private interface{}) (*ServiceAccountInfo, error)
 NewPrivateClaims() interface{}
}

func (j *jwtTokenAuthenticator) AuthenticateToken(ctx context.Context, tokenData string) (*authenticator.Response, bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !j.hasCorrectIssuer(tokenData) {
  return nil, false, nil
 }
 tok, err := jwt.ParseSigned(tokenData)
 if err != nil {
  return nil, false, nil
 }
 public := &jwt.Claims{}
 private := j.validator.NewPrivateClaims()
 var (
  found   bool
  errlist []error
 )
 for _, key := range j.keys {
  if err := tok.Claims(key, public, private); err != nil {
   errlist = append(errlist, err)
   continue
  }
  found = true
  break
 }
 if !found {
  return nil, false, utilerrors.NewAggregate(errlist)
 }
 tokenAudiences := authenticator.Audiences(public.Audience)
 if len(tokenAudiences) == 0 {
  tokenAudiences = j.implicitAuds
 }
 requestedAudiences, ok := authenticator.AudiencesFrom(ctx)
 if !ok {
  requestedAudiences = j.implicitAuds
 }
 auds := authenticator.Audiences(tokenAudiences).Intersect(requestedAudiences)
 if len(auds) == 0 && len(j.implicitAuds) != 0 {
  return nil, false, fmt.Errorf("token audiences %q is invalid for the target audiences %q", tokenAudiences, requestedAudiences)
 }
 sa, err := j.validator.Validate(tokenData, public, private)
 if err != nil {
  return nil, false, err
 }
 return &authenticator.Response{User: sa.UserInfo(), Audiences: auds}, true, nil
}
func (j *jwtTokenAuthenticator) hasCorrectIssuer(tokenData string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 parts := strings.Split(tokenData, ".")
 if len(parts) != 3 {
  return false
 }
 payload, err := base64.RawURLEncoding.DecodeString(parts[1])
 if err != nil {
  return false
 }
 claims := struct {
  Issuer string `json:"iss"`
 }{}
 if err := json.Unmarshal(payload, &claims); err != nil {
  return false
 }
 if claims.Issuer != j.iss {
  return false
 }
 return true
}
