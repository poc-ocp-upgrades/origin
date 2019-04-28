package util

const (
	ProductOpenShift = `OpenShift`
)

func GetProductName(binaryName string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ProductOpenShift
}
func GetPlatformName(binaryName string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "OpenShift Application Platform"
}
func GetDistributionName(binaryName string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "OpenShift distribution of Kubernetes"
}
