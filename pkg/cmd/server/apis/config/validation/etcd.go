package validation

import (
	"fmt"
	"strings"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"github.com/openshift/origin/pkg/cmd/server/apis/config"
	"github.com/openshift/origin/pkg/cmd/server/apis/config/validation/common"
)

func ValidateEtcdConnectionInfo(config config.EtcdConnectionInfo, server *config.EtcdConfig, fldPath *field.Path) field.ErrorList {
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
	allErrs := field.ErrorList{}
	if len(config.URLs) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("urls"), ""))
	}
	for i, u := range config.URLs {
		_, urlErrs := common.ValidateURL(u, fldPath.Child("urls").Index(i))
		if len(urlErrs) > 0 {
			allErrs = append(allErrs, urlErrs...)
		}
	}
	if len(config.CA) > 0 {
		allErrs = append(allErrs, common.ValidateFile(config.CA, fldPath.Child("ca"))...)
	}
	allErrs = append(allErrs, common.ValidateCertInfo(config.ClientCert, false, fldPath)...)
	if server != nil {
		var builtInAddress = fmt.Sprintf("https://%s", server.Address)
		if len(server.ServingInfo.ClientCA) > 0 {
			if len(config.ClientCert.CertFile) == 0 {
				allErrs = append(allErrs, field.Required(fldPath.Child("certFile"), "A client certificate must be provided for this etcd server"))
			}
		}
		clientURLs := sets.NewString(config.URLs...)
		if !clientURLs.Has(builtInAddress) {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("urls"), strings.Join(clientURLs.List(), ","), fmt.Sprintf("must include the etcd address %s", builtInAddress)))
		}
	}
	return allErrs
}
func ValidateEtcdConfig(config *config.EtcdConfig, fldPath *field.Path) common.ValidationResults {
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
	validationResults := common.ValidationResults{}
	servingInfoPath := fldPath.Child("servingInfo")
	validationResults.Append(common.ValidateServingInfo(config.ServingInfo, true, servingInfoPath))
	if config.ServingInfo.BindNetwork == "tcp6" {
		validationResults.AddErrors(field.Invalid(servingInfoPath.Child("bindNetwork"), config.ServingInfo.BindNetwork, "tcp6 is not a valid bindNetwork for etcd, must be tcp or tcp4"))
	}
	if len(config.ServingInfo.NamedCertificates) > 0 {
		validationResults.AddErrors(field.Invalid(servingInfoPath.Child("namedCertificates"), "<not shown>", "namedCertificates are not supported for etcd"))
	}
	peerServingInfoPath := fldPath.Child("peerServingInfo")
	validationResults.Append(common.ValidateServingInfo(config.PeerServingInfo, true, peerServingInfoPath))
	if config.ServingInfo.BindNetwork == "tcp6" {
		validationResults.AddErrors(field.Invalid(peerServingInfoPath.Child("bindNetwork"), config.ServingInfo.BindNetwork, "tcp6 is not a valid bindNetwork for etcd peers, must be tcp or tcp4"))
	}
	if len(config.ServingInfo.NamedCertificates) > 0 {
		validationResults.AddErrors(field.Invalid(peerServingInfoPath.Child("namedCertificates"), "<not shown>", "namedCertificates are not supported for etcd"))
	}
	validationResults.AddErrors(common.ValidateHostPort(config.Address, fldPath.Child("address"))...)
	validationResults.AddErrors(common.ValidateHostPort(config.PeerAddress, fldPath.Child("peerAddress"))...)
	if len(config.StorageDir) == 0 {
		validationResults.AddErrors(field.Required(fldPath.Child("storageDirectory"), ""))
	}
	return validationResults
}
