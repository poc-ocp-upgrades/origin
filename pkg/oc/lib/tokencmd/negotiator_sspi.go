package tokencmd

import (
	"fmt"
	"io"
	"os"
	"strings"
	"syscall"
	"k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/runtime"
	"github.com/openshift/origin/pkg/cmd/util/term"
	"github.com/alexbrainman/sspi"
	"github.com/alexbrainman/sspi/negotiate"
	"k8s.io/klog"
)

const (
	desiredFlags		= sspi.ISC_REQ_CONFIDENTIALITY | sspi.ISC_REQ_INTEGRITY | sspi.ISC_REQ_MUTUAL_AUTH | sspi.ISC_REQ_REPLAY_DETECT | sspi.ISC_REQ_SEQUENCE_DETECT
	requiredFlags		= sspi.ISC_REQ_CONFIDENTIALITY | sspi.ISC_REQ_INTEGRITY | sspi.ISC_REQ_MUTUAL_AUTH
	domainSeparator		= `\`
	upnSeparator		= "@"
	shortDomainEnvVar	= "USERDOMAIN"
	maxUsername		= 256
	maxPassword		= 256
	maxDomain		= 15
)

func SSPIEnabled() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return true
}

type sspiNegotiator struct {
	principalName	string
	password	string
	reader		io.Reader
	writer		io.Writer
	host		string
	cred		*sspi.Credentials
	ctx		*negotiate.ClientContext
	desiredFlags	uint32
	requiredFlags	uint32
	complete	bool
}

func NewSSPINegotiator(principalName, password, host string, reader io.Reader) Negotiator {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &sspiNegotiator{principalName: principalName, password: password, reader: reader, writer: os.Stdout, host: host, desiredFlags: desiredFlags, requiredFlags: requiredFlags}
}
func (s *sspiNegotiator) Load() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(5).Info("Attempt to load SSPI")
	return nil
}
func (s *sspiNegotiator) InitSecContext(requestURL string, challengeToken []byte) (tokenToSend []byte, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer runtime.HandleCrash()
	if needsInit := s.cred == nil || s.ctx == nil; needsInit {
		logSSPI("Start SSPI flow: %s", requestURL)
		return s.initContext(requestURL)
	}
	klog.V(5).Info("Continue SSPI flow")
	return s.updateContext(challengeToken)
}
func (s *sspiNegotiator) IsComplete() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.complete
}
func (s *sspiNegotiator) Release() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer runtime.HandleCrash()
	klog.V(5).Info("Attempt to release SSPI")
	var errs []error
	if err := s.ctx.Release(); err != nil {
		logSSPI("SSPI context release failed: %v", err)
		errs = append(errs, err)
	}
	if err := s.cred.Release(); err != nil {
		logSSPI("SSPI credential release failed: %v", err)
		errs = append(errs, err)
	}
	return errors.Reduce(errors.NewAggregate(errs))
}
func (s *sspiNegotiator) initContext(requestURL string) (outputToken []byte, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cred, err := s.getUserCredentials()
	if err != nil {
		logSSPI("getUserCredentials failed: %v", err)
		return nil, err
	}
	s.cred = cred
	klog.V(5).Info("getUserCredentials successful")
	serviceName, err := getServiceName('/', requestURL)
	if err != nil {
		return nil, err
	}
	logSSPI("importing service name %s", serviceName)
	ctx, outputToken, err := negotiate.NewClientContextWithFlags(s.cred, serviceName, s.desiredFlags)
	if err != nil {
		logSSPI("NewClientContextWithFlags failed: %v", err)
		return nil, err
	}
	s.ctx = ctx
	klog.V(5).Info("NewClientContextWithFlags successful")
	return outputToken, nil
}
func (s *sspiNegotiator) getUserCredentials() (*sspi.Credentials, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(s.principalName) == 0 && len(s.password) > 0 {
		return nil, fmt.Errorf("username cannot be empty with non-empty password")
	}
	if len(s.principalName) > 0 {
		domain, username, err := s.getDomainAndUsername()
		if err != nil {
			return nil, err
		}
		password, err := s.getPassword(domain, username)
		if err != nil {
			return nil, err
		}
		logSSPI("Using AcquireUserCredentials because principalName is not empty, principalName=%s, username=%s, domain=%s", s.principalName, username, domain)
		cred, err := negotiate.AcquireUserCredentials(domain, username, password)
		if err != nil {
			logSSPI("AcquireUserCredentials failed: %v", err)
			return nil, err
		}
		klog.V(5).Info("AcquireUserCredentials successful")
		return cred, nil
	}
	klog.V(5).Info("Using AcquireCurrentUserCredentials because principalName is empty")
	return negotiate.AcquireCurrentUserCredentials()
}
func (s *sspiNegotiator) getDomainAndUsername() (domain, username string, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch {
	case strings.Contains(s.principalName, domainSeparator):
		data := strings.Split(s.principalName, domainSeparator)
		if len(data) != 2 || len(data[1]) == 0 {
			return "", "", fmt.Errorf(`invalid username %s, fully qualified user name format must have single backslash and non-empty user (ex: DOMAIN\Username)`, s.principalName)
		}
		domain = data[0]
		username = data[1]
	case strings.Contains(s.principalName, upnSeparator):
		username = s.principalName
	default:
		domain, _ = os.LookupEnv(shortDomainEnvVar)
		username = s.principalName
	}
	if domainLen, usernameLen := len(domain), len(username); domainLen > maxDomain || usernameLen > maxUsername {
		return "", "", fmt.Errorf("the maximum character lengths for user name and domain are %d and %d, respectively:\n"+"input username=%s username=%s,len=%d domain=%s,len=%d", maxUsername, maxDomain, s.principalName, username, usernameLen, domain, domainLen)
	}
	return domain, username, nil
}
func (s *sspiNegotiator) getPassword(domain, username string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	password := s.password
	if missingPassword := len(password) == 0; missingPassword {
		if hasDomain := len(domain) > 0; hasDomain {
			fmt.Fprintf(s.writer, "Authentication required for %s (%s)\n", s.host, domain)
		} else {
			fmt.Fprintf(s.writer, "Authentication required for %s\n", s.host)
		}
		fmt.Fprintf(s.writer, "Username: %s\n", username)
		password = term.PromptForPasswordString(s.reader, s.writer, "Password: ")
	}
	if passwordLen := len(password); passwordLen > maxPassword {
		return "", fmt.Errorf("the maximum character length for password is %d: password=<redacted>,len=%d", maxPassword, passwordLen)
	}
	return password, nil
}
func (s *sspiNegotiator) updateContext(challengeToken []byte) (outputToken []byte, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	authCompleted, outputToken, err := s.ctx.Update(challengeToken)
	if err != nil {
		logSSPI("ClientContext.Update failed: %v", err)
		return nil, err
	}
	s.complete = authCompleted
	logSSPI("ClientContext.Update successful, complete=%v", s.complete)
	if nonFatalErr := s.ctx.VerifyFlags(); nonFatalErr == nil {
		klog.V(5).Info("ClientContext.VerifyFlags successful")
	} else {
		logSSPI("ClientContext.VerifyFlags failed: %v", nonFatalErr)
		if fatalErr := s.ctx.VerifySelectiveFlags(s.requiredFlags); fatalErr != nil {
			logSSPI("ClientContext.VerifySelectiveFlags failed: %v", fatalErr)
			return nil, fatalErr
		}
		klog.V(5).Info("ClientContext.VerifySelectiveFlags successful")
	}
	return outputToken, nil
}
func logSSPI(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if klog.V(5) {
		for i, arg := range args {
			if errno, ok := arg.(syscall.Errno); ok {
				args[i] = fmt.Sprintf("%v, code=%#v", errno, errno)
			}
		}
		s := fmt.Sprintf(format, args...)
		klog.InfoDepth(1, s)
	}
}
