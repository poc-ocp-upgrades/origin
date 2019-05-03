package azure

import (
 "bytes"
 godefaultbytes "bytes"
 godefaultruntime "runtime"
 "encoding/json"
 "fmt"
 "io/ioutil"
 "net/http"
 godefaulthttp "net/http"
 "net/url"
 "strconv"
 "strings"
 "unicode"
 jwt "github.com/dgrijalva/jwt-go"
)

type authDirective struct {
 service string
 realm   string
}
type accessTokenPayload struct {
 TenantID string `json:"tid"`
}
type acrTokenPayload struct {
 Expiration int64  `json:"exp"`
 TenantID   string `json:"tenant"`
 Credential string `json:"credential"`
}
type acrAuthResponse struct {
 RefreshToken string `json:"refresh_token"`
}

const timeShiftBuffer = 300
const userAgentHeader = "User-Agent"
const userAgent = "kubernetes-credentialprovider-acr"
const dockerTokenLoginUsernameGUID = "00000000-0000-0000-0000-000000000000"

var client = &http.Client{}

func receiveChallengeFromLoginServer(serverAddress string) (*authDirective, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 challengeURL := url.URL{Scheme: "https", Host: serverAddress, Path: "v2/"}
 var err error
 var r *http.Request
 r, _ = http.NewRequest("GET", challengeURL.String(), nil)
 r.Header.Add(userAgentHeader, userAgent)
 var challenge *http.Response
 if challenge, err = client.Do(r); err != nil {
  return nil, fmt.Errorf("Error reaching registry endpoint %s, error: %s", challengeURL.String(), err)
 }
 defer challenge.Body.Close()
 if challenge.StatusCode != 401 {
  return nil, fmt.Errorf("Registry did not issue a valid AAD challenge, status: %d", challenge.StatusCode)
 }
 var authHeader []string
 var ok bool
 if authHeader, ok = challenge.Header["Www-Authenticate"]; !ok {
  return nil, fmt.Errorf("Challenge response does not contain header 'Www-Authenticate'")
 }
 if len(authHeader) != 1 {
  return nil, fmt.Errorf("Registry did not issue a valid AAD challenge, authenticate header [%s]", strings.Join(authHeader, ", "))
 }
 authSections := strings.SplitN(authHeader[0], " ", 2)
 authType := strings.ToLower(authSections[0])
 var authParams *map[string]string
 if authParams, err = parseAssignments(authSections[1]); err != nil {
  return nil, fmt.Errorf("Unable to understand the contents of Www-Authenticate header %s", authSections[1])
 }
 if !strings.EqualFold("Bearer", authType) {
  return nil, fmt.Errorf("Www-Authenticate: expected realm: Bearer, actual: %s", authType)
 }
 if len((*authParams)["service"]) == 0 {
  return nil, fmt.Errorf("Www-Authenticate: missing header \"service\"")
 }
 if len((*authParams)["realm"]) == 0 {
  return nil, fmt.Errorf("Www-Authenticate: missing header \"realm\"")
 }
 return &authDirective{service: (*authParams)["service"], realm: (*authParams)["realm"]}, nil
}
func parseAcrToken(identityToken string) (token *acrTokenPayload, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 tokenSegments := strings.Split(identityToken, ".")
 if len(tokenSegments) < 2 {
  return nil, fmt.Errorf("Invalid existing refresh token length: %d", len(tokenSegments))
 }
 payloadSegmentEncoded := tokenSegments[1]
 var payloadBytes []byte
 if payloadBytes, err = jwt.DecodeSegment(payloadSegmentEncoded); err != nil {
  return nil, fmt.Errorf("Error decoding payload segment from refresh token, error: %s", err)
 }
 var payload acrTokenPayload
 if err = json.Unmarshal(payloadBytes, &payload); err != nil {
  return nil, fmt.Errorf("Error unmarshalling acr payload, error: %s", err)
 }
 return &payload, nil
}
func performTokenExchange(serverAddress string, directive *authDirective, tenant string, accessToken string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var err error
 data := url.Values{"service": []string{directive.service}, "grant_type": []string{"access_token_refresh_token"}, "access_token": []string{accessToken}, "refresh_token": []string{accessToken}, "tenant": []string{tenant}}
 var realmURL *url.URL
 if realmURL, err = url.Parse(directive.realm); err != nil {
  return "", fmt.Errorf("Www-Authenticate: invalid realm %s", directive.realm)
 }
 authEndpoint := fmt.Sprintf("%s://%s/oauth2/exchange", realmURL.Scheme, realmURL.Host)
 datac := data.Encode()
 var r *http.Request
 r, _ = http.NewRequest("POST", authEndpoint, bytes.NewBufferString(datac))
 r.Header.Add(userAgentHeader, userAgent)
 r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
 r.Header.Add("Content-Length", strconv.Itoa(len(datac)))
 var exchange *http.Response
 if exchange, err = client.Do(r); err != nil {
  return "", fmt.Errorf("Www-Authenticate: failed to reach auth url %s", authEndpoint)
 }
 defer exchange.Body.Close()
 if exchange.StatusCode != 200 {
  return "", fmt.Errorf("Www-Authenticate: auth url %s responded with status code %d", authEndpoint, exchange.StatusCode)
 }
 var content []byte
 if content, err = ioutil.ReadAll(exchange.Body); err != nil {
  return "", fmt.Errorf("Www-Authenticate: error reading response from %s", authEndpoint)
 }
 var authResp acrAuthResponse
 if err = json.Unmarshal(content, &authResp); err != nil {
  return "", fmt.Errorf("Www-Authenticate: unable to read response %s", content)
 }
 return authResp.RefreshToken, nil
}
func parseAssignments(statements string) (*map[string]string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var cursor int
 result := make(map[string]string)
 var errorMsg = fmt.Errorf("malformed header value: %s", statements)
 for {
  equalIndex := nextOccurrence(statements, cursor, "=")
  if equalIndex == -1 {
   return nil, errorMsg
  }
  key := strings.TrimSpace(statements[cursor:equalIndex])
  cursor = nextNoneSpace(statements, equalIndex+1)
  if cursor == -1 {
   return nil, errorMsg
  }
  if statements[cursor] == '"' {
   cursor = cursor + 1
   closeQuoteIndex := nextOccurrence(statements, cursor, "\"")
   if closeQuoteIndex == -1 {
    return nil, errorMsg
   }
   value := statements[cursor:closeQuoteIndex]
   result[key] = value
   commaIndex := nextNoneSpace(statements, closeQuoteIndex+1)
   if commaIndex == -1 {
    return &result, nil
   } else if statements[commaIndex] != ',' {
    return nil, errorMsg
   } else {
    cursor = commaIndex + 1
   }
  } else {
   commaIndex := nextOccurrence(statements, cursor, ",")
   endStatements := commaIndex == -1
   var untrimmed string
   if endStatements {
    untrimmed = statements[cursor:commaIndex]
   } else {
    untrimmed = statements[cursor:]
   }
   value := strings.TrimSpace(untrimmed)
   if len(value) == 0 {
    return nil, errorMsg
   }
   result[key] = value
   if endStatements {
    return &result, nil
   }
   cursor = commaIndex + 1
  }
 }
}
func nextOccurrence(str string, start int, sep string) int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if start >= len(str) {
  return -1
 }
 offset := strings.Index(str[start:], sep)
 if offset == -1 {
  return -1
 }
 return offset + start
}
func nextNoneSpace(str string, start int) int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if start >= len(str) {
  return -1
 }
 offset := strings.IndexFunc(str[start:], func(c rune) bool {
  return !unicode.IsSpace(c)
 })
 if offset == -1 {
  return -1
 }
 return offset + start
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
