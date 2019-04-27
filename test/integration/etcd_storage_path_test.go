package integration

import (
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"
	"golang.org/x/net/context"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/diff"
	"k8s.io/apimachinery/pkg/util/sets"
	discocache "k8s.io/client-go/discovery/cached"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/util/flowcontrol"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	kapihelper "k8s.io/kubernetes/pkg/apis/core/helper"
	kclientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	etcddata "k8s.io/kubernetes/test/integration/etcd"
	"github.com/openshift/origin/pkg/cmd/server/etcd"
	testutil "github.com/openshift/origin/test/util"
	testserver "github.com/openshift/origin/test/util/server"
	etcdv3 "github.com/coreos/etcd/clientv3"
	"github.com/openshift/origin/pkg/api/install"
)

var openshiftEtcdStorageData = map[schema.GroupVersionResource]etcddata.StorageData{gvr("authorization.openshift.io", "v1", "roles"): {Stub: `{"metadata": {"name": "r1b1o2"}, "rules": [{"verbs": ["create"], "apiGroups": ["authorization.k8s.io"], "resources": ["selfsubjectaccessreviews"]}]}`, ExpectedEtcdPath: "kubernetes.io/roles/etcdstoragepathtestnamespace/r1b1o2", ExpectedGVK: gvkP("rbac.authorization.k8s.io", "v1", "Role")}, gvr("authorization.openshift.io", "v1", "clusterroles"): {Stub: `{"metadata": {"name": "cr1a1o2"}, "rules": [{"verbs": ["create"], "apiGroups": ["authorization.k8s.io"], "resources": ["selfsubjectaccessreviews"]}]}`, ExpectedEtcdPath: "kubernetes.io/clusterroles/cr1a1o2", ExpectedGVK: gvkP("rbac.authorization.k8s.io", "v1", "ClusterRole")}, gvr("authorization.openshift.io", "v1", "rolebindings"): {Stub: `{"metadata": {"name": "rb1a1o2"}, "subjects": [{"kind": "Group", "name": "system:authenticated"}], "roleRef": {"kind": "Role", "name": "r1a1"}}`, ExpectedEtcdPath: "kubernetes.io/rolebindings/etcdstoragepathtestnamespace/rb1a1o2", ExpectedGVK: gvkP("rbac.authorization.k8s.io", "v1", "RoleBinding")}, gvr("authorization.openshift.io", "v1", "clusterrolebindings"): {Stub: `{"metadata": {"name": "crb1a1o2"}, "subjects": [{"kind": "Group", "name": "system:authenticated"}], "roleRef": {"kind": "ClusterRole", "name": "cr1a1"}}`, ExpectedEtcdPath: "kubernetes.io/clusterrolebindings/crb1a1o2", ExpectedGVK: gvkP("rbac.authorization.k8s.io", "v1", "ClusterRoleBinding")}, gvr("authorization.openshift.io", "v1", "rolebindingrestrictions"): {Stub: `{"metadata": {"name": "rbrg"}, "spec": {"serviceaccountrestriction": {"serviceaccounts": [{"name": "sa"}]}}}`, ExpectedEtcdPath: "openshift.io/rolebindingrestrictions/etcdstoragepathtestnamespace/rbrg"}, gvr("build.openshift.io", "v1", "builds"): {Stub: `{"metadata": {"name": "build1g"}, "spec": {"source": {"dockerfile": "Dockerfile1"}, "strategy": {"dockerStrategy": {"noCache": true}}}}`, ExpectedEtcdPath: "openshift.io/builds/etcdstoragepathtestnamespace/build1g"}, gvr("build.openshift.io", "v1", "buildconfigs"): {Stub: `{"metadata": {"name": "bc1g"}, "spec": {"source": {"dockerfile": "Dockerfile0"}, "strategy": {"dockerStrategy": {"noCache": true}}}}`, ExpectedEtcdPath: "openshift.io/buildconfigs/etcdstoragepathtestnamespace/bc1g"}, gvr("apps.openshift.io", "v1", "deploymentconfigs"): {Stub: `{"metadata": {"name": "dc1g"}, "spec": {"selector": {"d": "c"}, "template": {"metadata": {"labels": {"d": "c"}}, "spec": {"containers": [{"image": "fedora:latest", "name": "container2"}]}}}}`, ExpectedEtcdPath: "openshift.io/deploymentconfigs/etcdstoragepathtestnamespace/dc1g"}, gvr("image.openshift.io", "v1", "imagestreams"): {Stub: `{"metadata": {"name": "is1g"}, "spec": {"dockerImageRepository": "docker"}}`, ExpectedEtcdPath: "openshift.io/imagestreams/etcdstoragepathtestnamespace/is1g"}, gvr("image.openshift.io", "v1", "images"): {Stub: `{"dockerImageReference": "fedora:latest", "metadata": {"name": "image1g"}}`, ExpectedEtcdPath: "openshift.io/images/image1g"}, gvr("oauth.openshift.io", "v1", "oauthclientauthorizations"): {Stub: `{"clientName": "system:serviceaccount:etcdstoragepathtestnamespace:clientg", "metadata": {"name": "user:system:serviceaccount:etcdstoragepathtestnamespace:clientg"}, "scopes": ["user:info"], "userName": "user", "userUID": "cannot be empty"}`, ExpectedEtcdPath: "openshift.io/oauth/clientauthorizations/user:system:serviceaccount:etcdstoragepathtestnamespace:clientg", Prerequisites: []etcddata.Prerequisite{{GvrData: gvr("", "v1", "serviceaccounts"), Stub: `{"metadata": {"annotations": {"serviceaccounts.openshift.io/oauth-redirecturi.foo": "http://bar"}, "name": "clientg"}}`}, {GvrData: gvr("", "v1", "secrets"), Stub: `{"metadata": {"annotations": {"kubernetes.io/service-account.name": "clientg"}, "generateName": "clientg"}, "type": "kubernetes.io/service-account-token"}`}}}, gvr("oauth.openshift.io", "v1", "oauthaccesstokens"): {Stub: `{"clientName": "client1g", "metadata": {"name": "tokenneedstobelongenoughelseitwontworkg"}, "userName": "user", "userUID": "cannot be empty"}`, ExpectedEtcdPath: "openshift.io/oauth/accesstokens/tokenneedstobelongenoughelseitwontworkg", Prerequisites: []etcddata.Prerequisite{{GvrData: gvr("oauth.openshift.io", "v1", "oauthclients"), Stub: `{"metadata": {"name": "client1g"}}`}}}, gvr("oauth.openshift.io", "v1", "oauthauthorizetokens"): {Stub: `{"clientName": "client0g", "metadata": {"name": "tokenneedstobelongenoughelseitwontworkg"}, "userName": "user", "userUID": "cannot be empty", "expiresIn": 86400}`, ExpectedEtcdPath: "openshift.io/oauth/authorizetokens/tokenneedstobelongenoughelseitwontworkg", Prerequisites: []etcddata.Prerequisite{{GvrData: gvr("oauth.openshift.io", "v1", "oauthclients"), Stub: `{"metadata": {"name": "client0g"}}`}}}, gvr("oauth.openshift.io", "v1", "oauthclients"): {Stub: `{"metadata": {"name": "clientg"}}`, ExpectedEtcdPath: "openshift.io/oauth/clients/clientg"}, gvr("project.openshift.io", "v1", "projects"): {Stub: `{"metadata": {"name": "namespace2g"}, "spec": {"finalizers": ["kubernetes", "openshift.io/origin"]}}`, ExpectedEtcdPath: "kubernetes.io/namespaces/namespace2g", ExpectedGVK: gvkP("", "v1", "Namespace")}, gvr("route.openshift.io", "v1", "routes"): {Stub: `{"metadata": {"name": "route1g"}, "spec": {"host": "hostname1", "to": {"name": "service1"}}}`, ExpectedEtcdPath: "openshift.io/routes/etcdstoragepathtestnamespace/route1g"}, gvr("security.openshift.io", "v1", "securitycontextconstraints"): {Stub: `{"allowPrivilegedContainer": true, "fsGroup": {"type": "RunAsAny"}, "metadata": {"name": "scc2"}, "runAsUser": {"type": "RunAsAny"}, "seLinuxContext": {"type": "MustRunAs"}, "supplementalGroups": {"type": "RunAsAny"}}`, ExpectedEtcdPath: "openshift.io/securitycontextconstraints/scc2"}, gvr("security.openshift.io", "v1", "rangeallocations"): {Stub: `{"metadata": {"name": "scc2"}}`, ExpectedEtcdPath: "openshift.io/rangeallocations/scc2"}, gvr("template.openshift.io", "v1", "templates"): {Stub: `{"message": "Jenkins template", "metadata": {"name": "template1g"}}`, ExpectedEtcdPath: "openshift.io/templates/etcdstoragepathtestnamespace/template1g"}, gvr("template.openshift.io", "v1", "templateinstances"): {Stub: `{"metadata": {"name": "templateinstance1"}, "spec": {"template": {"metadata": {"name": "template1", "namespace": "etcdstoragepathtestnamespace"}}, "requester": {"username": "test"}}}`, ExpectedEtcdPath: "openshift.io/templateinstances/etcdstoragepathtestnamespace/templateinstance1"}, gvr("template.openshift.io", "v1", "brokertemplateinstances"): {Stub: `{"metadata": {"name": "brokertemplateinstance1"}, "spec": {"templateInstance": {"kind": "TemplateInstance", "name": "templateinstance1", "namespace": "etcdstoragepathtestnamespace"}, "secret": {"kind": "Secret", "name": "secret1", "namespace": "etcdstoragepathtestnamespace"}}}`, ExpectedEtcdPath: "openshift.io/brokertemplateinstances/brokertemplateinstance1"}, gvr("user.openshift.io", "v1", "groups"): {Stub: `{"metadata": {"name": "groupg"}, "users": ["user1", "user2"]}`, ExpectedEtcdPath: "openshift.io/groups/groupg"}, gvr("user.openshift.io", "v1", "users"): {Stub: `{"fullName": "user1g", "metadata": {"name": "user1g"}}`, ExpectedEtcdPath: "openshift.io/users/user1g"}, gvr("user.openshift.io", "v1", "identities"): {Stub: `{"metadata": {"name": "github:user2g"}, "providerName": "github", "providerUserName": "user2g"}`, ExpectedEtcdPath: "openshift.io/useridentities/github:user2g"}}
var kindWhiteList = sets.NewString("ImageStreamTag", "UserIdentityMapping", "ClusterResourceQuota")

const testNamespace = "etcdstoragepathtestnamespace"

func TestEtcd3StoragePath(t *testing.T) {
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
	install.InstallInternalOpenShift(legacyscheme.Scheme)
	install.InstallInternalKube(legacyscheme.Scheme)
	masterConfig, err := testserver.DefaultMasterOptions()
	if err != nil {
		t.Fatalf("error getting master config: %#v", err)
	}
	masterConfig.AdmissionConfig.PluginOrderOverride = []string{"PodNodeSelector"}
	masterConfig.KubernetesMasterConfig.APIServerArguments = map[string][]string{"runtime-config": {"auditregistration.k8s.io/v1alpha1=true", "rbac.authorization.k8s.io/v1alpha1=true", "scheduling.k8s.io/v1alpha1=true", "settings.k8s.io/v1alpha1=true", "storage.k8s.io/v1alpha1=true", "batch/v2alpha1=true"}, "storage-media-type": {"application/json"}}
	masterConfig.KubernetesMasterConfig.APIServerArguments["disable-admission-plugins"] = append(masterConfig.KubernetesMasterConfig.APIServerArguments["disable-admission-plugins"], "ServiceAccount")
	_, err = testserver.StartConfiguredMasterAPI(masterConfig)
	if err != nil {
		t.Fatalf("error starting server: %v", err.Error())
	}
	etcdClient3, err := etcd.MakeEtcdClientV3(masterConfig.EtcdClientInfo)
	if err != nil {
		t.Fatal(err)
	}
	kubeConfigFile := masterConfig.MasterClients.OpenShiftLoopbackKubeConfig
	kubeConfig := testutil.GetClusterAdminClientConfigOrDie(kubeConfigFile)
	kubeConfig.QPS = 99999
	kubeConfig.Burst = 9999
	kubeClient := kclientset.NewForConfigOrDie(kubeConfig)
	etcddata.CreateTestCRDs(t, apiextensionsclientset.NewForConfigOrDie(kubeConfig), false, etcddata.GetCustomResourceDefinitionData()...)
	if err := testutil.WaitForClusterResourceQuotaCRDAvailable(kubeConfig); err != nil {
		t.Fatal(err)
	}
	mapper := restmapper.NewDeferredDiscoveryRESTMapper(discocache.NewMemCacheClient(kubeClient.Discovery()))
	mapper.Reset()
	client, err := newClient(*kubeConfig)
	if err != nil {
		t.Fatalf("error creating client: %#v", err)
	}
	if _, err := kubeClient.Core().Namespaces().Create(&kapi.Namespace{ObjectMeta: metav1.ObjectMeta{Name: testNamespace}}); err != nil {
		t.Fatalf("error creating test namespace: %#v", err)
	}
	etcdStorageData := etcddata.GetEtcdStorageData()
	delete(etcdStorageData, gvr("admissionregistration.k8s.io", "v1alpha1", "initializerconfigurations"))
	for gvr := range etcdStorageData {
		data := etcdStorageData[gvr]
		path := data.ExpectedEtcdPath
		if !strings.HasPrefix(path, "/registry/") {
			t.Fatalf("%s does not have expected Kube prefix, data=%#v", gvr.String(), data)
		}
		data.ExpectedEtcdPath = "kubernetes.io/" + path[10:]
		etcdStorageData[gvr] = data
	}
	for gvr, data := range openshiftEtcdStorageData {
		if _, ok := etcdStorageData[gvr]; ok {
			t.Errorf("%s exists in both Kube and OpenShift ETCD data, data=%#v", gvr.String(), data)
		}
		if len(gvr.Group) != 0 {
			isOpenShiftResource := gvr.Group == "openshift.io" || strings.HasSuffix(gvr.Group, ".openshift.io")
			if !isOpenShiftResource {
				t.Errorf("%s should be added in the upstream test, data=%#v", gvr.String(), data)
			}
		}
		etcdStorageData[gvr] = data
	}
	kindSeen := sets.NewString()
	pathSeen := map[string][]schema.GroupVersionResource{}
	etcdSeen := map[schema.GroupVersionResource]empty{}
	cohabitatingResources := map[string]map[schema.GroupVersionKind]empty{}
	serverResources, err := kubeClient.Discovery().ServerResources()
	if err != nil {
		t.Fatal(err)
	}
	resourcesToPersist := append(etcddata.GetResources(t, serverResources))
	for _, resourceToPersist := range resourcesToPersist {
		mapping := resourceToPersist.Mapping
		gvResource := mapping.Resource
		gvk := mapping.GroupVersionKind
		kind := gvk.Kind
		if kindWhiteList.Has(kind) {
			kindSeen.Insert(kind)
			continue
		}
		etcdSeen[gvResource] = empty{}
		testData, hasTest := etcdStorageData[gvResource]
		if !hasTest {
			t.Errorf("no test data for %v.  Please add a test for your new type to etcdStorageData.", gvk)
			continue
		}
		if len(testData.ExpectedEtcdPath) == 0 {
			t.Errorf("empty test data for %v", gvk)
			continue
		}
		shouldCreate := len(testData.Stub) != 0
		var input *metaObject
		if shouldCreate {
			if input, err = jsonToMetaObject(testData.Stub); err != nil || input.isEmpty() {
				t.Errorf("invalid test data for %v: %v", gvk, err)
				continue
			}
		}
		func() {
			all := &[]cleanupData{}
			defer func() {
				if !t.Failed() {
					if err := client.cleanup(all); err != nil {
						t.Errorf("failed to clean up etcd: %#v", err)
					}
				}
			}()
			if err := client.createPrerequisites(mapper, testNamespace, testData.Prerequisites, all); err != nil {
				t.Errorf("failed to create prerequisites for %v: %#v", gvk, err)
				return
			}
			if shouldCreate {
				if err := client.create(testData.Stub, testNamespace, mapping, all); err != nil {
					t.Errorf("failed to create stub for %v: %#v", gvk, err)
					return
				}
			}
			output, err := getFromEtcd(etcdClient3.KV, testData.ExpectedEtcdPath)
			if err != nil {
				t.Errorf("failed to get from etcd for %v: %#v", gvk, err)
				return
			}
			expectedGVK := gvk
			if testData.ExpectedGVK != nil {
				expectedGVK = *testData.ExpectedGVK
			}
			actualGVK := output.getGVK()
			if actualGVK != expectedGVK {
				t.Errorf("GVK for %v does not match, expected %s got %s", gvk, expectedGVK, actualGVK)
			}
			if !kapihelper.Semantic.DeepDerivative(input, output) {
				t.Errorf("Test stub for %v does not match: %s", gvk, diff.ObjectGoPrintDiff(input, output))
			}
			addGVKToEtcdBucket(cohabitatingResources, actualGVK, getEtcdBucket(testData.ExpectedEtcdPath))
			pathSeen[testData.ExpectedEtcdPath] = append(pathSeen[testData.ExpectedEtcdPath], gvResource)
		}()
	}
	if inEtcdData, inEtcdSeen := diffMaps(etcdStorageData, etcdSeen); len(inEtcdData) != 0 || len(inEtcdSeen) != 0 {
		t.Errorf("etcd data does not match the types we saw:\nin etcd data but not seen:\n%s\nseen but not in etcd data:\n%s", inEtcdData, inEtcdSeen)
	}
	if inKindData, inKindSeen := diffMaps(kindWhiteList, kindSeen); len(inKindData) != 0 || len(inKindSeen) != 0 {
		t.Errorf("kind whitelist data does not match the types we saw:\nin kind whitelist but not seen:\n%s\nseen but not in kind whitelist:\n%s", inKindData, inKindSeen)
	}
	for bucket, gvks := range cohabitatingResources {
		if len(gvks) != 1 {
			gvkStrings := []string{}
			for key := range gvks {
				gvkStrings = append(gvkStrings, keyStringer(key))
			}
			t.Errorf("cohabitating resources in etcd bucket %s have inconsistent GVKs\nyou may need to use DefaultStorageFactory.AddCohabitatingResources to sync the GVK of these resources:\n%s", bucket, gvkStrings)
		}
	}
	for path, gvrs := range pathSeen {
		if len(gvrs) != 1 {
			gvrStrings := []string{}
			for _, key := range gvrs {
				gvrStrings = append(gvrStrings, keyStringer(key))
			}
			t.Errorf("invalid test data, please ensure all expectedEtcdPath are unique, path %s has duplicate GVRs:\n%s", path, gvrStrings)
		}
	}
}
func addGVKToEtcdBucket(cohabitatingResources map[string]map[schema.GroupVersionKind]empty, gvk schema.GroupVersionKind, bucket string) {
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
	if cohabitatingResources[bucket] == nil {
		cohabitatingResources[bucket] = map[schema.GroupVersionKind]empty{}
	}
	cohabitatingResources[bucket][gvk] = empty{}
}
func getEtcdBucket(path string) string {
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
	idx := strings.LastIndex(path, "/")
	if idx == -1 {
		panic("path with no slashes " + path)
	}
	bucket := path[:idx]
	if len(bucket) == 0 {
		panic("invalid bucket for path " + path)
	}
	return bucket
}

type metaObject struct {
	Kind		string	`json:"kind,omitempty"`
	APIVersion	string	`json:"apiVersion,omitempty"`
	Metadata	struct {
		Name		string	`json:"name,omitempty"`
		Namespace	string	`json:"namespace,omitempty"`
	}	`json:"metadata,omitempty"`
}

func (obj *metaObject) getGVK() schema.GroupVersionKind {
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
	return schema.FromAPIVersionAndKind(obj.APIVersion, obj.Kind)
}
func (obj *metaObject) isEmpty() bool {
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
	return obj == nil || *obj == metaObject{}
}
func (obj *metaObject) GetObjectKind() schema.ObjectKind {
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
	return schema.EmptyObjectKind
}
func (obj *metaObject) DeepCopyObject() runtime.Object {
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
	out := new(metaObject)
	out.Kind = obj.Kind
	out.APIVersion = obj.APIVersion
	out.Metadata.Name = obj.Metadata.Name
	out.Metadata.Namespace = obj.Metadata.Namespace
	return out
}

type empty struct{}
type cleanupData struct {
	obj	runtime.Object
	mapping	*meta.RESTMapping
}

func gvr(g, v, r string) schema.GroupVersionResource {
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
	return schema.GroupVersionResource{Group: g, Version: v, Resource: r}
}
func gvkP(g, v, k string) *schema.GroupVersionKind {
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
	return &schema.GroupVersionKind{Group: g, Version: v, Kind: k}
}
func jsonToMetaObject(stub string) (*metaObject, error) {
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
	obj := &metaObject{}
	if err := json.Unmarshal([]byte(stub), &obj); err != nil {
		return nil, err
	}
	obj.Kind = ""
	obj.APIVersion = ""
	return obj, nil
}
func keyStringer(i interface{}) string {
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
	base := "\n\t"
	switch key := i.(type) {
	case string:
		return base + key
	case schema.GroupVersionResource:
		return base + key.String()
	case schema.GroupVersionKind:
		return base + key.String()
	default:
		panic("unexpected type")
	}
}

type allClient struct {
	client	*http.Client
	config	*restclient.Config
	backoff	restclient.BackoffManager
}

func (c *allClient) verb(verb string, gvk schema.GroupVersionKind) (*restclient.Request, error) {
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
	apiPath := "/apis"
	switch {
	case gvk.Group == kapi.GroupName:
		apiPath = "/api"
	}
	baseURL, versionedAPIPath, err := restclient.DefaultServerURL(c.config.Host, apiPath, gvk.GroupVersion(), true)
	if err != nil {
		return nil, err
	}
	contentConfig := c.config.ContentConfig
	gv := gvk.GroupVersion()
	contentConfig.GroupVersion = &gv
	serializers, err := createSerializers(contentConfig)
	if err != nil {
		return nil, err
	}
	return restclient.NewRequest(c.client, verb, baseURL, versionedAPIPath, contentConfig, *serializers, c.backoff, c.config.RateLimiter, 0), nil
}
func (c *allClient) create(stub, ns string, mapping *meta.RESTMapping, all *[]cleanupData) error {
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
	req, err := c.verb("POST", mapping.GroupVersionKind)
	if err != nil {
		return err
	}
	namespaced := mapping.Scope.Name() == meta.RESTScopeNameNamespace
	output, err := req.NamespaceIfScoped(ns, namespaced).Resource(mapping.Resource.Resource).Body(strings.NewReader(stub)).Do().Get()
	if err != nil {
		if runtime.IsNotRegisteredError(err) {
			return nil
		}
		return err
	}
	*all = append(*all, cleanupData{output, mapping})
	return nil
}
func (c *allClient) destroy(obj runtime.Object, mapping *meta.RESTMapping) error {
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
	req, err := c.verb("DELETE", mapping.GroupVersionKind)
	if err != nil {
		return err
	}
	namespaced := mapping.Scope.Name() == meta.RESTScopeNameNamespace
	metadata, err := meta.Accessor(obj)
	if err != nil {
		return err
	}
	return req.NamespaceIfScoped(metadata.GetNamespace(), namespaced).Resource(mapping.Resource.Resource).Name(metadata.GetName()).Do().Error()
}
func (c *allClient) cleanup(all *[]cleanupData) error {
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
	for i := len(*all) - 1; i >= 0; i-- {
		obj := (*all)[i].obj
		mapping := (*all)[i].mapping
		if err := c.destroy(obj, mapping); err != nil {
			return err
		}
	}
	return nil
}
func (c *allClient) createPrerequisites(mapper meta.RESTMapper, ns string, prerequisites []etcddata.Prerequisite, all *[]cleanupData) error {
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
	for _, prerequisite := range prerequisites {
		gvk, err := mapper.KindFor(prerequisite.GvrData)
		if err != nil {
			return err
		}
		mapping, err := mapper.RESTMapping(gvk.GroupKind(), gvk.Version)
		if err != nil {
			return err
		}
		if err := c.create(prerequisite.Stub, ns, mapping, all); err != nil {
			return err
		}
	}
	return nil
}
func newClient(config restclient.Config) (*allClient, error) {
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
	config.ContentConfig.NegotiatedSerializer = legacyscheme.Codecs
	config.ContentConfig.ContentType = "application/json"
	config.Timeout = 30 * time.Second
	config.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(9999, 9999)
	transport, err := restclient.TransportFor(&config)
	if err != nil {
		return nil, err
	}
	client := &http.Client{Transport: transport, Timeout: config.Timeout}
	backoff := &restclient.URLBackoff{Backoff: flowcontrol.NewBackOff(1*time.Second, 10*time.Second)}
	return &allClient{client: client, config: &config, backoff: backoff}, nil
}
func createSerializers(config restclient.ContentConfig) (*restclient.Serializers, error) {
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
	mediaTypes := config.NegotiatedSerializer.SupportedMediaTypes()
	contentType := config.ContentType
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return nil, fmt.Errorf("the content type specified in the client configuration is not recognized: %v", err)
	}
	info, ok := runtime.SerializerInfoForMediaType(mediaTypes, mediaType)
	if !ok {
		if len(contentType) != 0 || len(mediaTypes) == 0 {
			return nil, fmt.Errorf("no serializers registered for %s", contentType)
		}
		info = mediaTypes[0]
	}
	internalGV := schema.GroupVersions{{Group: config.GroupVersion.Group, Version: runtime.APIVersionInternal}, {Group: "", Version: runtime.APIVersionInternal}}
	s := &restclient.Serializers{Encoder: config.NegotiatedSerializer.EncoderForVersion(info.Serializer, *config.GroupVersion), Decoder: config.NegotiatedSerializer.DecoderToVersion(info.Serializer, internalGV), RenegotiatedDecoder: func(contentType string, params map[string]string) (runtime.Decoder, error) {
		info, ok := runtime.SerializerInfoForMediaType(mediaTypes, contentType)
		if !ok {
			return nil, fmt.Errorf("serializer for %s not registered", contentType)
		}
		return config.NegotiatedSerializer.DecoderToVersion(info.Serializer, internalGV), nil
	}}
	if info.StreamSerializer != nil {
		s.StreamingSerializer = info.StreamSerializer.Serializer
		s.Framer = info.StreamSerializer.Framer
	}
	return s, nil
}
func getFromEtcd(kv etcdv3.KV, path string) (*metaObject, error) {
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
	response, err := kv.Get(context.Background(), "/"+path, etcdv3.WithSerializable())
	if err != nil {
		return nil, err
	}
	if len(response.Kvs) == 0 {
		return nil, fmt.Errorf("no keys found for %q", "/"+path)
	}
	into := &metaObject{}
	if _, _, err := legacyscheme.Codecs.UniversalDeserializer().Decode(response.Kvs[0].Value, nil, into); err != nil {
		return nil, err
	}
	return into, nil
}
func diffMaps(a, b interface{}) ([]string, []string) {
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
	inA := diffMapKeys(a, b, keyStringer)
	inB := diffMapKeys(b, a, keyStringer)
	return inA, inB
}
func diffMapKeys(a, b interface{}, stringer func(interface{}) string) []string {
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
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)
	ret := []string{}
	for _, ka := range av.MapKeys() {
		kat := ka.Interface()
		found := false
		for _, kb := range bv.MapKeys() {
			kbt := kb.Interface()
			if kat == kbt {
				found = true
				break
			}
		}
		if !found {
			ret = append(ret, stringer(kat))
		}
	}
	return ret
}
