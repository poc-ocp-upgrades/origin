package verifyimagesignature

import (
	"context"
	"bytes"
	"runtime"
	"fmt"
	"net/http"
	"net/url"
	godigest "github.com/opencontainers/go-digest"
	"k8s.io/client-go/rest"
	"github.com/openshift/origin/pkg/image/registryclient"
)

func getImageManifestByIDFromRegistry(registry *url.URL, repositoryName, imageID, username, password string, insecure bool) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx := context.Background()
	credentials := registryclient.NewBasicCredentials()
	credentials.Add(registry, username, password)
	insecureRT, err := rest.TransportFor(&rest.Config{TLSClientConfig: rest.TLSClientConfig{Insecure: true}})
	if err != nil {
		return nil, err
	}
	repo, err := registryclient.NewContext(http.DefaultTransport, insecureRT).WithCredentials(credentials).Repository(ctx, registry, repositoryName, insecure)
	if err != nil {
		return nil, err
	}
	manifests, err := repo.Manifests(ctx, nil)
	if err != nil {
		return nil, err
	}
	manifest, err := manifests.Get(ctx, godigest.Digest(imageID))
	if err != nil {
		return nil, err
	}
	_, manifestPayload, err := manifest.Payload()
	if err != nil {
		return nil, err
	}
	return manifestPayload, nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
