package tokencmd

import (
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
	"k8s.io/klog"
)

type Negotiator interface {
	Load() error
	InitSecContext(requestURL string, challengeToken []byte) (tokenToSend []byte, err error)
	IsComplete() bool
	Release() error
}
type NegotiateChallengeHandler struct{ negotiator Negotiator }

func NewNegotiateChallengeHandler(negotiator Negotiator) ChallengeHandler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &NegotiateChallengeHandler{negotiator: negotiator}
}
func (c *NegotiateChallengeHandler) CanHandle(headers http.Header) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if isNegotiate, _, err := getNegotiateToken(headers); err != nil || !isNegotiate {
		return false
	}
	if err := c.negotiator.Load(); err != nil {
		return false
	}
	return true
}
func (c *NegotiateChallengeHandler) HandleChallenge(requestURL string, headers http.Header) (http.Header, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, incomingToken, err := getNegotiateToken(headers)
	if err != nil {
		return nil, false, err
	}
	outgoingToken, err := c.negotiator.InitSecContext(requestURL, incomingToken)
	if err != nil {
		klog.V(5).Infof("InitSecContext returned error: %v", err)
		return nil, false, err
	}
	responseHeaders := http.Header{}
	responseHeaders.Set("Authorization", "Negotiate "+base64.StdEncoding.EncodeToString(outgoingToken))
	return responseHeaders, true, nil
}
func (c *NegotiateChallengeHandler) CompleteChallenge(requestURL string, headers http.Header) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c.negotiator.IsComplete() {
		return nil
	}
	klog.V(5).Infof("continue needed")
	isNegotiate, incomingToken, err := getNegotiateToken(headers)
	if err != nil {
		return err
	}
	if !isNegotiate {
		return errors.New("client requires final negotiate token, none provided")
	}
	_, err = c.negotiator.InitSecContext(requestURL, incomingToken)
	if err != nil {
		klog.V(5).Infof("InitSecContext returned error during final negotiation: %v", err)
		return err
	}
	if !c.negotiator.IsComplete() {
		return errors.New("InitSecContext did not indicate final negotiation completed")
	}
	return nil
}
func (c *NegotiateChallengeHandler) Release() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.negotiator.Release()
}

const negotiateScheme = "negotiate"

func getNegotiateToken(headers http.Header) (bool, []byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, challengeHeader := range headers[http.CanonicalHeaderKey("WWW-Authenticate")] {
		caseInsensitiveHeader := strings.ToLower(challengeHeader)
		if caseInsensitiveHeader == negotiateScheme {
			return true, nil, nil
		}
		if strings.HasPrefix(caseInsensitiveHeader, negotiateScheme+" ") {
			payload := challengeHeader[len(negotiateScheme):]
			payload = strings.Replace(payload, " ", "", -1)
			data, err := base64.StdEncoding.DecodeString(payload)
			if err != nil {
				return false, nil, err
			}
			return true, data, nil
		}
	}
	return false, nil, nil
}
