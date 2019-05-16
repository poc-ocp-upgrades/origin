package install

import (
	goformat "fmt"
	"k8s.io/apimachinery/pkg/runtime/schema"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newIgnoredResources := map[schema.GroupResource]struct{}{{Group: "extensions", Resource: "replicationcontrollers"}: {}, {Group: "extensions", Resource: "networkpolicies"}: {}, {Group: "", Resource: "bindings"}: {}, {Group: "", Resource: "componentstatuses"}: {}, {Group: "", Resource: "events"}: {}, {Group: "authentication.k8s.io", Resource: "tokenreviews"}: {}, {Group: "authorization.k8s.io", Resource: "subjectaccessreviews"}: {}, {Group: "authorization.k8s.io", Resource: "selfsubjectaccessreviews"}: {}, {Group: "authorization.k8s.io", Resource: "localsubjectaccessreviews"}: {}, {Group: "authorization.k8s.io", Resource: "selfsubjectrulesreviews"}: {}, {Group: "authorization.openshift.io", Resource: "selfsubjectaccessreviews"}: {}, {Group: "authorization.openshift.io", Resource: "subjectaccessreviews"}: {}, {Group: "authorization.openshift.io", Resource: "localsubjectaccessreviews"}: {}, {Group: "authorization.openshift.io", Resource: "resourceaccessreviews"}: {}, {Group: "authorization.openshift.io", Resource: "localresourceaccessreviews"}: {}, {Group: "authorization.openshift.io", Resource: "selfsubjectrulesreviews"}: {}, {Group: "authorization.openshift.io", Resource: "subjectrulesreviews"}: {}, {Group: "authorization.openshift.io", Resource: "roles"}: {}, {Group: "authorization.openshift.io", Resource: "rolebindings"}: {}, {Group: "authorization.openshift.io", Resource: "clusterroles"}: {}, {Group: "authorization.openshift.io", Resource: "clusterrolebindings"}: {}, {Group: "apiregistration.k8s.io", Resource: "apiservices"}: {}, {Group: "apiextensions.k8s.io", Resource: "customresourcedefinitions"}: {}}
	for k, v := range newIgnoredResources {
		ignoredResources[k] = v
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
