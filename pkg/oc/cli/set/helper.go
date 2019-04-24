package set

import (
	"fmt"
	"strings"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	"k8s.io/cli-runtime/pkg/genericclioptions/resource"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/kubectl/scheme"
)

func selectContainers(containers []corev1.Container, spec string) ([]*corev1.Container, []*corev1.Container) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := []*corev1.Container{}
	skipped := []*corev1.Container{}
	for i, c := range containers {
		if selectString(c.Name, spec) {
			out = append(out, &containers[i])
		} else {
			skipped = append(skipped, &containers[i])
		}
	}
	return out, skipped
}
func selectString(s, spec string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if spec == "*" {
		return true
	}
	if !strings.Contains(spec, "*") {
		return s == spec
	}
	pos := 0
	match := true
	parts := strings.Split(spec, "*")
	for i, part := range parts {
		if len(part) == 0 {
			continue
		}
		next := strings.Index(s[pos:], part)
		switch {
		case next < pos:
			fallthrough
		case i == 0 && pos != 0:
			fallthrough
		case i == (len(parts)-1) && len(s) != (len(part)+next):
			match = false
			break
		default:
			pos = next
		}
	}
	return match
}
func updateEnv(existing []corev1.EnvVar, env []corev1.EnvVar, remove []string) []corev1.EnvVar {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := []corev1.EnvVar{}
	covered := sets.NewString(remove...)
	for _, e := range existing {
		if covered.Has(e.Name) {
			continue
		}
		newer, ok := findEnv(env, e.Name)
		if ok {
			covered.Insert(e.Name)
			out = append(out, newer)
			continue
		}
		out = append(out, e)
	}
	for _, e := range env {
		if covered.Has(e.Name) {
			continue
		}
		covered.Insert(e.Name)
		out = append(out, e)
	}
	return out
}
func findEnv(env []corev1.EnvVar, name string) (corev1.EnvVar, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, e := range env {
		if e.Name == name {
			return e, true
		}
	}
	return corev1.EnvVar{}, false
}

type Patch struct {
	Info	*resource.Info
	Err	error
	Before	[]byte
	After	[]byte
	Patch	[]byte
}

func CalculatePatches(infos []*resource.Info, encoder runtime.Encoder, mutateFn func(*resource.Info) (bool, error)) []*Patch {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var patches []*Patch
	for _, info := range infos {
		patch := &Patch{Info: info}
		versionedEncoder := legacyscheme.Codecs.EncoderForVersion(encoder, patch.Info.Mapping.GroupVersionKind.GroupVersion())
		patch.Before, patch.Err = runtime.Encode(versionedEncoder, info.Object)
		ok, err := mutateFn(info)
		if !ok {
			continue
		}
		if err != nil {
			patch.Err = err
		}
		patches = append(patches, patch)
		if patch.Err != nil {
			continue
		}
		patch.After, patch.Err = runtime.Encode(versionedEncoder, info.Object)
		if patch.Err != nil {
			continue
		}
		versioned, err := legacyscheme.Scheme.ConvertToVersion(info.Object, info.Mapping.GroupVersionKind.GroupVersion())
		if err != nil {
			patch.Err = err
			continue
		}
		patch.Patch, patch.Err = strategicpatch.CreateTwoWayMergePatch(patch.Before, patch.After, versioned)
	}
	return patches
}
func CalculatePatchesExternal(infos []*resource.Info, mutateFn func(*resource.Info) (bool, error)) []*Patch {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var patches []*Patch
	for _, info := range infos {
		patch := &Patch{Info: info}
		patch.Before, patch.Err = runtime.Encode(scheme.DefaultJSONEncoder(), info.Object)
		ok, err := mutateFn(info)
		if !ok {
			continue
		}
		if err != nil {
			patch.Err = err
		}
		patches = append(patches, patch)
		if patch.Err != nil {
			continue
		}
		patch.After, patch.Err = runtime.Encode(scheme.DefaultJSONEncoder(), info.Object)
		if patch.Err != nil {
			continue
		}
		patch.Patch, patch.Err = strategicpatch.CreateTwoWayMergePatch(patch.Before, patch.After, info.Object)
	}
	return patches
}
func getObjectName(info *resource.Info) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if info.Mapping != nil {
		return fmt.Sprintf("%s/%s", info.Mapping.Resource.Resource, info.Name)
	}
	gvk := info.Object.GetObjectKind().GroupVersionKind()
	if len(gvk.Group) == 0 {
		return fmt.Sprintf("%s/%s", strings.ToLower(gvk.Kind), info.Name)
	}
	return fmt.Sprintf("%s.%s/%s\n", strings.ToLower(gvk.Kind), gvk.Group, info.Name)
}
