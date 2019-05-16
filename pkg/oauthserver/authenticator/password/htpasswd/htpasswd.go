package htpasswd

import (
	"bufio"
	"context"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	goformat "fmt"
	authapi "github.com/openshift/origin/pkg/oauthserver/api"
	"github.com/openshift/origin/pkg/oauthserver/authenticator/identitymapper"
	"golang.org/x/crypto/bcrypt"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/klog"
	"os"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

type Authenticator struct {
	providerName string
	file         string
	fileInfo     os.FileInfo
	mapper       authapi.UserIdentityMapper
	usernames    map[string]string
}

func New(providerName string, file string, mapper authapi.UserIdentityMapper) (authenticator.Password, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	auth := &Authenticator{providerName: providerName, file: file, mapper: mapper}
	if err := auth.loadIfNeeded(); err != nil {
		return nil, err
	}
	return auth, nil
}
func (a *Authenticator) AuthenticatePassword(ctx context.Context, username, password string) (*authenticator.Response, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := a.loadIfNeeded(); err != nil {
		return nil, false, err
	}
	if len(username) > 255 {
		username = username[:255]
	}
	if strings.Contains(username, ":") {
		return nil, false, errors.New("Usernames may not contain : characters")
	}
	hash, ok := a.usernames[username]
	if !ok {
		return nil, false, nil
	}
	if ok, err := testPassword(password, hash); !ok || err != nil {
		return nil, false, err
	}
	identity := authapi.NewDefaultUserIdentityInfo(a.providerName, username)
	return identitymapper.ResponseFor(a.mapper, identity)
}
func (a *Authenticator) load() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	file, err := os.Open(a.file)
	if err != nil {
		return err
	}
	defer file.Close()
	newusernames := map[string]string{}
	warnedusernames := map[string]bool{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			klog.Warningf("Ignoring malformed htpasswd line: %s", line)
			continue
		}
		username := parts[0]
		password := parts[1]
		if _, duplicate := newusernames[username]; duplicate {
			if _, warned := warnedusernames[username]; !warned {
				warnedusernames[username] = true
				klog.Warningf("%s contains multiple passwords for user '%s'. The last one specified will be used.", a.file, username)
			}
		}
		newusernames[username] = password
	}
	a.usernames = newusernames
	return nil
}
func (a *Authenticator) loadIfNeeded() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	info, err := os.Stat(a.file)
	if err != nil {
		return err
	}
	if a.fileInfo == nil || a.fileInfo.ModTime() != info.ModTime() {
		klog.V(4).Infof("Loading htpasswd file %s...", a.file)
		loadingErr := a.load()
		if loadingErr != nil {
			return err
		}
		a.fileInfo = info
		return nil
	}
	return nil
}
func testPassword(password, hash string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch {
	case strings.HasPrefix(hash, "$apr1$"):
		return testMD5Password(password, hash)
	case strings.HasPrefix(hash, "$2y$") || strings.HasPrefix(hash, "$2a$"):
		return testBCryptPassword(password, hash)
	case strings.HasPrefix(hash, "{SHA}"):
		return testSHAPassword(password, hash[5:])
	case len(hash) == 13:
		return testCryptPassword(password, hash)
	default:
		return false, errors.New("Unrecognized hash type")
	}
}
func testSHAPassword(password, hash string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(hash) == 0 {
		return false, errors.New("Invalid SHA hash")
	}
	shasum := sha1.Sum([]byte(password))
	base64shasum := base64.StdEncoding.EncodeToString(shasum[:])
	match := hash == base64shasum
	return match, nil
}
func testBCryptPassword(password, hash string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
func testMD5Password(password, hash string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	parts := strings.Split(hash, "$")
	if len(parts) != 4 {
		return false, errors.New("Malformed MD5 hash")
	}
	salt := parts[2]
	if len(salt) == 0 {
		return false, errors.New("Malformed MD5 hash: missing salt")
	}
	if len(salt) > 8 {
		salt = salt[:8]
	}
	md5hash := parts[3]
	if len(md5hash) == 0 {
		return false, errors.New("Malformed MD5 hash: missing hash")
	}
	testhash := string(aprMD5([]byte(password), []byte(salt)))
	match := testhash == hash
	return match, nil
}
func testCryptPassword(password, hash string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false, errors.New("crypt password hashes are not supported")
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
