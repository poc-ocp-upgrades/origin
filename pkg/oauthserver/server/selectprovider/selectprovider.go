package selectprovider

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/openshift/origin/pkg/oauthserver/api"
	"github.com/openshift/origin/pkg/oauthserver/oauth/handlers"
	"html/template"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"net/http"
)

type SelectProviderRenderer interface {
	Render(redirectors []api.ProviderInfo, w http.ResponseWriter, req *http.Request)
}
type selectProvider struct {
	render            SelectProviderRenderer
	forceInterstitial bool
}

func NewSelectProvider(render SelectProviderRenderer, forceInterstitial bool) handlers.AuthenticationSelectionHandler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &selectProvider{render: render, forceInterstitial: forceInterstitial}
}

type ProviderData struct{ Providers []api.ProviderInfo }

func NewSelectProviderRenderer(customSelectProviderTemplateFile string) (*selectProviderTemplateRenderer, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := &selectProviderTemplateRenderer{}
	if len(customSelectProviderTemplateFile) > 0 {
		customTemplate, err := template.ParseFiles(customSelectProviderTemplateFile)
		if err != nil {
			return nil, err
		}
		r.selectProviderTemplate = customTemplate
	} else {
		r.selectProviderTemplate = defaultSelectProviderTemplate
	}
	return r, nil
}
func (s *selectProvider) SelectAuthentication(providers []api.ProviderInfo, w http.ResponseWriter, req *http.Request) (*api.ProviderInfo, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(providers) == 0 {
		return nil, false, nil
	}
	if len(providers) == 1 && !s.forceInterstitial {
		return &providers[0], false, nil
	}
	s.render.Render(providers, w, req)
	return nil, true, nil
}
func ValidateSelectProviderTemplate(templateContent []byte) []error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var allErrs []error
	template, err := template.New("selectProviderTemplateTest").Parse(string(templateContent))
	if err != nil {
		return append(allErrs, err)
	}
	providerData := ProviderData{Providers: []api.ProviderInfo{{Name: "provider_1", URL: "http://example.com/redirect_1/"}, {Name: "provider_2", URL: "http://example.com/redirect_2/"}}}
	var buffer bytes.Buffer
	err = template.Execute(&buffer, providerData)
	if err != nil {
		return append(allErrs, err)
	}
	output := buffer.Bytes()
	if !bytes.Contains(output, []byte(providerData.Providers[1].URL)) {
		allErrs = append(allErrs, errors.New("template must iterate over all {{.Providers}} and use the {{ .URL }} for each one"))
	}
	return allErrs
}

type selectProviderTemplateRenderer struct{ selectProviderTemplate *template.Template }

func (r selectProviderTemplateRenderer) Render(providers []api.ProviderInfo, w http.ResponseWriter, req *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	w.Header().Add("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := r.selectProviderTemplate.Execute(w, ProviderData{Providers: providers}); err != nil {
		utilruntime.HandleError(fmt.Errorf("unable to render select provider template: %v", err))
	}
}
