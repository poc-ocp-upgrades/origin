package describe

import (
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
	"k8s.io/kubernetes/pkg/kubectl/describe/versioned"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/apis/autoscaling"
	"k8s.io/kubernetes/pkg/kubectl/describe"
	"github.com/openshift/api/apps"
	appsv1 "github.com/openshift/api/apps/v1"
	appstypedclient "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	"github.com/openshift/origin/pkg/api/legacy"
	appsutil "github.com/openshift/origin/pkg/apps/util"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	appsedges "github.com/openshift/origin/pkg/oc/lib/graph/appsgraph"
	appsgraph "github.com/openshift/origin/pkg/oc/lib/graph/appsgraph/nodes"
	"github.com/openshift/origin/pkg/oc/lib/graph/genericgraph"
	kubegraph "github.com/openshift/origin/pkg/oc/lib/graph/kubegraph/nodes"
)

const (
	maxDisplayDeployments		= 3
	maxDisplayDeploymentsEvents	= 8
)

type DeploymentConfigDescriber struct {
	appsClient	appstypedclient.AppsV1Interface
	kubeClient	kubernetes.Interface
	config		*appsv1.DeploymentConfig
}

func NewDeploymentConfigDescriber(client appstypedclient.AppsV1Interface, kclient kubernetes.Interface, config *appsv1.DeploymentConfig) *DeploymentConfigDescriber {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &DeploymentConfigDescriber{appsClient: client, kubeClient: kclient, config: config}
}
func (d *DeploymentConfigDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var deploymentConfig *appsv1.DeploymentConfig
	if d.config != nil {
		deploymentConfig = d.config
	} else {
		var err error
		deploymentConfig, err = d.appsClient.DeploymentConfigs(namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return "", err
		}
	}
	return tabbedString(func(out *tabwriter.Writer) error {
		formatMeta(out, deploymentConfig.ObjectMeta)
		var (
			deploymentsHistory	[]*corev1.ReplicationController
			activeDeploymentName	string
		)
		if d.config == nil {
			if rcs, err := d.kubeClient.CoreV1().ReplicationControllers(namespace).List(metav1.ListOptions{LabelSelector: appsutil.ConfigSelector(deploymentConfig.Name).String()}); err == nil {
				deploymentsHistory = make([]*corev1.ReplicationController, 0, len(rcs.Items))
				for i := range rcs.Items {
					deploymentsHistory = append(deploymentsHistory, &rcs.Items[i])
				}
			}
		}
		if deploymentConfig.Status.LatestVersion == 0 {
			formatString(out, "Latest Version", "Not deployed")
		} else {
			formatString(out, "Latest Version", strconv.FormatInt(deploymentConfig.Status.LatestVersion, 10))
		}
		printDeploymentConfigSpec(d.kubeClient, *deploymentConfig, out)
		fmt.Fprintln(out)
		latestDeploymentName := appsutil.LatestDeploymentNameForConfig(deploymentConfig)
		if activeDeployment := appsutil.ActiveDeployment(deploymentsHistory); activeDeployment != nil {
			activeDeploymentName = activeDeployment.Name
		}
		var deployment *corev1.ReplicationController
		isNotDeployed := len(deploymentsHistory) == 0
		for _, item := range deploymentsHistory {
			if item.Name == latestDeploymentName {
				deployment = item
			}
		}
		if deployment == nil {
			isNotDeployed = true
		}
		if isNotDeployed {
			formatString(out, "Latest Deployment", "<none>")
		} else {
			header := fmt.Sprintf("Deployment #%d (latest)", appsutil.DeploymentVersionFor(deployment))
			printDeploymentRc(deployment, d.kubeClient, out, header, (deployment.Name == activeDeploymentName) || len(deploymentsHistory) == 1)
		}
		if d.config == nil && !isNotDeployed {
			var sorted []*corev1.ReplicationController
			for i := range deploymentsHistory {
				sorted = append(sorted, deploymentsHistory[i])
			}
			sort.Sort(sort.Reverse(OverlappingControllers(sorted)))
			counter := 1
			for _, item := range sorted {
				if item.Name != latestDeploymentName && deploymentConfig.Name == appsutil.DeploymentConfigNameFor(item) {
					header := fmt.Sprintf("Deployment #%d", appsutil.DeploymentVersionFor(item))
					printDeploymentRc(item, d.kubeClient, out, header, item.Name == activeDeploymentName)
					counter++
				}
				if counter == maxDisplayDeployments {
					break
				}
			}
		}
		if settings.ShowEvents {
			if events, err := d.kubeClient.CoreV1().Events(deploymentConfig.Namespace).Search(legacyscheme.Scheme, deploymentConfig); err == nil && events != nil {
				latestDeploymentEvents := &corev1.EventList{Items: []corev1.Event{}}
				for i := len(events.Items); i != 0 && i > len(events.Items)-maxDisplayDeploymentsEvents; i-- {
					latestDeploymentEvents.Items = append(latestDeploymentEvents.Items, events.Items[i-1])
				}
				fmt.Fprintln(out)
				pw := versioned.NewPrefixWriter(out)
				versioned.DescribeEvents(latestDeploymentEvents, pw)
			}
		}
		return nil
	})
}

type OverlappingControllers []*corev1.ReplicationController

func (o OverlappingControllers) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(o)
}
func (o OverlappingControllers) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o[i], o[j] = o[j], o[i]
}
func (o OverlappingControllers) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if o[i].CreationTimestamp.Equal(&o[j].CreationTimestamp) {
		return o[i].Name < o[j].Name
	}
	return o[i].CreationTimestamp.Before(&o[j].CreationTimestamp)
}
func multilineStringArray(sep, indent string, args ...string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i, s := range args {
		if strings.HasSuffix(s, "\n") {
			s = strings.TrimSuffix(s, "\n")
		}
		if strings.Contains(s, "\n") {
			s = "\n" + indent + strings.Join(strings.Split(s, "\n"), "\n"+indent)
		}
		args[i] = s
	}
	strings.TrimRight(args[len(args)-1], "\n ")
	return strings.Join(args, " ")
}
func printStrategy(strategy appsv1.DeploymentStrategy, indent string, w *tabwriter.Writer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if strategy.CustomParams != nil {
		if len(strategy.CustomParams.Image) == 0 {
			fmt.Fprintf(w, "%sImage:\t%s\n", indent, "<default>")
		} else {
			fmt.Fprintf(w, "%sImage:\t%s\n", indent, strategy.CustomParams.Image)
		}
		if len(strategy.CustomParams.Environment) > 0 {
			fmt.Fprintf(w, "%sEnvironment:\t%s\n", indent, formatLabels(convertEnv(strategy.CustomParams.Environment)))
		}
		if len(strategy.CustomParams.Command) > 0 {
			fmt.Fprintf(w, "%sCommand:\t%v\n", indent, multilineStringArray(" ", "\t  ", strategy.CustomParams.Command...))
		}
	}
	if strategy.RecreateParams != nil {
		pre := strategy.RecreateParams.Pre
		mid := strategy.RecreateParams.Mid
		post := strategy.RecreateParams.Post
		if pre != nil {
			printHook("Pre-deployment", pre, indent, w)
		}
		if mid != nil {
			printHook("Mid-deployment", mid, indent, w)
		}
		if post != nil {
			printHook("Post-deployment", post, indent, w)
		}
	}
	if strategy.RollingParams != nil {
		pre := strategy.RollingParams.Pre
		post := strategy.RollingParams.Post
		if pre != nil {
			printHook("Pre-deployment", pre, indent, w)
		}
		if post != nil {
			printHook("Post-deployment", post, indent, w)
		}
	}
}
func printHook(prefix string, hook *appsv1.LifecycleHook, indent string, w io.Writer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if hook.ExecNewPod != nil {
		fmt.Fprintf(w, "%s%s hook (pod type, failure policy: %s):\n", indent, prefix, hook.FailurePolicy)
		fmt.Fprintf(w, "%s  Container:\t%s\n", indent, hook.ExecNewPod.ContainerName)
		fmt.Fprintf(w, "%s  Command:\t%v\n", indent, multilineStringArray(" ", "\t  ", hook.ExecNewPod.Command...))
		if len(hook.ExecNewPod.Env) > 0 {
			fmt.Fprintf(w, "%s  Env:\t%s\n", indent, formatLabels(convertEnv(hook.ExecNewPod.Env)))
		}
	}
	if len(hook.TagImages) > 0 {
		fmt.Fprintf(w, "%s%s hook (tag images, failure policy: %s):\n", indent, prefix, hook.FailurePolicy)
		for _, image := range hook.TagImages {
			fmt.Fprintf(w, "%s  Tag:\tcontainer %s to %s %s %s\n", indent, image.ContainerName, image.To.Kind, image.To.Name, image.To.Namespace)
		}
	}
}
func printTriggers(triggers []appsv1.DeploymentTriggerPolicy, w *tabwriter.Writer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(triggers) == 0 {
		formatString(w, "Triggers", "<none>")
		return
	}
	printLabels := []string{}
	for _, t := range triggers {
		switch t.Type {
		case appsv1.DeploymentTriggerOnConfigChange:
			printLabels = append(printLabels, "Config")
		case appsv1.DeploymentTriggerOnImageChange:
			if len(t.ImageChangeParams.From.Name) > 0 {
				name, tag, _ := imageapi.SplitImageStreamTag(t.ImageChangeParams.From.Name)
				printLabels = append(printLabels, fmt.Sprintf("Image(%s@%s, auto=%v)", name, tag, t.ImageChangeParams.Automatic))
			}
		}
	}
	desc := strings.Join(printLabels, ", ")
	formatString(w, "Triggers", desc)
}
func printDeploymentConfigSpec(kc kubernetes.Interface, dc appsv1.DeploymentConfig, w *tabwriter.Writer) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	spec := dc.Spec
	formatString(w, "Selector", formatLabels(spec.Selector))
	test := ""
	if spec.Test {
		test = " (test, will be scaled down between deployments)"
	}
	formatString(w, "Replicas", fmt.Sprintf("%d%s", spec.Replicas, test))
	if spec.Paused {
		formatString(w, "Paused", "yes")
	}
	printAutoscalingInfo([]schema.GroupResource{apps.Resource("DeploymentConfig"), legacy.Resource("DeploymentConfig")}, dc.Namespace, dc.Name, kc, w)
	printTriggers(spec.Triggers, w)
	formatString(w, "Strategy", spec.Strategy.Type)
	printStrategy(spec.Strategy, "  ", w)
	if dc.Spec.MinReadySeconds > 0 {
		formatString(w, "MinReadySeconds", fmt.Sprintf("%d", spec.MinReadySeconds))
	}
	fmt.Fprintf(w, "Template:\n")
	versioned.DescribePodTemplate(spec.Template, versioned.NewPrefixWriter(w))
	return nil
}
func printAutoscalingInfo(res []schema.GroupResource, namespace, name string, kclient kubernetes.Interface, w *tabwriter.Writer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	hpaList, err := kclient.AutoscalingV1().HorizontalPodAutoscalers(namespace).List(metav1.ListOptions{LabelSelector: labels.Everything().String()})
	if err != nil {
		return
	}
	scaledBy := []autoscalingv1.HorizontalPodAutoscaler{}
	for _, hpa := range hpaList.Items {
		for _, r := range res {
			if hpa.Spec.ScaleTargetRef.Name == name && hpa.Spec.ScaleTargetRef.Kind == r.String() {
				scaledBy = append(scaledBy, hpa)
			}
		}
	}
	for _, hpa := range scaledBy {
		fmt.Fprintf(w, "Autoscaling:\tbetween %d and %d replicas", *hpa.Spec.MinReplicas, hpa.Spec.MaxReplicas)
		legacyHpa := &autoscaling.HorizontalPodAutoscaler{}
		if err := legacyscheme.Scheme.Convert(&hpa, legacyHpa, nil); err != nil {
			panic(err)
		}
		targetDescriptions := formatHPATargets(legacyHpa)
		if len(targetDescriptions) == 1 {
			fmt.Fprintf(w, " targeting %s\n", targetDescriptions[0])
		} else {
			fmt.Fprintf(w, "\n")
			for _, description := range targetDescriptions {
				fmt.Fprintf(w, "\t  targeting %s\n", description)
			}
		}
		break
	}
}
func formatHPATargets(hpa *autoscaling.HorizontalPodAutoscaler) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	descriptions := make([]string, len(hpa.Spec.Metrics))
	for i, metricSpec := range hpa.Spec.Metrics {
		switch metricSpec.Type {
		case autoscaling.PodsMetricSourceType:
			descriptions[i] = fmt.Sprintf("%s %s average per pod", metricSpec.Pods.Target.AverageValue.String(), metricSpec.Pods.Metric.Name)
		case autoscaling.ObjectMetricSourceType:
			targetObjDesc := fmt.Sprintf("%s %s", metricSpec.Object.Target.Type, metricSpec.Object.Metric.Name)
			descriptions[i] = fmt.Sprintf("%s %s on %s", metricSpec.Object.Target.Value.String(), metricSpec.Object.Metric.Name, targetObjDesc)
		case autoscaling.ResourceMetricSourceType:
			if metricSpec.Resource.Target.AverageValue != nil {
				descriptions[i] = fmt.Sprintf("%s %s average per pod", metricSpec.Resource.Target.AverageValue.String(), metricSpec.Resource.Name)
			} else if metricSpec.Resource.Target.AverageUtilization != nil {
				descriptions[i] = fmt.Sprintf("%d%% %s average per pod", *metricSpec.Resource.Target.AverageUtilization, metricSpec.Resource.Name)
			} else {
				descriptions[i] = "<unset resource metric>"
			}
		default:
			descriptions[i] = "<unknown metric type>"
		}
	}
	return descriptions
}
func printDeploymentRc(deployment *corev1.ReplicationController, kubeClient kubernetes.Interface, w io.Writer, header string, verbose bool) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(header) > 0 {
		fmt.Fprintf(w, "%v:\n", header)
	}
	if verbose {
		fmt.Fprintf(w, "\tName:\t%s\n", deployment.Name)
	}
	timeAt := strings.ToLower(formatRelativeTime(deployment.CreationTimestamp.Time))
	fmt.Fprintf(w, "\tCreated:\t%s ago\n", timeAt)
	fmt.Fprintf(w, "\tStatus:\t%s\n", appsutil.DeploymentStatusFor(deployment))
	if deployment.Spec.Replicas != nil {
		fmt.Fprintf(w, "\tReplicas:\t%d current / %d desired\n", deployment.Status.Replicas, *deployment.Spec.Replicas)
	}
	if verbose {
		fmt.Fprintf(w, "\tSelector:\t%s\n", formatLabels(deployment.Spec.Selector))
		fmt.Fprintf(w, "\tLabels:\t%s\n", formatLabels(deployment.Labels))
		running, waiting, succeeded, failed, err := getPodStatusForDeployment(deployment, kubeClient)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "\tPods Status:\t%d Running / %d Waiting / %d Succeeded / %d Failed\n", running, waiting, succeeded, failed)
	}
	return nil
}
func getPodStatusForDeployment(deployment *corev1.ReplicationController, kubeClient kubernetes.Interface) (running, waiting, succeeded, failed int, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rcPods, err := kubeClient.CoreV1().Pods(deployment.Namespace).List(metav1.ListOptions{LabelSelector: labels.Set(deployment.Spec.Selector).AsSelector().String()})
	if err != nil {
		return
	}
	for _, pod := range rcPods.Items {
		switch pod.Status.Phase {
		case corev1.PodRunning:
			running++
		case corev1.PodPending:
			waiting++
		case corev1.PodSucceeded:
			succeeded++
		case corev1.PodFailed:
			failed++
		}
	}
	return
}

type LatestDeploymentsDescriber struct {
	count		int
	appsClient	appstypedclient.AppsV1Interface
	kubeClient	kubernetes.Interface
}

func NewLatestDeploymentsDescriber(client appstypedclient.AppsV1Interface, kclient kubernetes.Interface, count int) *LatestDeploymentsDescriber {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &LatestDeploymentsDescriber{count: count, appsClient: client, kubeClient: kclient}
}
func (d *LatestDeploymentsDescriber) Describe(namespace, name string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var f formatter
	config, err := d.appsClient.DeploymentConfigs(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	var deployments []corev1.ReplicationController
	if d.count == -1 || d.count > 1 {
		list, err := d.kubeClient.CoreV1().ReplicationControllers(namespace).List(metav1.ListOptions{LabelSelector: appsutil.ConfigSelector(name).String()})
		if err != nil && !kerrors.IsNotFound(err) {
			return "", err
		}
		deployments = list.Items
	} else {
		deploymentName := appsutil.LatestDeploymentNameForConfig(config)
		deployment, err := d.kubeClient.CoreV1().ReplicationControllers(config.Namespace).Get(deploymentName, metav1.GetOptions{})
		if err != nil && !kerrors.IsNotFound(err) {
			return "", err
		}
		if deployment != nil {
			deployments = []corev1.ReplicationController{*deployment}
		}
	}
	g := genericgraph.New()
	dcNode := appsgraph.EnsureDeploymentConfigNode(g, config)
	for i := range deployments {
		kubegraph.EnsureReplicationControllerNode(g, &deployments[i])
	}
	appsedges.AddTriggerDeploymentConfigsEdges(g, dcNode)
	appsedges.AddDeploymentConfigsDeploymentEdges(g, dcNode)
	activeDeployment, inactiveDeployments := appsedges.RelevantDeployments(g, dcNode)
	return tabbedString(func(out *tabwriter.Writer) error {
		descriptions := describeDeploymentConfigDeployments(f, dcNode, activeDeployment, inactiveDeployments, nil, d.count)
		for i, description := range descriptions {
			descriptions[i] = fmt.Sprintf("%v %v", name, description)
		}
		printLines(out, "", 0, descriptions...)
		return nil
	})
}
