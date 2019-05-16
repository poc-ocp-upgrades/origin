package util

const (
	ProductOpenShift = `OpenShift`
)

func GetProductName(binaryName string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ProductOpenShift
}
func GetPlatformName(binaryName string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "OpenShift Application Platform"
}
func GetDistributionName(binaryName string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "OpenShift distribution of Kubernetes"
}
