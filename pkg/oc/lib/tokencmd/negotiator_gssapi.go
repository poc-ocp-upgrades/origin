package tokencmd

import (
	"errors"
	"runtime"
	"sync"
	"time"
	"github.com/apcera/gssapi"
	"k8s.io/klog"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
)

func GSSAPIEnabled() bool {
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

type gssapiNegotiator struct {
	lib		*gssapi.Lib
	loadError	error
	loadOnce	sync.Once
	name		*gssapi.Name
	ctx		*gssapi.CtxId
	flags		uint32
	principalName	string
	cred		*gssapi.CredId
	complete	bool
}

func NewGSSAPINegotiator(principalName string) Negotiator {
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
	return &gssapiNegotiator{principalName: principalName}
}
func (g *gssapiNegotiator) Load() error {
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
	_, err := g.loadLib()
	return err
}
func (g *gssapiNegotiator) InitSecContext(requestURL string, challengeToken []byte) (tokenToSend []byte, err error) {
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
	lib, err := g.loadLib()
	if err != nil {
		return nil, err
	}
	if g.ctx == nil {
		if len(g.principalName) > 0 {
			klog.V(5).Infof("acquiring credentials for principal name %s", g.principalName)
			credBuffer, err := lib.MakeBufferString(g.principalName)
			if err != nil {
				return nil, convertGSSAPIError(err)
			}
			defer credBuffer.Release()
			credName, err := credBuffer.Name(lib.GSS_KRB5_NT_PRINCIPAL_NAME)
			if err != nil {
				return nil, convertGSSAPIError(err)
			}
			defer credName.Release()
			cred, _, _, err := lib.AcquireCred(credName, time.Duration(0), lib.GSS_C_NO_OID_SET, gssapi.GSS_C_INITIATE)
			if err != nil {
				klog.V(5).Infof("AcquireCred returned error: %v", err)
				return nil, convertGSSAPIError(err)
			}
			g.cred = cred
		} else {
			g.cred = lib.GSS_C_NO_CREDENTIAL
		}
		serviceName, err := getServiceName('@', requestURL)
		if err != nil {
			return nil, err
		}
		klog.V(5).Infof("importing service name %s", serviceName)
		nameBuf, err := lib.MakeBufferString(serviceName)
		if err != nil {
			return nil, convertGSSAPIError(err)
		}
		defer nameBuf.Release()
		name, err := nameBuf.Name(lib.GSS_C_NT_HOSTBASED_SERVICE)
		if err != nil {
			return nil, convertGSSAPIError(err)
		}
		g.name = name
		g.ctx = lib.GSS_C_NO_CONTEXT
	}
	incomingTokenBuffer, err := lib.MakeBufferBytes(challengeToken)
	if err != nil {
		return nil, convertGSSAPIError(err)
	}
	defer incomingTokenBuffer.Release()
	var outgoingToken *gssapi.Buffer
	g.ctx, _, outgoingToken, _, _, err = lib.InitSecContext(g.cred, g.ctx, g.name, lib.GSS_C_NO_OID, g.flags, time.Duration(0), lib.GSS_C_NO_CHANNEL_BINDINGS, incomingTokenBuffer)
	defer outgoingToken.Release()
	switch err {
	case nil:
		klog.V(5).Infof("InitSecContext returned GSS_S_COMPLETE")
		g.complete = true
		return outgoingToken.Bytes(), nil
	case gssapi.ErrContinueNeeded:
		klog.V(5).Infof("InitSecContext returned GSS_S_CONTINUE_NEEDED")
		g.complete = false
		return outgoingToken.Bytes(), nil
	default:
		klog.V(5).Infof("InitSecContext returned error: %v", err)
		return nil, convertGSSAPIError(err)
	}
}
func (g *gssapiNegotiator) IsComplete() bool {
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
	return g.complete
}
func (g *gssapiNegotiator) Release() error {
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
	var errs []error
	if err := g.name.Release(); err != nil {
		errs = append(errs, convertGSSAPIError(err))
	}
	if err := g.ctx.Release(); err != nil {
		errs = append(errs, convertGSSAPIError(err))
	}
	if err := g.cred.Release(); err != nil {
		errs = append(errs, convertGSSAPIError(err))
	}
	if err := g.lib.Unload(); err != nil {
		errs = append(errs, convertGSSAPIError(err))
	}
	return utilerrors.NewAggregate(errs)
}
func (g *gssapiNegotiator) loadLib() (*gssapi.Lib, error) {
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
	g.loadOnce.Do(func() {
		klog.V(5).Infof("loading gssapi")
		var libPaths []string
		switch runtime.GOOS {
		case "darwin":
			libPaths = []string{"libgssapi_krb5.dylib"}
		case "linux":
			libPaths = []string{"libgssapi_krb5.so.2", "libgssapi.so.3"}
		default:
			libPaths = []string{""}
		}
		var loadErrors []error
		for _, libPath := range libPaths {
			lib, loadError := gssapi.Load(&gssapi.Options{LibPath: libPath})
			if loadError == nil {
				klog.V(5).Infof("loaded gssapi %s", libPath)
				g.lib = lib
				return
			}
			klog.V(5).Infof("%v", loadError)
			loadErrors = append(loadErrors, convertGSSAPIError(loadError))
		}
		g.loadError = utilerrors.NewAggregate(loadErrors)
	})
	return g.lib, g.loadError
}
func convertGSSAPIError(err error) error {
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
	return errors.New(err.Error())
}
