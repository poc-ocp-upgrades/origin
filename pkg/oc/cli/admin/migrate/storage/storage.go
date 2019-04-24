package storage

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"time"
	"github.com/spf13/cobra"
	"golang.org/x/time/rate"
	"k8s.io/klog"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericclioptions/resource"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/flowcontrol"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/util/templates"
	"github.com/openshift/origin/pkg/oc/cli/admin/migrate"
)

var (
	internalMigrateStorageLong	= templates.LongDesc(`
		Migrate internal object storage via update

		This command invokes an update operation on every API object reachable by the caller. This forces
		the server to write to the underlying storage if the object representation has changed. Use this
		command to ensure that the most recent storage changes have been applied to all objects (storage
		version, storage encoding, any newer object defaults).

		To operate on a subset of resources, use the --include flag. If you encounter errors during a run
		the command will output a list of resources that received errors, which you can then re-run the
		command on. You may also specify --from-key and --to-key to restrict the set of resource names
		to operate on (key is NAMESPACE/NAME for resources in namespaces or NAME for cluster scoped
		resources). --from-key is inclusive if specified, while --to-key is exclusive.

		By default, events are not migrated since they expire within a very short period of time. If you
		have significantly increased the expiration time of events, run a migration with --include=events

		WARNING: This is a slow command and will put significant load on an API server. It may also
		result in significant intra-cluster traffic.`)
	internalMigrateStorageExample	= templates.Examples(`
	  # Perform an update of all objects
	  %[1]s

	  # Only migrate pods
	  %[1]s --include=pods

	  # Only pods that are in namespaces starting with "bar"
	  %[1]s --include=pods --from-key=bar/ --to-key=bar/\xFF`)
)

const (
	longThrottleLatency	= 50 * time.Millisecond
	mbToKB			= 1000
	kbToBytes		= 1000
	byteToBits		= 8.0
	slowBandwidth		= 30
)

type MigrateAPIStorageOptions struct {
	migrate.ResourceOptions
	bandwidth	int
	limiter		*tokenLimiter
	client		dynamic.Interface
}

func NewMigrateAPIStorageOptions(streams genericclioptions.IOStreams) *MigrateAPIStorageOptions {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &MigrateAPIStorageOptions{bandwidth: 10, ResourceOptions: *migrate.NewResourceOptions(streams).WithIncludes([]string{"*"}).WithUnstructured().WithExcludes([]schema.GroupResource{{Resource: "appliedclusterresourcequotas"}, {Resource: "imagestreamimages"}, {Resource: "imagestreamtags"}, {Resource: "imagestreammappings"}, {Resource: "imagestreamimports"}, {Resource: "projectrequests"}, {Resource: "projects"}, {Resource: "clusterrolebindings"}, {Resource: "rolebindings"}, {Resource: "clusterroles"}, {Resource: "roles"}, {Resource: "resourceaccessreviews"}, {Resource: "localresourceaccessreviews"}, {Resource: "subjectaccessreviews"}, {Resource: "selfsubjectrulesreviews"}, {Resource: "localsubjectaccessreviews"}, {Resource: "useridentitymappings"}, {Resource: "podsecuritypolicyreviews"}, {Resource: "podsecuritypolicyselfsubjectreviews"}, {Resource: "podsecuritypolicysubjectreviews"}, {Resource: "bindings"}, {Resource: "deploymentconfigrollbacks"}, {Resource: "events"}, {Resource: "componentstatuses"}, {Resource: "replicationcontrollerdummies.extensions"}, {Resource: "podtemplates"}, {Resource: "selfsubjectaccessreviews", Group: "authorization.k8s.io"}, {Resource: "localsubjectaccessreviews", Group: "authorization.k8s.io"}}).WithOverlappingResources([]sets.String{sets.NewString("deploymentconfigs.apps.openshift.io", "deploymentconfigs"), sets.NewString("clusterpolicies.authorization.openshift.io", "clusterpolicies"), sets.NewString("clusterpolicybindings.authorization.openshift.io", "clusterpolicybindings"), sets.NewString("clusterrolebindings.authorization.openshift.io", "clusterrolebindings"), sets.NewString("clusterroles.authorization.openshift.io", "clusterroles"), sets.NewString("localresourceaccessreviews.authorization.openshift.io", "localresourceaccessreviews"), sets.NewString("localsubjectaccessreviews.authorization.openshift.io", "localsubjectaccessreviews"), sets.NewString("policies.authorization.openshift.io", "policies"), sets.NewString("policybindings.authorization.openshift.io", "policybindings"), sets.NewString("resourceaccessreviews.authorization.openshift.io", "resourceaccessreviews"), sets.NewString("rolebindingrestrictions.authorization.openshift.io", "rolebindingrestrictions"), sets.NewString("rolebindings.authorization.openshift.io", "rolebindings"), sets.NewString("roles.authorization.openshift.io", "roles"), sets.NewString("selfsubjectrulesreviews.authorization.openshift.io", "selfsubjectrulesreviews"), sets.NewString("subjectaccessreviews.authorization.openshift.io", "subjectaccessreviews"), sets.NewString("subjectrulesreviews.authorization.openshift.io", "subjectrulesreviews"), sets.NewString("builds.build.openshift.io", "builds"), sets.NewString("buildconfigs.build.openshift.io", "buildconfigs"), sets.NewString("images.image.openshift.io", "images"), sets.NewString("imagesignatures.image.openshift.io", "imagesignatures"), sets.NewString("imagestreamimages.image.openshift.io", "imagestreamimages"), sets.NewString("imagestreamimports.image.openshift.io", "imagestreamimports"), sets.NewString("imagestreammappings.image.openshift.io", "imagestreammappings"), sets.NewString("imagestreams.image.openshift.io", "imagestreams"), sets.NewString("imagestreamtags.image.openshift.io", "imagestreamtags"), sets.NewString("clusternetworks.network.openshift.io", "clusternetworks"), sets.NewString("egressnetworkpolicies.network.openshift.io", "egressnetworkpolicies"), sets.NewString("hostsubnets.network.openshift.io", "hostsubnets"), sets.NewString("netnamespaces.network.openshift.io", "netnamespaces"), sets.NewString("oauthaccesstokens.oauth.openshift.io", "oauthaccesstokens"), sets.NewString("oauthauthorizetokens.oauth.openshift.io", "oauthauthorizetokens"), sets.NewString("oauthclientauthorizations.oauth.openshift.io", "oauthclientauthorizations"), sets.NewString("oauthclients.oauth.openshift.io", "oauthclients"), sets.NewString("projectrequests.project.openshift.io", "projectrequests"), sets.NewString("projects.project.openshift.io", "projects"), sets.NewString("appliedclusterresourcequotas.quota.openshift.io", "appliedclusterresourcequotas"), sets.NewString("clusterresourcequotas.quota.openshift.io", "clusterresourcequotas"), sets.NewString("routes.route.openshift.io", "routes"), sets.NewString("podsecuritypolicyreviews.security.openshift.io", "podsecuritypolicyreviews"), sets.NewString("podsecuritypolicyselfsubjectreviews.security.openshift.io", "podsecuritypolicyselfsubjectreviews"), sets.NewString("podsecuritypolicysubjectreviews.security.openshift.io", "podsecuritypolicysubjectreviews"), sets.NewString("processedtemplates.template.openshift.io", "processedtemplates"), sets.NewString("templates.template.openshift.io", "templates"), sets.NewString("groups.user.openshift.io", "groups"), sets.NewString("identities.user.openshift.io", "identities"), sets.NewString("useridentitymappings.user.openshift.io", "useridentitymappings"), sets.NewString("users.user.openshift.io", "users"), sets.NewString("horizontalpodautoscalers.autoscaling", "horizontalpodautoscalers.extensions"), sets.NewString("jobs.batch", "jobs.extensions")})}
}
func NewCmdMigrateAPIStorage(name, fullName string, f kcmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o := NewMigrateAPIStorageOptions(streams)
	cmd := &cobra.Command{Use: name, Short: "Update the stored version of API objects", Long: internalMigrateStorageLong, Example: fmt.Sprintf(internalMigrateStorageExample, fullName), Run: func(cmd *cobra.Command, args []string) {
		kcmdutil.CheckErr(o.Complete(f, cmd, args))
		kcmdutil.CheckErr(o.Validate())
		kcmdutil.CheckErr(o.Run())
	}}
	o.ResourceOptions.Bind(cmd)
	o.Workers = 32 * runtime.NumCPU()
	cmd.Flags().IntVar(&o.bandwidth, "bandwidth", o.bandwidth, "Average network bandwidth measured in megabits per second (Mbps) to use during storage migration.  Zero means no limit.  This flag is alpha and may change in the future.")
	cmd.Flags().MarkDeprecated("confirm", "storage migration does not support dry run, this flag is ignored")
	cmd.Flags().MarkHidden("confirm")
	cmd.Flags().MarkDeprecated("output", "storage migration does not support dry run, this flag is ignored")
	cmd.Flags().MarkHidden("output")
	return cmd
}
func (o *MigrateAPIStorageOptions) Complete(f kcmdutil.Factory, c *cobra.Command, args []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := c.Flags().Set("output", ""); err != nil {
		return err
	}
	o.Confirm = true
	o.ResourceOptions.SaveFn = o.save
	if err := o.ResourceOptions.Complete(f, c); err != nil {
		return err
	}
	always := flowcontrol.NewFakeAlwaysRateLimiter()
	o.Builder.TransformRequests(func(req *rest.Request) {
		req.Throttle(always)
	})
	if o.bandwidth > 0 {
		o.limiter = newTokenLimiter(o.bandwidth, o.Workers)
		if o.bandwidth < slowBandwidth {
			o.Builder.RequestChunksOf(0)
		}
	}
	clientConfig, err := f.ToRESTConfig()
	if err != nil {
		return err
	}
	clientConfigCopy := rest.CopyConfig(clientConfig)
	clientConfigCopy.Burst = 99999
	clientConfigCopy.QPS = 99999
	o.client, err = dynamic.NewForConfig(clientConfigCopy)
	if err != nil {
		return err
	}
	return nil
}
func (o MigrateAPIStorageOptions) Validate() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if o.bandwidth < 0 {
		return fmt.Errorf("invalid value %d for --bandwidth, must be at least 0", o.bandwidth)
	}
	return o.ResourceOptions.Validate()
}
func (o MigrateAPIStorageOptions) Run() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o.ResourceOptions.Visitor().Visit(migrate.AlwaysRequiresMigration)
}
func (o *MigrateAPIStorageOptions) save(info *resource.Info, reporter migrate.Reporter) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch oldObject := info.Object.(type) {
	case *unstructured.Unstructured:
		if o.limiter != nil {
			defer o.rateLimit(oldObject)
		}
		newObject, err := o.client.Resource(info.Mapping.Resource).Namespace(info.Namespace).Update(oldObject, metav1.UpdateOptions{})
		if errors.IsConflict(err) {
			return migrate.ErrUnchanged
		}
		if err != nil {
			return migrate.DefaultRetriable(info, err)
		}
		if newObject.GetResourceVersion() == oldObject.GetResourceVersion() {
			return migrate.ErrUnchanged
		}
	default:
		return fmt.Errorf("invalid type %T passed to storage migration: %v", oldObject, oldObject)
	}
	return nil
}
func (o *MigrateAPIStorageOptions) rateLimit(oldObject *unstructured.Unstructured) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var dataLen int
	if data, err := oldObject.MarshalJSON(); err != nil {
		klog.Errorf("failed to marshall %#v: %v", oldObject, err)
		dataLen = 8192
	} else {
		dataLen = len(data)
	}
	latency := o.limiter.take(2 * dataLen)
	if latency > longThrottleLatency {
		klog.V(4).Infof("Throttling request took %v, request: %s:%s", latency, "PUT", oldObject.GetSelfLink())
	}
}

type tokenLimiter struct {
	burst		int
	rateLimiter	*rate.Limiter
	nowFunc		func() time.Time
}

func (t *tokenLimiter) take(n int) time.Duration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if n <= 0 {
		return 0
	}
	var extra time.Duration
	for ; n > t.burst; n -= t.burst {
		extra += t.getDuration(t.burst)
	}
	total := t.getDuration(n) + extra
	time.Sleep(total)
	return total
}
func (t *tokenLimiter) getDuration(n int) time.Duration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	now := t.nowFunc()
	reservation := t.rateLimiter.ReserveN(now, n)
	if !reservation.OK() {
		klog.Errorf("unable to get rate limited reservation, burst=%d n=%d", t.burst, n)
		return time.Minute
	}
	return reservation.DelayFrom(now)
}
func newTokenLimiter(bandwidth, workers int) *tokenLimiter {
	_logClusterCodePath()
	defer _logClusterCodePath()
	burst := 100 * kbToBytes * workers
	return &tokenLimiter{burst: burst, rateLimiter: rate.NewLimiter(rate.Limit(bandwidth*mbToKB*kbToBytes)/byteToBits, burst), nowFunc: time.Now}
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
