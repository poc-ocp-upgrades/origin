package release

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"github.com/blang/semver"
	"github.com/ghodss/yaml"
	imageapi "github.com/openshift/api/image/v1"
	imagereference "github.com/openshift/origin/pkg/image/apis/image/reference"
	"k8s.io/klog"
)

type Payload struct {
	path		string
	references	*imageapi.ImageStream
}

func NewPayload(path string) *Payload {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Payload{path: path}
}
func (p *Payload) Path() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.path
}
func (p *Payload) Rewrite(allowTags bool, fn func(component string) imagereference.DockerImageReference) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	is, err := p.References()
	if err != nil {
		return err
	}
	replacements, err := ReplacementsForImageStream(is, allowTags, fn)
	if err != nil {
		return err
	}
	mapper, err := NewExactMapper(replacements)
	if err != nil {
		return err
	}
	files, err := ioutil.ReadDir(p.path)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if filepath.Base(file.Name()) == "image-references" {
			continue
		}
		path := filepath.Join(p.path, file.Name())
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		out, err := mapper(data)
		if err != nil {
			return fmt.Errorf("unable to rewrite the contents of %s: %v", path, err)
		}
		if bytes.Equal(data, out) {
			continue
		}
		klog.V(6).Infof("Rewrote\n%s\n\nto\n\n%s\n", string(data), string(out))
		if err := ioutil.WriteFile(path, out, file.Mode()); err != nil {
			return err
		}
	}
	return nil
}
func (p *Payload) References() (*imageapi.ImageStream, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if p.references != nil {
		return p.references, nil
	}
	is, err := parseImageStream(filepath.Join(p.path, "image-references"))
	if err != nil {
		return nil, err
	}
	p.references = is
	return is, nil
}
func parseImageStream(path string) (*imageapi.ImageStream, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	data, err := ioutil.ReadFile(path)
	if os.IsNotExist(err) {
		return nil, err
	}
	if err != nil {
		return nil, fmt.Errorf("unable to read release image info from release contents: %v", err)
	}
	return readReleaseImageReferences(data)
}
func readReleaseImageReferences(data []byte) (*imageapi.ImageStream, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	is := &imageapi.ImageStream{}
	if err := yaml.Unmarshal(data, &is); err != nil {
		return nil, fmt.Errorf("unable to load release image-references: %v", err)
	}
	if is.Kind != "ImageStream" || is.APIVersion != "image.openshift.io/v1" {
		return nil, fmt.Errorf("unrecognized image-references in release payload")
	}
	return is, nil
}

type ManifestMapper func(data []byte) ([]byte, error)

func NewTransformFromImageStreamFile(path string, input *imageapi.ImageStream, allowMissingImages bool) (ManifestMapper, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	is, err := parseImageStream(path)
	if err != nil {
		return nil, err
	}
	references := make(map[string]ImageReference)
	for _, tag := range is.Spec.Tags {
		if tag.From == nil || tag.From.Kind != "DockerImage" {
			continue
		}
		if len(tag.From.Name) == 0 {
			return nil, fmt.Errorf("no from.name for the tag %s", tag.Name)
		}
		ref := ImageReference{SourceRepository: tag.From.Name}
		for _, inputTag := range input.Spec.Tags {
			if inputTag.Name == tag.Name {
				ref.TargetPullSpec = inputTag.From.Name
				break
			}
		}
		if len(ref.TargetPullSpec) == 0 {
			if allowMissingImages {
				klog.V(2).Infof("Image file %q referenced an image %q that is not part of the input images, skipping", path, tag.From.Name)
				continue
			}
			return nil, fmt.Errorf("no input image tag named %q", tag.Name)
		}
		references[tag.Name] = ref
	}
	imageMapper, err := NewImageMapper(references)
	if err != nil {
		return nil, err
	}
	versions := make(map[string]string)
	tagsByName := make(map[string][]string)
	for _, tag := range input.Spec.Tags {
		if _, ok := references[tag.Name]; !ok {
			continue
		}
		value, ok := tag.Annotations[annotationBuildVersions]
		if !ok {
			continue
		}
		klog.V(4).Infof("Found build versions from %s: %s", tag.Name, value)
		items, err := parseComponentVersionsLabel(value)
		if err != nil {
			return nil, fmt.Errorf("input image stream has an invalid version annotation for tag %q: %v", tag.Name, value)
		}
		for k, v := range items {
			existing, ok := versions[k]
			if ok {
				if existing != v {
					return nil, fmt.Errorf("input image stream has multiple versions defined for version %s: %s defines %s but was already set to %s on %s", k, tag.Name, v, existing, strings.Join(tagsByName[k], ", "))
				}
			} else {
				versions[k] = v
				klog.V(4).Infof("Found version %s=%s from %s", k, v, tag.Name)
			}
			tagsByName[k] = append(tagsByName[k], tag.Name)
		}
	}
	defaults, err := parseComponentVersionsLabel(input.Annotations[annotationBuildVersions])
	if err != nil {
		return nil, fmt.Errorf("unable to read default versions label on input image stream: %v", err)
	}
	for k, v := range defaults {
		if _, ok := versions[k]; !ok {
			versions[k] = v
		}
	}
	versionMapper := NewComponentVersionsMapper(input.Name, versions, tagsByName)
	return func(data []byte) ([]byte, error) {
		data, err := imageMapper(data)
		if err != nil {
			return nil, err
		}
		return versionMapper(data)
	}, nil
}

type ImageReference struct {
	SourceRepository	string
	TargetPullSpec		string
}

func NopManifestMapper(data []byte) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return data, nil
}

const patternImageFormat = `([\W]|^)(%s)(:[\w][\w.-]{0,127}|@[A-Za-z][A-Za-z0-9]*(?:[-_+.][A-Za-z][A-Za-z0-9]*)*[:][[:xdigit:]]{2,})?([\s"']|$)`

func NewImageMapper(images map[string]ImageReference) (ManifestMapper, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	repositories := make([]string, 0, len(images))
	bySource := make(map[string]string)
	for name, ref := range images {
		if len(ref.SourceRepository) == 0 {
			return nil, fmt.Errorf("an empty source repository is not allowed for name %q", name)
		}
		if existing, ok := bySource[ref.SourceRepository]; ok {
			return nil, fmt.Errorf("the source repository %q was defined more than once (for %q and %q)", ref.SourceRepository, existing, name)
		}
		bySource[ref.SourceRepository] = name
		repositories = append(repositories, regexp.QuoteMeta(ref.SourceRepository))
	}
	if len(repositories) == 0 {
		klog.V(5).Infof("No images are mapped, will not replace any contents")
		return NopManifestMapper, nil
	}
	pattern := fmt.Sprintf(patternImageFormat, strings.Join(repositories, "|"))
	re := regexp.MustCompile(pattern)
	return func(data []byte) ([]byte, error) {
		out := re.ReplaceAllFunc(data, func(in []byte) []byte {
			parts := re.FindSubmatch(in)
			repository := string(parts[2])
			name, ok := bySource[repository]
			if !ok {
				klog.V(4).Infof("found potential image %q, but no matching definition", repository)
				return in
			}
			ref := images[name]
			suffix := parts[3]
			klog.V(2).Infof("found repository %q with locator %q in the input, switching to %q (from pattern %s)", string(repository), string(suffix), ref.TargetPullSpec, pattern)
			switch {
			case len(suffix) == 0:
				return []byte(string(parts[1]) + ref.TargetPullSpec + string(parts[4]))
			case suffix[0] == '@':
				return []byte(string(parts[1]) + ref.TargetPullSpec + string(parts[4]))
			default:
				return []byte(string(parts[1]) + ref.TargetPullSpec + string(parts[4]))
			}
		})
		return out, nil
	}, nil
}

const exactImageFormat = `\b%s\b`

func NewExactMapper(mappings map[string]string) (ManifestMapper, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	patterns := make(map[string]*regexp.Regexp)
	for from, to := range mappings {
		pattern := fmt.Sprintf(exactImageFormat, regexp.QuoteMeta(from))
		re, err := regexp.Compile(pattern)
		if err != nil {
			return nil, err
		}
		patterns[to] = re
	}
	return func(data []byte) ([]byte, error) {
		for to, pattern := range patterns {
			data = pattern.ReplaceAll(data, []byte(to))
		}
		return data, nil
	}, nil
}
func ComponentReferencesForImageStream(is *imageapi.ImageStream) (func(string) imagereference.DockerImageReference, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	components := make(map[string]imagereference.DockerImageReference)
	for _, tag := range is.Spec.Tags {
		if tag.From == nil || tag.From.Kind != "DockerImage" {
			continue
		}
		ref, err := imagereference.Parse(tag.From.Name)
		if err != nil {
			return nil, fmt.Errorf("reference for %q is invalid: %v", tag.Name, err)
		}
		components[tag.Name] = ref
	}
	return func(component string) imagereference.DockerImageReference {
		ref, ok := components[component]
		if !ok {
			panic(fmt.Errorf("unknown component %s", component))
		}
		return ref
	}, nil
}

const (
	componentVersionFormat = `([\W]|^)0\.0\.1-snapshot([a-z0-9\-]*)`
)

func NewComponentVersionsMapper(releaseName string, versions map[string]string, tagsByName map[string][]string) ManifestMapper {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if v, err := semver.Parse(releaseName); err == nil {
		v.Build = nil
		releaseName = v.String()
	} else {
		releaseName = ""
	}
	re, err := regexp.Compile(componentVersionFormat)
	if err != nil {
		return func([]byte) ([]byte, error) {
			return nil, fmt.Errorf("component versions mapper regex: %v", err)
		}
	}
	return func(data []byte) ([]byte, error) {
		var missing []string
		var conflicts []string
		data = re.ReplaceAllFunc(data, func(part []byte) []byte {
			matches := re.FindSubmatch(part)
			if matches == nil {
				return part
			}
			key := string(matches[2])
			if len(key) == 0 && len(releaseName) > 0 {
				buf := &bytes.Buffer{}
				buf.Write(matches[1])
				buf.WriteString(releaseName)
				return buf.Bytes()
			}
			if !strings.HasPrefix(key, "-") {
				return part
			}
			key = key[1:]
			value, ok := versions[key]
			if !ok {
				missing = append(missing, key)
				return part
			}
			if len(tagsByName[key]) > 1 {
				conflicts = append(conflicts, key)
				return part
			}
			buf := &bytes.Buffer{}
			buf.Write(matches[1])
			buf.WriteString(value)
			return buf.Bytes()
		})
		if len(missing) > 0 {
			switch len(missing) {
			case 1:
				if len(missing[0]) == 0 {
					return nil, fmt.Errorf("empty version references are not allowed")
				}
				return nil, fmt.Errorf("unknown version reference %q", missing[0])
			default:
				return nil, fmt.Errorf("unknown version references: %s", strings.Join(missing, ", "))
			}
		}
		if len(conflicts) > 0 {
			allImageTags := tagsByName[conflicts[0]]
			sort.Strings(allImageTags)
			return nil, fmt.Errorf("the version for %q is inconsistent across the referenced images: %s", conflicts[0], strings.Join(allImageTags, ", "))
		}
		return data, nil
	}
}

var (
	reAllowedVersionKey = regexp.MustCompile(`^[a-z0-9]+[\-a-z0-9]*[a-z0-9]+$`)
)

type ComponentVersions map[string]string

func (v ComponentVersions) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var keys []string
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	buf := &bytes.Buffer{}
	for i, k := range keys {
		if i != 0 {
			buf.WriteRune(',')
		}
		fmt.Fprintf(buf, "%s=%s", k, v[k])
	}
	return buf.String()
}
func parseComponentVersionsLabel(label string) (ComponentVersions, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	label = strings.TrimSpace(label)
	if len(label) == 0 {
		return nil, nil
	}
	labels := make(map[string]string)
	for _, pair := range strings.Split(label, ",") {
		pair = strings.TrimSpace(pair)
		parts := strings.SplitN(pair, "=", 2)
		if len(parts) == 1 {
			return nil, fmt.Errorf("the version pair %q must be NAME=VERSION", pair)
		}
		if !reAllowedVersionKey.MatchString(parts[0]) {
			return nil, fmt.Errorf("the version name %q must only be ASCII alphanumerics and internal hyphens", parts[0])
		}
		v, err := semver.Parse(parts[1])
		if err != nil {
			return nil, fmt.Errorf("the version pair %q must have a valid semantic version: %v", pair, err)
		}
		v.Build = nil
		labels[parts[0]] = v.String()
	}
	return labels, nil
}
func ReplacementsForImageStream(is *imageapi.ImageStream, allowTags bool, fn func(component string) imagereference.DockerImageReference) (map[string]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	replacements := make(map[string]string)
	for i := range is.Spec.Tags {
		tag := &is.Spec.Tags[i]
		if tag.From == nil || tag.From.Kind != "DockerImage" {
			continue
		}
		oldImage := tag.From.Name
		oldRef, err := imagereference.Parse(oldImage)
		if err != nil {
			return nil, fmt.Errorf("unable to parse image reference for tag %q from payload: %v", tag.Name, err)
		}
		if len(oldRef.Tag) > 0 || len(oldRef.ID) == 0 {
			if !allowTags {
				return nil, fmt.Errorf("image reference tag %q in payload does not point to an image digest - unable to rewrite payload", tag.Name)
			}
		}
		ref := fn(tag.Name)
		if !allowTags {
			if len(ref.ID) == 0 {
				ref.Tag = ""
				ref.ID = oldRef.ID
			}
		}
		newImage := ref.Exact()
		replacements[oldImage] = newImage
		tag.From.Name = newImage
	}
	if klog.V(5) {
		for k, v := range replacements {
			klog.Infof("Mapping %s -> %s", k, v)
		}
	}
	return replacements, nil
}
