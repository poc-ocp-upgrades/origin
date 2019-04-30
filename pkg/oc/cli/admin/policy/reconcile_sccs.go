package policy

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"github.com/spf13/cobra"
	sccutil "github.com/openshift/origin/pkg/security/securitycontextconstraints/util"
	kapierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericclioptions/printers"
	corev1typedclient "k8s.io/client-go/kubernetes/typed/core/v1"
	kapihelper "k8s.io/kubernetes/pkg/apis/core/helper"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/scheme"
	"k8s.io/kubernetes/pkg/kubectl/util/templates"
	securityv1 "github.com/openshift/api/security/v1"
	securityv1typedclient "github.com/openshift/client-go/security/clientset/versioned/typed/security/v1"
	"github.com/openshift/origin/pkg/cmd/server/bootstrappolicy"
)

const ReconcileSCCRecommendedName = "reconcile-sccs"

type ReconcileSCCOptions struct {
	PrintFlags	*genericclioptions.PrintFlags
	Printer		printers.ResourcePrinter
	Confirmed	bool
	Union		bool
	InfraNamespace	string
	SCCClient	securityv1typedclient.SecurityContextConstraintsInterface
	NSClient	corev1typedclient.NamespaceInterface
	genericclioptions.IOStreams
}

var (
	reconcileSCCLong	= templates.LongDesc(`
		Replace cluster SCCs to match the recommended bootstrap policy

		This command will inspect the cluster SCCs against the recommended bootstrap SCCs.
		Any cluster SCC that does not match will be replaced by the recommended SCC.
		This command will not remove any additional cluster SCCs.  By default, this command
		will not remove additional users and groups that have been granted access to the SCC and
		will preserve existing priorities (but will always reconcile unset priorities and the policy
		definition), labels, and annotations.

		You can see which cluster SCCs have recommended changes by choosing an output type.`)
	reconcileSCCExample	= templates.Examples(`
		# Display the cluster SCCs that would be modified
	  %[1]s

	  # Update cluster SCCs that don't match the current defaults preserving additional grants
	  # for users, groups, labels, annotations and keeping any priorities that are already set
	  %[1]s --confirm

	  # Replace existing users, groups, labels, annotations, and priorities that do not match defaults
	  %[1]s --additive-only=false --confirm`)
)

func NewDefaultReconcileSCCOptions(streams genericclioptions.IOStreams) *ReconcileSCCOptions {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ReconcileSCCOptions{PrintFlags: genericclioptions.NewPrintFlags("").WithTypeSetter(scheme.Scheme).WithDefaultOutput("yaml"), Union: true, InfraNamespace: bootstrappolicy.DefaultOpenShiftInfraNamespace, IOStreams: streams}
}
func NewCmdReconcileSCC(name, fullName string, f kcmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o := NewDefaultReconcileSCCOptions(streams)
	cmd := &cobra.Command{Use: name, Short: "Replace cluster SCCs to match the recommended bootstrap policy", Long: reconcileSCCLong, Example: fmt.Sprintf(reconcileSCCExample, fullName), Run: func(cmd *cobra.Command, args []string) {
		kcmdutil.CheckErr(o.Complete(cmd, f, args))
		kcmdutil.CheckErr(o.Validate())
		kcmdutil.CheckErr(o.RunReconcileSCCs(cmd, f))
	}}
	cmd.Flags().BoolVar(&o.Confirmed, "confirm", o.Confirmed, "If true, specify that cluster SCCs should be modified. Defaults to false, displaying what would be replaced but not actually replacing anything.")
	cmd.Flags().BoolVar(&o.Union, "additive-only", o.Union, "If true, preserves extra users, groups, labels and annotations in the SCC as well as existing priorities.")
	cmd.Flags().StringVar(&o.InfraNamespace, "infrastructure-namespace", o.InfraNamespace, "Name of the infrastructure namespace.")
	o.PrintFlags.AddFlags(cmd)
	return cmd
}
func (o *ReconcileSCCOptions) Complete(cmd *cobra.Command, f kcmdutil.Factory, args []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) != 0 {
		return kcmdutil.UsageErrorf(cmd, "no arguments are allowed")
	}
	clientConfig, err := f.ToRESTConfig()
	if err != nil {
		return err
	}
	kClient, err := corev1typedclient.NewForConfig(clientConfig)
	if err != nil {
		return err
	}
	securityClient, err := securityv1typedclient.NewForConfig(clientConfig)
	if err != nil {
		return err
	}
	o.SCCClient = securityClient.SecurityContextConstraints()
	o.NSClient = kClient.Namespaces()
	o.Printer, err = o.PrintFlags.ToPrinter()
	if err != nil {
		return err
	}
	return nil
}
func (o *ReconcileSCCOptions) Validate() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if o.SCCClient == nil {
		return errors.New("a SCC client is required")
	}
	if _, err := o.NSClient.Get(o.InfraNamespace, metav1.GetOptions{}); err != nil {
		return fmt.Errorf("Failed to GET reconcile SCC namespace %s: %v", o.InfraNamespace, err)
	}
	return nil
}
func (o *ReconcileSCCOptions) RunReconcileSCCs(cmd *cobra.Command, f kcmdutil.Factory) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newSCCs, changedSCCs, err := o.ChangedSCCs()
	if err != nil {
		return err
	}
	if (len(changedSCCs) + len(newSCCs)) == 0 {
		return nil
	}
	if !o.Confirmed {
		objs := []runtime.Object{}
		for _, obj := range newSCCs {
			objs = append(objs, obj)
		}
		for _, obj := range changedSCCs {
			objs = append(objs, obj)
		}
		if err := printObjectList(objs, o.Printer, o.Out); err != nil {
			return err
		}
	}
	if o.Confirmed {
		return o.ReplaceChangedSCCs(newSCCs, changedSCCs)
	}
	return nil
}
func printObjectList(objs []runtime.Object, printer printers.ResourcePrinter, out io.Writer) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	list := &unstructured.UnstructuredList{Object: map[string]interface{}{"kind": "List", "apiVersion": "v1", "metadata": map[string]interface{}{}}}
	for _, obj := range objs {
		unstrObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
		if err != nil {
			return err
		}
		list.Items = append(list.Items, unstructured.Unstructured{Object: unstrObj})
	}
	return printer.PrintObj(list, out)
}
func (o *ReconcileSCCOptions) ChangedSCCs() ([]*securityv1.SecurityContextConstraints, []*securityv1.SecurityContextConstraints, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	toUpdateSCCs := []*securityv1.SecurityContextConstraints{}
	toCreateSCCs := []*securityv1.SecurityContextConstraints{}
	groups, users := bootstrappolicy.GetBoostrapSCCAccess(o.InfraNamespace)
	bootstrapSCCs := bootstrappolicy.GetBootstrapSecurityContextConstraints(groups, users)
	for _, expectedSCC := range bootstrapSCCs {
		actualSCC, err := o.SCCClient.Get(expectedSCC.Name, metav1.GetOptions{})
		if kapierrors.IsNotFound(err) {
			toCreateSCCs = append(toCreateSCCs, expectedSCC)
			continue
		}
		if err != nil {
			return nil, nil, err
		}
		if updatedSCC, needsUpdating := o.computeUpdatedSCC(*expectedSCC, *actualSCC); needsUpdating {
			toUpdateSCCs = append(toUpdateSCCs, updatedSCC)
		}
	}
	return toCreateSCCs, toUpdateSCCs, nil
}
func (o *ReconcileSCCOptions) ReplaceChangedSCCs(newSCCs, changedSCCs []*securityv1.SecurityContextConstraints) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	applyOnConstraints := func(sccs []*securityv1.SecurityContextConstraints, fn func(*securityv1.SecurityContextConstraints) (*securityv1.SecurityContextConstraints, error)) error {
		for i := range sccs {
			updatedSCC, err := fn(sccs[i])
			if err != nil {
				return err
			}
			fmt.Fprintf(o.Out, "securitycontextconstraints/%s\n", updatedSCC.Name)
		}
		return nil
	}
	err := applyOnConstraints(newSCCs, o.SCCClient.Create)
	if err != nil {
		return err
	}
	return applyOnConstraints(changedSCCs, o.SCCClient.Update)
}
func (o *ReconcileSCCOptions) computeUpdatedSCC(expected securityv1.SecurityContextConstraints, actual securityv1.SecurityContextConstraints) (*securityv1.SecurityContextConstraints, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	needsUpdate := false
	if o.Union {
		groupSet := sets.NewString(actual.Groups...)
		groupSet.Insert(expected.Groups...)
		expected.Groups = groupSet.List()
		userSet := sets.NewString(actual.Users...)
		userSet.Insert(expected.Users...)
		expected.Users = userSet.List()
		if actual.Priority != nil {
			expected.Priority = actual.Priority
		}
		expected.Labels = MergeMaps(expected.Labels, actual.Labels)
		expected.Annotations = MergeMaps(expected.Annotations, actual.Annotations)
	}
	sortVolumes(&expected)
	sortVolumes(&actual)
	sort.StringSlice(actual.Groups).Sort()
	sort.StringSlice(actual.Users).Sort()
	sort.StringSlice(expected.Groups).Sort()
	sort.StringSlice(expected.Users).Sort()
	updated := expected
	updated.ObjectMeta = actual.ObjectMeta
	updated.ObjectMeta.Labels = expected.Labels
	updated.ObjectMeta.Annotations = expected.Annotations
	if !kapihelper.Semantic.DeepEqual(updated, actual) {
		needsUpdate = true
	}
	return &updated, needsUpdate
}
func sortVolumes(scc *securityv1.SecurityContextConstraints) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if scc.Volumes == nil || len(scc.Volumes) == 0 {
		return
	}
	volumes := sccutil.FSTypeToStringSet(scc.Volumes).List()
	sort.StringSlice(volumes).Sort()
	scc.Volumes = sliceToFSType(volumes)
}
func sliceToFSType(s []string) []securityv1.FSType {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fsTypes := []securityv1.FSType{}
	for _, v := range s {
		fsTypes = append(fsTypes, securityv1.FSType(v))
	}
	return fsTypes
}
func MergeMaps(a, b map[string]string) map[string]string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if a == nil && b == nil {
		return nil
	}
	res := make(map[string]string)
	for k, v := range a {
		res[k] = v
	}
	for k, v := range b {
		res[k] = v
	}
	return res
}
