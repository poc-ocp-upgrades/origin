package describe

import (
	"bytes"
	"fmt"
	"io"
	"net/url"
	"strings"
	"text/tabwriter"
	"time"
	"github.com/docker/go-units"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/rest"
	api "k8s.io/kubernetes/pkg/apis/core"
	buildv1 "github.com/openshift/api/build/v1"
	"github.com/openshift/library-go/pkg/image/reference"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
	buildmanualclient "github.com/openshift/origin/pkg/build/client/v1"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	templateapi "github.com/openshift/origin/pkg/template/apis/template"
)

const emptyString = "<none>"

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
func toString(v interface{}) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	value := fmt.Sprintf("%v", v)
	if len(value) == 0 {
		value = emptyString
	}
	return value
}
func bold(v interface{}) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "\033[1m" + toString(v) + "\033[0m"
}
func convertEnv(env []corev1.EnvVar) map[string]string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result := make(map[string]string, len(env))
	for _, e := range env {
		result[e.Name] = toString(e.Value)
	}
	return result
}
func formatEnv(env corev1.EnvVar) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if env.ValueFrom != nil && env.ValueFrom.FieldRef != nil {
		return fmt.Sprintf("%s=<%s>", env.Name, env.ValueFrom.FieldRef.FieldPath)
	}
	return fmt.Sprintf("%s=%s", env.Name, env.Value)
}
func formatString(out *tabwriter.Writer, label string, v interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	labelVals := strings.Split(toString(v), "\n")
	fmt.Fprintf(out, fmt.Sprintf("%s:", label))
	for _, lval := range labelVals {
		fmt.Fprintln(out, fmt.Sprintf("\t%s", lval))
	}
}
func formatTime(out *tabwriter.Writer, label string, t time.Time) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Fprintf(out, fmt.Sprintf("%s:\t%s ago\n", label, formatRelativeTime(t)))
}
func formatLabels(labelMap map[string]string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return labels.Set(labelMap).String()
}
func extractAnnotations(annotations map[string]string, keys ...string) ([]string, map[string]string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	extracted := make([]string, len(keys))
	remaining := make(map[string]string)
	for k, v := range annotations {
		remaining[k] = v
	}
	for i, key := range keys {
		extracted[i] = remaining[key]
		delete(remaining, key)
	}
	return extracted, remaining
}
func formatMapStringString(out *tabwriter.Writer, label string, items map[string]string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	keys := sets.NewString()
	for k := range items {
		keys.Insert(k)
	}
	if keys.Len() == 0 {
		formatString(out, label, "")
		return
	}
	for i, key := range keys.List() {
		if i == 0 {
			formatString(out, label, fmt.Sprintf("%s=%s", key, items[key]))
		} else {
			fmt.Fprintf(out, "%s\t%s=%s\n", "", key, items[key])
		}
	}
}
func formatAnnotations(out *tabwriter.Writer, m metav1.ObjectMeta, prefix string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	values, annotations := extractAnnotations(m.Annotations, "description")
	if len(values[0]) > 0 {
		formatString(out, prefix+"Description", values[0])
	}
	formatMapStringString(out, prefix+"Annotations", annotations)
}

var timeNowFn = func() time.Time {
	return time.Now()
}

func formatToHumanDuration(dur time.Duration) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return units.HumanDuration(dur)
}
func formatRelativeTime(t time.Time) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return units.HumanDuration(timeNowFn().Sub(t))
}
func FormatRelativeTime(t time.Time) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return formatRelativeTime(t)
}
func formatMeta(out *tabwriter.Writer, m metav1.ObjectMeta) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	formatString(out, "Name", m.Name)
	if len(m.Namespace) > 0 {
		formatString(out, "Namespace", m.Namespace)
	}
	if !m.CreationTimestamp.IsZero() {
		formatTime(out, "Created", m.CreationTimestamp.Time)
	}
	formatMapStringString(out, "Labels", m.Labels)
	formatAnnotations(out, m, "")
}

type DescribeWebhook struct {
	URL		string
	AllowEnv	*bool
}

func webHooksDescribe(triggers []buildv1.BuildTriggerPolicy, name, namespace string, c rest.Interface) map[string][]DescribeWebhook {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result := map[string][]DescribeWebhook{}
	for _, trigger := range triggers {
		var allowEnv *bool
		switch trigger.Type {
		case buildv1.GitHubWebHookBuildTriggerType, buildv1.GitLabWebHookBuildTriggerType, buildv1.BitbucketWebHookBuildTriggerType:
		case buildv1.GenericWebHookBuildTriggerType:
			allowEnv = &trigger.GenericWebHook.AllowEnv
		default:
			continue
		}
		webHookDesc := result[string(trigger.Type)]
		var urlStr string
		webhookClient := buildmanualclient.NewWebhookURLClient(c, namespace)
		u, err := webhookClient.WebHookURL(name, &trigger)
		if err != nil {
			urlStr = fmt.Sprintf("<error: %s>", err.Error())
		} else {
			urlStr, _ = url.PathUnescape(u.String())
		}
		webHookDesc = append(webHookDesc, DescribeWebhook{URL: urlStr, AllowEnv: allowEnv})
		result[string(trigger.Type)] = webHookDesc
	}
	return result
}
func formatImageStreamTags(out *tabwriter.Writer, stream *imageapi.ImageStream) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(stream.Status.Tags) == 0 && len(stream.Spec.Tags) == 0 {
		fmt.Fprintf(out, "Tags:\t<none>\n")
		return
	}
	now := timeNowFn()
	images := make(map[string]string)
	for tag, tags := range stream.Status.Tags {
		for _, item := range tags.Items {
			switch {
			case len(item.Image) > 0:
				if _, ok := images[item.Image]; !ok {
					images[item.Image] = tag
				}
			case len(item.DockerImageReference) > 0:
				if _, ok := images[item.DockerImageReference]; !ok {
					images[item.Image] = item.DockerImageReference
				}
			}
		}
	}
	sortedTags := []string{}
	for k := range stream.Status.Tags {
		sortedTags = append(sortedTags, k)
	}
	var localReferences sets.String
	var referentialTags map[string]sets.String
	for k := range stream.Spec.Tags {
		if target, _, multiple, err := imageapi.FollowTagReference(stream, k); err == nil && multiple {
			if referentialTags == nil {
				referentialTags = make(map[string]sets.String)
			}
			if localReferences == nil {
				localReferences = sets.NewString()
			}
			localReferences.Insert(k)
			v := referentialTags[target]
			if v == nil {
				v = sets.NewString()
				referentialTags[target] = v
			}
			v.Insert(k)
		}
		if _, ok := stream.Status.Tags[k]; !ok {
			sortedTags = append(sortedTags, k)
		}
	}
	fmt.Fprintf(out, "Unique Images:\t%d\nTags:\t%d\n\n", len(images), len(sortedTags))
	first := true
	imageapi.PrioritizeTags(sortedTags)
	for _, tag := range sortedTags {
		if localReferences.Has(tag) {
			continue
		}
		if first {
			first = false
		} else {
			fmt.Fprintf(out, "\n")
		}
		taglist, _ := stream.Status.Tags[tag]
		tagRef, hasSpecTag := stream.Spec.Tags[tag]
		scheduled := false
		insecure := false
		importing := false
		var name string
		if hasSpecTag && tagRef.From != nil {
			if len(tagRef.From.Namespace) > 0 && tagRef.From.Namespace != stream.Namespace {
				name = fmt.Sprintf("%s/%s", tagRef.From.Namespace, tagRef.From.Name)
			} else {
				name = tagRef.From.Name
			}
			scheduled, insecure = tagRef.ImportPolicy.Scheduled, tagRef.ImportPolicy.Insecure
			gen := imageapi.LatestObservedTagGeneration(stream, tag)
			importing = !tagRef.Reference && tagRef.Generation != nil && *tagRef.Generation > gen
		}
		if referentialTags[tag].Len() > 0 {
			references := referentialTags[tag].List()
			imageapi.PrioritizeTags(references)
			fmt.Fprintf(out, "%s (%s)\n", tag, strings.Join(references, ", "))
		} else {
			fmt.Fprintf(out, "%s\n", tag)
		}
		switch {
		case !hasSpecTag:
			fmt.Fprintf(out, "  no spec tag\n")
		case tagRef.From == nil:
			fmt.Fprintf(out, "  tag without source image\n")
		case tagRef.From.Kind == "ImageStreamTag":
			switch {
			case tagRef.Reference:
				fmt.Fprintf(out, "  reference to %s\n", name)
			case scheduled:
				fmt.Fprintf(out, "  updates automatically from %s\n", name)
			default:
				fmt.Fprintf(out, "  tagged from %s\n", name)
			}
		case tagRef.From.Kind == "DockerImage":
			switch {
			case tagRef.Reference:
				fmt.Fprintf(out, "  reference to registry %s\n", name)
			case scheduled:
				fmt.Fprintf(out, "  updates automatically from registry %s\n", name)
			default:
				fmt.Fprintf(out, "  tagged from %s\n", name)
			}
		case tagRef.From.Kind == "ImageStreamImage":
			switch {
			case tagRef.Reference:
				fmt.Fprintf(out, "  reference to image %s\n", name)
			default:
				fmt.Fprintf(out, "  tagged from %s\n", name)
			}
		default:
			switch {
			case tagRef.Reference:
				fmt.Fprintf(out, "  reference to %s %s\n", tagRef.From.Kind, name)
			default:
				fmt.Fprintf(out, "  updates from %s %s\n", tagRef.From.Kind, name)
			}
		}
		if insecure {
			fmt.Fprintf(out, "    will use insecure HTTPS or HTTP connections\n")
		}
		switch tagRef.ReferencePolicy.Type {
		case imageapi.LocalTagReferencePolicy:
			fmt.Fprintf(out, "    prefer registry pullthrough when referencing this tag\n")
		}
		fmt.Fprintln(out)
		extraOutput := false
		if d := tagRef.Annotations["description"]; len(d) > 0 {
			fmt.Fprintf(out, "  %s\n", d)
			extraOutput = true
		}
		if t := tagRef.Annotations["tags"]; len(t) > 0 {
			fmt.Fprintf(out, "  Tags: %s\n", strings.Join(strings.Split(t, ","), ", "))
			extraOutput = true
		}
		if t := tagRef.Annotations["supports"]; len(t) > 0 {
			fmt.Fprintf(out, "  Supports: %s\n", strings.Join(strings.Split(t, ","), ", "))
			extraOutput = true
		}
		if t := tagRef.Annotations["sampleRepo"]; len(t) > 0 {
			fmt.Fprintf(out, "  Example Repo: %s\n", t)
			extraOutput = true
		}
		if extraOutput {
			fmt.Fprintln(out)
		}
		if importing {
			fmt.Fprintf(out, "  ~ importing latest image ...\n")
		}
		for i := range taglist.Conditions {
			condition := &taglist.Conditions[i]
			switch condition.Type {
			case imageapi.ImportSuccess:
				if condition.Status == api.ConditionFalse {
					d := now.Sub(condition.LastTransitionTime.Time)
					fmt.Fprintf(out, "  ! error: Import failed (%s): %s\n      %s ago\n", condition.Reason, condition.Message, units.HumanDuration(d))
				}
			}
		}
		if len(taglist.Items) == 0 {
			continue
		}
		for i, event := range taglist.Items {
			d := now.Sub(event.Created.Time)
			if i == 0 {
				fmt.Fprintf(out, "  * %s\n", event.DockerImageReference)
			} else {
				fmt.Fprintf(out, "    %s\n", event.DockerImageReference)
			}
			ref, err := reference.Parse(event.DockerImageReference)
			id := event.Image
			if len(id) > 0 && err == nil && ref.ID != id {
				fmt.Fprintf(out, "      %s ago\t%s\n", units.HumanDuration(d), id)
			} else {
				fmt.Fprintf(out, "      %s ago\n", units.HumanDuration(d))
			}
		}
	}
}
func roleBindingRestrictionType(rbr *authorizationapi.RoleBindingRestriction) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch {
	case rbr.Spec.UserRestriction != nil:
		return "User"
	case rbr.Spec.GroupRestriction != nil:
		return "Group"
	case rbr.Spec.ServiceAccountRestriction != nil:
		return "ServiceAccount"
	}
	return ""
}
func PrintTemplateParameters(params []templateapi.Parameter, output io.Writer) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	w := tabwriter.NewWriter(output, 20, 5, 3, ' ', 0)
	defer w.Flush()
	parameterColumns := []string{"NAME", "DESCRIPTION", "GENERATOR", "VALUE"}
	fmt.Fprintf(w, "%s\n", strings.Join(parameterColumns, "\t"))
	for _, p := range params {
		value := p.Value
		if len(p.Generate) != 0 {
			value = p.From
		}
		_, err := fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", p.Name, p.Description, p.Generate, value)
		if err != nil {
			return err
		}
	}
	return nil
}
