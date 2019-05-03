package gce

import (
 "encoding/json"
 "net/http"
 "strings"
 "time"
 "k8s.io/client-go/util/flowcontrol"
 "github.com/prometheus/client_golang/prometheus"
 "golang.org/x/oauth2"
 "golang.org/x/oauth2/google"
 "google.golang.org/api/googleapi"
)

const (
 tokenURLQPS   = .05
 tokenURLBurst = 3
)

var (
 getTokenCounter     = prometheus.NewCounter(prometheus.CounterOpts{Name: "get_token_count", Help: "Counter of total Token() requests to the alternate token source"})
 getTokenFailCounter = prometheus.NewCounter(prometheus.CounterOpts{Name: "get_token_fail_count", Help: "Counter of failed Token() requests to the alternate token source"})
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 prometheus.MustRegister(getTokenCounter)
 prometheus.MustRegister(getTokenFailCounter)
}

type AltTokenSource struct {
 oauthClient *http.Client
 tokenURL    string
 tokenBody   string
 throttle    flowcontrol.RateLimiter
}

func (a *AltTokenSource) Token() (*oauth2.Token, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 a.throttle.Accept()
 getTokenCounter.Inc()
 t, err := a.token()
 if err != nil {
  getTokenFailCounter.Inc()
 }
 return t, err
}
func (a *AltTokenSource) token() (*oauth2.Token, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 req, err := http.NewRequest("POST", a.tokenURL, strings.NewReader(a.tokenBody))
 if err != nil {
  return nil, err
 }
 res, err := a.oauthClient.Do(req)
 if err != nil {
  return nil, err
 }
 defer res.Body.Close()
 if err := googleapi.CheckResponse(res); err != nil {
  return nil, err
 }
 var tok struct {
  AccessToken string    `json:"accessToken"`
  ExpireTime  time.Time `json:"expireTime"`
 }
 if err := json.NewDecoder(res.Body).Decode(&tok); err != nil {
  return nil, err
 }
 return &oauth2.Token{AccessToken: tok.AccessToken, Expiry: tok.ExpireTime}, nil
}
func NewAltTokenSource(tokenURL, tokenBody string) oauth2.TokenSource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 client := oauth2.NewClient(oauth2.NoContext, google.ComputeTokenSource(""))
 a := &AltTokenSource{oauthClient: client, tokenURL: tokenURL, tokenBody: tokenBody, throttle: flowcontrol.NewTokenBucketRateLimiter(tokenURLQPS, tokenURLBurst)}
 return oauth2.ReuseTokenSource(nil, a)
}
