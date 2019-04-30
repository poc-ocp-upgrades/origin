package deploymentconfigs

import (
	"bytes"
	"fmt"
	"sort"
	"text/tabwriter"
	"k8s.io/kubernetes/pkg/kubectl/describe/versioned"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/kubernetes/pkg/kubectl"
	appsv1 "github.com/openshift/api/apps/v1"
	appsutil "github.com/openshift/origin/pkg/apps/util"
)

func NewDeploymentConfigHistoryViewer(kc kubernetes.Interface) kubectl.HistoryViewer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &DeploymentConfigHistoryViewer{rn: kc.CoreV1()}
}

type DeploymentConfigHistoryViewer struct {
	rn corev1client.ReplicationControllersGetter
}

var _ kubectl.HistoryViewer = &DeploymentConfigHistoryViewer{}

func (h *DeploymentConfigHistoryViewer) ViewHistory(namespace, name string, revision int64) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	opts := metav1.ListOptions{LabelSelector: appsutil.ConfigSelector(name).String()}
	deploymentList, err := h.rn.ReplicationControllers(namespace).List(opts)
	if err != nil {
		return "", err
	}
	if len(deploymentList.Items) == 0 {
		return "No rollout history found.", nil
	}
	items := deploymentList.Items
	history := make([]*v1.ReplicationController, 0, len(items))
	for i := range items {
		history = append(history, &items[i])
	}
	if revision > 0 {
		var desired *v1.PodTemplateSpec
		for i := range history {
			rc := history[i]
			if appsutil.DeploymentVersionFor(rc) == revision {
				desired = rc.Spec.Template
				break
			}
		}
		if desired == nil {
			return "", fmt.Errorf("unable to find the specified revision")
		}
		buf := bytes.NewBuffer([]byte{})
		versioned.DescribePodTemplate(desired, versioned.NewPrefixWriter(buf))
		return buf.String(), nil
	}
	sort.Sort(appsutil.ByLatestVersionAsc(history))
	return tabbedString(func(out *tabwriter.Writer) error {
		fmt.Fprintf(out, "REVISION\tSTATUS\tCAUSE\n")
		for i := range history {
			rc := history[i]
			rev := appsutil.DeploymentVersionFor(rc)
			status := appsutil.AnnotationFor(rc, appsv1.DeploymentStatusAnnotation)
			cause := rc.Annotations[appsv1.DeploymentStatusReasonAnnotation]
			if len(cause) == 0 {
				cause = "<unknown>"
			}
			fmt.Fprintf(out, "%d\t%s\t%s\n", rev, status, cause)
		}
		return nil
	})
}
func tabbedString(f func(*tabwriter.Writer) error) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(tabwriter.Writer)
	buf := &bytes.Buffer{}
	out.Init(buf, 0, 8, 1, '\t', 0)
	err := f(out)
	if err != nil {
		return "", err
	}
	out.Flush()
	str := string(buf.String())
	return str, nil
}
