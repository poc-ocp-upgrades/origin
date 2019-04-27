package login

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	restclient "k8s.io/client-go/rest"
	kclientcmd "k8s.io/client-go/tools/clientcmd"
	kclientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	kterm "k8s.io/kubernetes/pkg/kubectl/util/term"
	userv1 "github.com/openshift/api/user/v1"
	projectv1typedclient "github.com/openshift/client-go/project/clientset/versioned/typed/project/v1"
	"github.com/openshift/origin/pkg/client/config"
	cmdutil "github.com/openshift/origin/pkg/cmd/util"
	"github.com/openshift/origin/pkg/cmd/util/term"
	"github.com/openshift/origin/pkg/oc/lib/errors"
	cliconfig "github.com/openshift/origin/pkg/oc/lib/kubeconfig"
	"github.com/openshift/origin/pkg/oc/lib/tokencmd"
	"github.com/openshift/origin/pkg/oc/util/project"
	loginutil "github.com/openshift/origin/pkg/oc/util/project"
)

const defaultClusterURL = "https://localhost:8443"
const projectsItemsSuppressThreshold = 50

type LoginOptions struct {
	Server			string
	CAFile			string
	InsecureTLS		bool
	Username		string
	Password		string
	Project			string
	StartingKubeConfig	*kclientcmdapi.Config
	DefaultNamespace	string
	Config			*restclient.Config
	CertFile		string
	KeyFile			string
	Token			string
	PathOptions		*kclientcmd.PathOptions
	CommandName		string
	RequestTimeout		time.Duration
	genericclioptions.IOStreams
}

func NewLoginOptions(streams genericclioptions.IOStreams) *LoginOptions {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &LoginOptions{IOStreams: streams}
}
func (o *LoginOptions) GatherInfo() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := o.gatherAuthInfo(); err != nil {
		return err
	}
	if err := o.gatherProjectInfo(); err != nil {
		return err
	}
	return nil
}
func (o *LoginOptions) getClientConfig() (*restclient.Config, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if o.Config != nil {
		return o.Config, nil
	}
	if len(o.Server) == 0 {
		if kterm.IsTerminal(o.In) {
			for !o.serverProvided() {
				defaultServer := defaultClusterURL
				promptMsg := fmt.Sprintf("Server [%s]: ", defaultServer)
				o.Server = term.PromptForStringWithDefault(o.In, o.Out, defaultServer, promptMsg)
			}
		}
	}
	clientConfig := &restclient.Config{}
	if o.RequestTimeout > 0 {
		clientConfig.Timeout = o.RequestTimeout
	}
	serverNormalized, err := config.NormalizeServerURL(o.Server)
	if err != nil {
		return nil, err
	}
	o.Server = serverNormalized
	clientConfig.Host = o.Server
	if len(o.CAFile) > 0 {
		clientConfig.CAFile = o.CAFile
		clientConfig.CAData = nil
	} else if caFile, caData, ok := findExistingClientCA(clientConfig.Host, *o.StartingKubeConfig); ok {
		clientConfig.CAFile = caFile
		clientConfig.CAData = caData
	}
	if err := dialToServer(*clientConfig); err != nil {
		switch err.(type) {
		case x509.UnknownAuthorityError, x509.HostnameError, x509.CertificateInvalidError:
			if o.InsecureTLS || hasExistingInsecureCluster(*clientConfig, *o.StartingKubeConfig) || promptForInsecureTLS(o.In, o.Out, err) {
				clientConfig.Insecure = true
				clientConfig.CAFile = ""
				clientConfig.CAData = nil
			} else {
				return nil, getPrettyErrorForServer(err, o.Server)
			}
		case tls.RecordHeaderError:
			return nil, getPrettyErrorForServer(err, o.Server)
		default:
			if _, ok := err.(*net.OpError); ok {
				return nil, fmt.Errorf("%v - verify you have provided the correct host and port and that the server is currently running.", err)
			}
			return nil, err
		}
	}
	o.Config = clientConfig
	return o.Config, nil
}
func (o *LoginOptions) gatherAuthInfo() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	directClientConfig, err := o.getClientConfig()
	if err != nil {
		return err
	}
	t := *directClientConfig
	clientConfig := &t
	if o.tokenProvided() {
		clientConfig.BearerToken = o.Token
		if me, err := project.WhoAmI(clientConfig); err == nil {
			o.Username = me.Name
			o.Config = clientConfig
			fmt.Fprintf(o.Out, "Logged into %q as %q using the token provided.\n\n", o.Config.Host, o.Username)
			return nil
		} else {
			if kerrors.IsUnauthorized(err) {
				return fmt.Errorf("The token provided is invalid or expired.\n\n")
			}
			return err
		}
	}
	if o.usernameProvided() && !o.passwordProvided() {
		kubeconfig := *o.StartingKubeConfig
		matchingClusters := getMatchingClusters(*clientConfig, kubeconfig)
		for key, context := range o.StartingKubeConfig.Contexts {
			if matchingClusters.Has(context.Cluster) {
				clientcmdConfig := kclientcmd.NewDefaultClientConfig(kubeconfig, &kclientcmd.ConfigOverrides{CurrentContext: key})
				if kubeconfigClientConfig, err := clientcmdConfig.ClientConfig(); err == nil {
					if me, err := project.WhoAmI(kubeconfigClientConfig); err == nil && (o.Username == me.Name) {
						clientConfig.BearerToken = kubeconfigClientConfig.BearerToken
						clientConfig.CertFile = kubeconfigClientConfig.CertFile
						clientConfig.CertData = kubeconfigClientConfig.CertData
						clientConfig.KeyFile = kubeconfigClientConfig.KeyFile
						clientConfig.KeyData = kubeconfigClientConfig.KeyData
						o.Config = clientConfig
						fmt.Fprintf(o.Out, "Logged into %q as %q using existing credentials.\n\n", o.Config.Host, o.Username)
						return nil
					}
				}
			}
		}
	}
	clientConfig.BearerToken = ""
	clientConfig.CertData = []byte{}
	clientConfig.KeyData = []byte{}
	clientConfig.CertFile = o.CertFile
	clientConfig.KeyFile = o.KeyFile
	token, err := tokencmd.RequestToken(o.Config, o.In, o.Username, o.Password)
	if err != nil {
		return err
	}
	clientConfig.BearerToken = token
	me, err := project.WhoAmI(clientConfig)
	if err != nil {
		return err
	}
	o.Username = me.Name
	o.Config = clientConfig
	fmt.Fprint(o.Out, "Login successful.\n\n")
	return nil
}
func (o *LoginOptions) gatherProjectInfo() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	me, err := o.whoAmI()
	if err != nil {
		return err
	}
	if o.Username != me.Name {
		return fmt.Errorf("current user, %v, does not match expected user %v", me.Name, o.Username)
	}
	projectClient, err := projectv1typedclient.NewForConfig(o.Config)
	if err != nil {
		return err
	}
	projectsList, err := projectClient.Projects().List(metav1.ListOptions{})
	if kerrors.IsNotFound(err) || kerrors.IsForbidden(err) {
		fmt.Fprintf(o.Out, "Using \"default\".  You can switch projects with:\n\n '%s project <projectname>'\n", o.CommandName)
		o.Project = "default"
		return nil
	}
	if err != nil {
		return err
	}
	projectsItems := projectsList.Items
	projects := sets.String{}
	for _, project := range projectsItems {
		projects.Insert(project.Name)
	}
	if len(o.DefaultNamespace) > 0 && !projects.Has(o.DefaultNamespace) {
		if currentProject, err := projectClient.Projects().Get(o.DefaultNamespace, metav1.GetOptions{}); err == nil {
			projectsItems = append(projectsItems, *currentProject)
			projects.Insert(currentProject.Name)
		}
	}
	switch len(projectsItems) {
	case 0:
		canRequest, err := loginutil.CanRequestProjects(o.Config, o.DefaultNamespace)
		if err != nil {
			return err
		}
		msg := errors.NoProjectsExistMessage(canRequest, o.CommandName)
		fmt.Fprintf(o.Out, msg)
		o.Project = ""
	case 1:
		o.Project = projectsItems[0].Name
		fmt.Fprintf(o.Out, "You have one project on this server: %q\n\n", o.Project)
		fmt.Fprintf(o.Out, "Using project %q.\n", o.Project)
	default:
		namespace := o.DefaultNamespace
		if !projects.Has(namespace) {
			if namespace != metav1.NamespaceDefault && projects.Has(metav1.NamespaceDefault) {
				namespace = metav1.NamespaceDefault
			} else {
				namespace = projects.List()[0]
			}
		}
		current, err := projectClient.Projects().Get(namespace, metav1.GetOptions{})
		if err != nil && !kerrors.IsNotFound(err) && !kerrors.IsForbidden(err) {
			return err
		}
		o.Project = current.Name
		if len(projectsItems) > projectsItemsSuppressThreshold {
			fmt.Fprintf(o.Out, "You have access to %d projects, the list has been suppressed. You can list all projects with '%s projects'\n\n", len(projectsItems), o.CommandName)
		} else {
			fmt.Fprintf(o.Out, "You have access to the following projects and can switch between them with '%s project <projectname>':\n\n", o.CommandName)
			for _, p := range projects.List() {
				if o.Project == p {
					fmt.Fprintf(o.Out, "  * %s\n", p)
				} else {
					fmt.Fprintf(o.Out, "    %s\n", p)
				}
			}
			fmt.Fprintln(o.Out)
		}
		fmt.Fprintf(o.Out, "Using project %q.\n", o.Project)
	}
	return nil
}
func (o *LoginOptions) SaveConfig() (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(o.Username) == 0 {
		return false, fmt.Errorf("Insufficient data to merge configuration.")
	}
	globalExistedBefore := true
	if _, err := os.Stat(o.PathOptions.GlobalFile); os.IsNotExist(err) {
		globalExistedBefore = false
	}
	newConfig, err := cliconfig.CreateConfig(o.Project, o.Config)
	if err != nil {
		return false, err
	}
	cwd, err := os.Getwd()
	if err != nil {
		return false, err
	}
	baseDir, err := cmdutil.MakeAbs(filepath.Dir(o.PathOptions.GetDefaultFilename()), cwd)
	if err != nil {
		return false, err
	}
	if err := cliconfig.RelativizeClientConfigPaths(newConfig, baseDir); err != nil {
		return false, err
	}
	configToWrite, err := cliconfig.MergeConfig(*o.StartingKubeConfig, *newConfig)
	if err != nil {
		return false, err
	}
	if err := kclientcmd.ModifyConfig(o.PathOptions, *configToWrite, true); err != nil {
		if !os.IsPermission(err) {
			return false, err
		}
		out := &bytes.Buffer{}
		fmt.Fprintf(out, errors.ErrKubeConfigNotWriteable(o.PathOptions.GetDefaultFilename(), o.PathOptions.IsExplicitFile(), err).Error())
		return false, fmt.Errorf("%v", out)
	}
	created := false
	if _, err := os.Stat(o.PathOptions.GlobalFile); err == nil {
		created = created || !globalExistedBefore
	}
	return created, nil
}
func (o LoginOptions) whoAmI() (*userv1.User, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return project.WhoAmI(o.Config)
}
func (o *LoginOptions) usernameProvided() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(o.Username) > 0
}
func (o *LoginOptions) passwordProvided() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(o.Password) > 0
}
func (o *LoginOptions) serverProvided() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return (len(o.Server) > 0)
}
func (o *LoginOptions) tokenProvided() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(o.Token) > 0
}
