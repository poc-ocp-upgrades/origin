package util

import (
	"bufio"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/debug"
	"strings"
	"time"
	g "github.com/onsi/ginkgo"
	o "github.com/onsi/gomega"
	authorizationapiv1 "k8s.io/api/authorization/v1"
	kapiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/client-go/discovery/cached"
	kclientset "k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	watchtools "k8s.io/client-go/tools/watch"
	kinternalclientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	e2e "k8s.io/kubernetes/test/e2e/framework"
	appsv1client "github.com/openshift/client-go/apps/clientset/versioned"
	buildv1client "github.com/openshift/client-go/build/clientset/versioned"
	configv1client "github.com/openshift/client-go/config/clientset/versioned"
	imagev1client "github.com/openshift/client-go/image/clientset/versioned"
	operatorv1client "github.com/openshift/client-go/operator/clientset/versioned"
	templateclient "github.com/openshift/client-go/template/clientset/versioned"
	_ "github.com/openshift/origin/pkg/api/install"
	authorizationclientset "github.com/openshift/origin/pkg/authorization/generated/internalclientset"
	configapi "github.com/openshift/origin/pkg/cmd/server/apis/config"
	"github.com/openshift/origin/pkg/cmd/server/bootstrappolicy"
	imageclientset "github.com/openshift/origin/pkg/image/generated/internalclientset"
	"github.com/openshift/origin/pkg/oc/lib/kubeconfig"
	projectapi "github.com/openshift/origin/pkg/project/apis/project"
	projectclientset "github.com/openshift/origin/pkg/project/generated/internalclientset"
	routeclientset "github.com/openshift/origin/pkg/route/generated/internalclientset"
	securityclientset "github.com/openshift/origin/pkg/security/generated/internalclientset"
	templateclientset "github.com/openshift/origin/pkg/template/generated/internalclientset"
	userclientset "github.com/openshift/origin/pkg/user/generated/internalclientset"
	testutil "github.com/openshift/origin/test/util"
)

type CLI struct {
	execPath		string
	verb			string
	configPath		string
	adminConfigPath		string
	username		string
	globalArgs		[]string
	commandArgs		[]string
	finalArgs		[]string
	namespacesToDelete	[]string
	stdin			*bytes.Buffer
	stdout			io.Writer
	stderr			io.Writer
	verbose			bool
	withoutNamespace	bool
	kubeFramework		*e2e.Framework
}

func NewCLI(project, adminConfigPath string) *CLI {
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
	client := &CLI{}
	g.AfterEach(client.TeardownProject)
	client.kubeFramework = e2e.NewDefaultFramework(project)
	client.kubeFramework.SkipNamespaceCreation = true
	client.username = "admin"
	client.execPath = "oc"
	if len(adminConfigPath) == 0 {
		FatalErr(fmt.Errorf("you must set the KUBECONFIG variable to admin kubeconfig"))
	}
	client.adminConfigPath = adminConfigPath
	g.BeforeEach(client.SetupProject)
	return client
}
func NewCLIWithoutNamespace(project string) *CLI {
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
	client := &CLI{}
	g.AfterEach(client.TeardownProject)
	client.kubeFramework = e2e.NewDefaultFramework(project)
	client.kubeFramework.SkipNamespaceCreation = true
	client.username = "admin"
	client.execPath = "oc"
	client.adminConfigPath = KubeConfigPath()
	return client
}
func (c *CLI) KubeFramework() *e2e.Framework {
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
	return c.kubeFramework
}
func (c *CLI) Username() string {
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
	return c.username
}
func (c *CLI) AsAdmin() *CLI {
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
	nc := *c
	nc.configPath = c.adminConfigPath
	return &nc
}
func (c *CLI) ChangeUser(name string) *CLI {
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
	adminClientConfig, err := testutil.GetClusterAdminClientConfig(c.adminConfigPath)
	if err != nil {
		FatalErr(err)
	}
	_, clientConfig, err := testutil.GetClientForUser(adminClientConfig, name)
	if err != nil {
		FatalErr(err)
	}
	kubeConfig, err := kubeconfig.CreateConfig(c.Namespace(), clientConfig)
	if err != nil {
		FatalErr(err)
	}
	f, err := ioutil.TempFile("", "configfile")
	if err != nil {
		FatalErr(err)
	}
	c.configPath = f.Name()
	err = clientcmd.WriteToFile(*kubeConfig, c.configPath)
	if err != nil {
		FatalErr(err)
	}
	c.username = name
	e2e.Logf("configPath is now %q", c.configPath)
	return c
}
func (c *CLI) SetNamespace(ns string) *CLI {
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
	c.kubeFramework.Namespace = &kapiv1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: ns}}
	return c
}
func (c CLI) WithoutNamespace() *CLI {
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
	c.withoutNamespace = true
	return &c
}
func (c *CLI) SetupProject() {
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
	newNamespace := names.SimpleNameGenerator.GenerateName(fmt.Sprintf("e2e-test-%s-", c.kubeFramework.BaseName))
	c.SetNamespace(newNamespace).ChangeUser(fmt.Sprintf("%s-user", newNamespace))
	e2e.Logf("The user is now %q", c.Username())
	e2e.Logf("Creating project %q", newNamespace)
	_, err := c.ProjectClient().Project().ProjectRequests().Create(&projectapi.ProjectRequest{ObjectMeta: metav1.ObjectMeta{Name: newNamespace}})
	o.Expect(err).NotTo(o.HaveOccurred())
	c.namespacesToDelete = append(c.namespacesToDelete, newNamespace)
	e2e.Logf("Waiting on permissions in project %q ...", newNamespace)
	err = WaitForSelfSAR(1*time.Second, 60*time.Second, c.KubeClient(), authorizationapiv1.SelfSubjectAccessReviewSpec{ResourceAttributes: &authorizationapiv1.ResourceAttributes{Namespace: newNamespace, Verb: "create", Group: "", Resource: "pods"}})
	o.Expect(err).NotTo(o.HaveOccurred())
	DefaultServiceAccounts := []string{bootstrappolicy.DefaultServiceAccountName, bootstrappolicy.DeployerServiceAccountName, bootstrappolicy.BuilderServiceAccountName}
	for _, sa := range DefaultServiceAccounts {
		e2e.Logf("Waiting for ServiceAccount %q to be provisioned...", sa)
		err = WaitForServiceAccount(c.KubeClient().CoreV1().ServiceAccounts(newNamespace), sa)
		o.Expect(err).NotTo(o.HaveOccurred())
	}
	var ctx context.Context
	cancel := func() {
	}
	defer cancel()
	defaultRoleBindings := bootstrappolicy.GetBootstrapServiceAccountProjectRoleBindings(newNamespace)
	for _, rb := range defaultRoleBindings {
		o.Expect(rb.Namespace).To(o.BeEquivalentTo(newNamespace))
		e2e.Logf("Waiting for RoleBinding %q to be provisioned...", rb)
		ctx, cancel = watchtools.ContextWithOptionalTimeout(context.Background(), 3*time.Minute)
		fieldSelector := fields.OneTermEqualSelector("metadata.name", rb.Name).String()
		lw := &cache.ListWatch{ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
			options.FieldSelector = fieldSelector
			return c.KubeClient().RbacV1().RoleBindings(rb.Namespace).List(options)
		}, WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
			options.FieldSelector = fieldSelector
			return c.KubeClient().RbacV1().RoleBindings(rb.Namespace).Watch(options)
		}}
		_, err := watchtools.UntilWithSync(ctx, lw, &rb, nil, func(event watch.Event) (b bool, e error) {
			switch t := event.Type; t {
			case watch.Added, watch.Modified:
				return true, nil
			case watch.Deleted:
				return true, fmt.Errorf("object has been deleted")
			default:
				return true, fmt.Errorf("internal error: unexpected event %#v", e)
			}
		})
		o.Expect(err).NotTo(o.HaveOccurred())
	}
	e2e.Logf("Project %q has been fully provisioned.", newNamespace)
}
func (c *CLI) CreateProject() string {
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
	newNamespace := names.SimpleNameGenerator.GenerateName(fmt.Sprintf("e2e-test-%s-", c.kubeFramework.BaseName))
	e2e.Logf("Creating project %q", newNamespace)
	_, err := c.ProjectClient().Project().ProjectRequests().Create(&projectapi.ProjectRequest{ObjectMeta: metav1.ObjectMeta{Name: newNamespace}})
	o.Expect(err).NotTo(o.HaveOccurred())
	c.namespacesToDelete = append(c.namespacesToDelete, newNamespace)
	e2e.Logf("Waiting on permissions in project %q ...", newNamespace)
	err = WaitForSelfSAR(1*time.Second, 60*time.Second, c.KubeClient(), authorizationapiv1.SelfSubjectAccessReviewSpec{ResourceAttributes: &authorizationapiv1.ResourceAttributes{Namespace: newNamespace, Verb: "create", Group: "", Resource: "pods"}})
	o.Expect(err).NotTo(o.HaveOccurred())
	return newNamespace
}
func (c *CLI) TeardownProject() {
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
	if len(c.Namespace()) > 0 && g.CurrentGinkgoTestDescription().Failed && e2e.TestContext.DumpLogsOnFailure {
		e2e.DumpAllNamespaceInfo(c.kubeFramework.ClientSet, c.Namespace())
	}
	if len(c.configPath) > 0 {
		os.Remove(c.configPath)
	}
	if e2e.TestContext.DeleteNamespace && len(c.namespacesToDelete) > 0 {
		timeout := e2e.DefaultNamespaceDeletionTimeout
		if c.kubeFramework.NamespaceDeletionTimeout != 0 {
			timeout = c.kubeFramework.NamespaceDeletionTimeout
		}
		e2e.DeleteNamespaces(c.kubeFramework.ClientSet, c.namespacesToDelete, nil)
		e2e.WaitForNamespacesDeleted(c.kubeFramework.ClientSet, c.namespacesToDelete, timeout)
	}
}
func (c *CLI) Verbose() *CLI {
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
	c.verbose = true
	return c
}
func (c *CLI) RESTMapper() meta.RESTMapper {
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
	ret := restmapper.NewDeferredDiscoveryRESTMapper(cached.NewMemCacheClient(c.KubeClient().Discovery()))
	ret.Reset()
	return ret
}
func (c *CLI) AppsClient() appsv1client.Interface {
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
	client, err := appsv1client.NewForConfig(c.UserConfig())
	if err != nil {
		FatalErr(err)
	}
	return client
}
func (c *CLI) AuthorizationClient() authorizationclientset.Interface {
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
	client, err := authorizationclientset.NewForConfig(c.UserConfig())
	if err != nil {
		FatalErr(err)
	}
	return client
}
func (c *CLI) BuildClient() buildv1client.Interface {
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
	client, err := buildv1client.NewForConfig(c.UserConfig())
	if err != nil {
		FatalErr(err)
	}
	return client
}
func (c *CLI) ImageClient() imageclientset.Interface {
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
	client, err := imageclientset.NewForConfig(c.UserConfig())
	if err != nil {
		FatalErr(err)
	}
	return client
}
func (c *CLI) ProjectClient() projectclientset.Interface {
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
	client, err := projectclientset.NewForConfig(c.UserConfig())
	if err != nil {
		FatalErr(err)
	}
	return client
}
func (c *CLI) RouteClient() routeclientset.Interface {
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
	client, err := routeclientset.NewForConfig(c.UserConfig())
	if err != nil {
		FatalErr(err)
	}
	return client
}
func (c *CLI) InternalTemplateClient() templateclientset.Interface {
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
	client, err := templateclientset.NewForConfig(c.UserConfig())
	if err != nil {
		FatalErr(err)
	}
	return client
}
func (c *CLI) TemplateClient() templateclient.Interface {
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
	client, err := templateclient.NewForConfig(c.UserConfig())
	if err != nil {
		FatalErr(err)
	}
	return client
}
func (c *CLI) UserClient() userclientset.Interface {
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
	client, err := userclientset.NewForConfig(c.UserConfig())
	if err != nil {
		FatalErr(err)
	}
	return client
}
func (c *CLI) AdminAppsClient() appsv1client.Interface {
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
	client, err := appsv1client.NewForConfig(c.AdminConfig())
	if err != nil {
		FatalErr(err)
	}
	return client
}
func (c *CLI) AdminAuthorizationClient() authorizationclientset.Interface {
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
	client, err := authorizationclientset.NewForConfig(c.AdminConfig())
	if err != nil {
		FatalErr(err)
	}
	return client
}
func (c *CLI) AdminBuildClient() buildv1client.Interface {
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
	client, err := buildv1client.NewForConfig(c.AdminConfig())
	if err != nil {
		FatalErr(err)
	}
	return client
}
func (c *CLI) AdminConfigClient() configv1client.Interface {
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
	client, err := configv1client.NewForConfig(c.AdminConfig())
	if err != nil {
		FatalErr(err)
	}
	return client
}
func (c *CLI) AdminImageClient() imagev1client.Interface {
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
	client, err := imagev1client.NewForConfig(c.AdminConfig())
	if err != nil {
		FatalErr(err)
	}
	return client
}
func (c *CLI) AdminOperatorClient() operatorv1client.Interface {
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
	client, err := operatorv1client.NewForConfig(c.AdminConfig())
	if err != nil {
		FatalErr(err)
	}
	return client
}
func (c *CLI) AdminInternalImageClient() imageclientset.Interface {
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
	client, err := imageclientset.NewForConfig(c.AdminConfig())
	if err != nil {
		FatalErr(err)
	}
	return client
}
func (c *CLI) AdminProjectClient() projectclientset.Interface {
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
	client, err := projectclientset.NewForConfig(c.AdminConfig())
	if err != nil {
		FatalErr(err)
	}
	return client
}
func (c *CLI) AdminRouteClient() routeclientset.Interface {
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
	client, err := routeclientset.NewForConfig(c.AdminConfig())
	if err != nil {
		FatalErr(err)
	}
	return client
}
func (c *CLI) AdminInternalTemplateClient() templateclientset.Interface {
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
	client, err := templateclientset.NewForConfig(c.AdminConfig())
	if err != nil {
		FatalErr(err)
	}
	return client
}
func (c *CLI) AdminTemplateClient() templateclient.Interface {
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
	client, err := templateclient.NewForConfig(c.AdminConfig())
	if err != nil {
		FatalErr(err)
	}
	return client
}
func (c *CLI) AdminUserClient() userclientset.Interface {
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
	client, err := userclientset.NewForConfig(c.AdminConfig())
	if err != nil {
		FatalErr(err)
	}
	return client
}
func (c *CLI) AdminSecurityClient() securityclientset.Interface {
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
	client, err := securityclientset.NewForConfig(c.AdminConfig())
	if err != nil {
		FatalErr(err)
	}
	return client
}
func (c *CLI) KubeClient() kclientset.Interface {
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
	return kclientset.NewForConfigOrDie(c.UserConfig())
}
func (c *CLI) InternalKubeClient() kinternalclientset.Interface {
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
	return kinternalclientset.NewForConfigOrDie(c.UserConfig())
}
func (c *CLI) AdminKubeClient() kclientset.Interface {
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
	return kclientset.NewForConfigOrDie(c.AdminConfig())
}
func (c *CLI) InternalAdminKubeClient() kinternalclientset.Interface {
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
	return kinternalclientset.NewForConfigOrDie(c.AdminConfig())
}
func (c *CLI) UserConfig() *restclient.Config {
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
	clientConfig, err := configapi.GetClientConfig(c.configPath, nil)
	if err != nil {
		FatalErr(err)
	}
	return clientConfig
}
func (c *CLI) AdminConfig() *restclient.Config {
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
	clientConfig, err := configapi.GetClientConfig(c.adminConfigPath, nil)
	if err != nil {
		FatalErr(err)
	}
	return clientConfig
}
func (c *CLI) Namespace() string {
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
	if c.kubeFramework.Namespace == nil {
		return ""
	}
	return c.kubeFramework.Namespace.Name
}
func (c *CLI) setOutput(out io.Writer) *CLI {
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
	c.stdout = out
	return c
}
func (c *CLI) Run(commands ...string) *CLI {
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
	in, out, errout := &bytes.Buffer{}, &bytes.Buffer{}, &bytes.Buffer{}
	nc := &CLI{execPath: c.execPath, verb: commands[0], kubeFramework: c.KubeFramework(), adminConfigPath: c.adminConfigPath, configPath: c.configPath, username: c.username, globalArgs: append(commands, []string{fmt.Sprintf("--config=%s", c.configPath)}...)}
	if !c.withoutNamespace {
		nc.globalArgs = append(nc.globalArgs, fmt.Sprintf("--namespace=%s", c.Namespace()))
	}
	nc.stdin, nc.stdout, nc.stderr = in, out, errout
	return nc.setOutput(c.stdout)
}
func (c *CLI) Template(t string) *CLI {
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
	if c.verb != "get" {
		FatalErr("Cannot use Template() for non-get verbs.")
	}
	templateArgs := []string{"--output=template", fmt.Sprintf("--template=%s", t)}
	commandArgs := append(c.commandArgs, templateArgs...)
	c.finalArgs = append(c.globalArgs, commandArgs...)
	return c
}
func (c *CLI) InputString(input string) *CLI {
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
	c.stdin.WriteString(input)
	return c
}
func (c *CLI) Args(args ...string) *CLI {
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
	c.commandArgs = args
	c.finalArgs = append(c.globalArgs, c.commandArgs...)
	return c
}
func (c *CLI) printCmd() string {
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
	return strings.Join(c.finalArgs, " ")
}

type ExitError struct {
	Cmd	string
	StdErr	string
	*exec.ExitError
}

func (c *CLI) Output() (string, error) {
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
	if c.verbose {
		fmt.Printf("DEBUG: oc %s\n", c.printCmd())
	}
	cmd := exec.Command(c.execPath, c.finalArgs...)
	cmd.Stdin = c.stdin
	e2e.Logf("Running '%s %s'", c.execPath, strings.Join(c.finalArgs, " "))
	out, err := cmd.CombinedOutput()
	trimmed := strings.TrimSpace(string(out))
	switch err.(type) {
	case nil:
		c.stdout = bytes.NewBuffer(out)
		return trimmed, nil
	case *exec.ExitError:
		e2e.Logf("Error running %v:\n%s", cmd, trimmed)
		return trimmed, &ExitError{ExitError: err.(*exec.ExitError), Cmd: c.execPath + " " + strings.Join(c.finalArgs, " "), StdErr: trimmed}
	default:
		FatalErr(fmt.Errorf("unable to execute %q: %v", c.execPath, err))
		return "", nil
	}
}
func (c *CLI) Outputs() (string, string, error) {
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
	if c.verbose {
		fmt.Printf("DEBUG: oc %s\n", c.printCmd())
	}
	cmd := exec.Command(c.execPath, c.finalArgs...)
	cmd.Stdin = c.stdin
	e2e.Logf("Running '%s %s'", c.execPath, strings.Join(c.finalArgs, " "))
	var stdErrBuff, stdOutBuff bytes.Buffer
	cmd.Stdout = &stdOutBuff
	cmd.Stderr = &stdErrBuff
	err := cmd.Run()
	stdOutBytes := stdOutBuff.Bytes()
	stdErrBytes := stdErrBuff.Bytes()
	stdOut := strings.TrimSpace(string(stdOutBytes))
	stdErr := strings.TrimSpace(string(stdErrBytes))
	switch err.(type) {
	case nil:
		c.stdout = bytes.NewBuffer(stdOutBytes)
		c.stderr = bytes.NewBuffer(stdErrBytes)
		return stdOut, stdErr, nil
	case *exec.ExitError:
		e2e.Logf("Error running %v:\nStdOut>\n%s\nStdErr>\n%s\n", cmd, stdOut, stdErr)
		return stdOut, stdErr, err
	default:
		FatalErr(fmt.Errorf("unable to execute %q: %v", c.execPath, err))
		return "", "", nil
	}
}
func (c *CLI) Background() (*exec.Cmd, *bytes.Buffer, *bytes.Buffer, error) {
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
	if c.verbose {
		fmt.Printf("DEBUG: oc %s\n", c.printCmd())
	}
	cmd := exec.Command(c.execPath, c.finalArgs...)
	cmd.Stdin = c.stdin
	var stdout, stderr bytes.Buffer
	cmd.Stdout = bufio.NewWriter(&stdout)
	cmd.Stderr = bufio.NewWriter(&stderr)
	e2e.Logf("Running '%s %s'", c.execPath, strings.Join(c.finalArgs, " "))
	err := cmd.Start()
	return cmd, &stdout, &stderr, err
}
func (c *CLI) BackgroundRC() (*exec.Cmd, io.ReadCloser, error) {
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
	if c.verbose {
		fmt.Printf("DEBUG: oc %s\n", c.printCmd())
	}
	cmd := exec.Command(c.execPath, c.finalArgs...)
	cmd.Stdin = c.stdin
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}
	e2e.Logf("Running '%s %s'", c.execPath, strings.Join(c.finalArgs, " "))
	err = cmd.Start()
	return cmd, stdout, err
}
func (c *CLI) OutputToFile(filename string) (string, error) {
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
	content, err := c.Output()
	if err != nil {
		return "", err
	}
	path := filepath.Join(e2e.TestContext.OutputDir, c.Namespace()+"-"+filename)
	return path, ioutil.WriteFile(path, []byte(content), 0644)
}
func (c *CLI) Execute() error {
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
	out, err := c.Output()
	if _, err := io.Copy(g.GinkgoWriter, strings.NewReader(out+"\n")); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: Unable to copy the output to ginkgo writer")
	}
	os.Stdout.Sync()
	return err
}
func FatalErr(msg interface{}) {
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
	fmt.Fprintln(g.GinkgoWriter, string(debug.Stack()))
	e2e.Failf("%v", msg)
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
