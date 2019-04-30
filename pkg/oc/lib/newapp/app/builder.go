package app

import (
	"fmt"
	"strings"
	"k8s.io/apimachinery/pkg/util/errors"
	dockerv10 "github.com/openshift/api/image/docker10"
	imagev1 "github.com/openshift/api/image/v1"
	imageutil "github.com/openshift/origin/pkg/image/util"
)

var s2iEnvironmentNames = []string{"STI_LOCATION", "STI_SCRIPTS_URL", "STI_BUILDER"}

const s2iScriptsLabel = "io.openshift.s2i.scripts-url"

func IsBuilderImage(image *dockerv10.DockerImage) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if image == nil || image.Config == nil {
		return false
	}
	if _, ok := image.Config.Labels[s2iScriptsLabel]; ok {
		return true
	}
	for _, env := range image.Config.Env {
		for _, name := range s2iEnvironmentNames {
			if strings.HasPrefix(env, name+"=") {
				return true
			}
		}
	}
	return false
}
func IsBuilderStreamTag(stream *imagev1.ImageStream, tag string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if stream == nil {
		return false
	}
	t, hasTag := imageutil.SpecHasTag(stream, tag)
	if !hasTag {
		return false
	}
	return imageutil.HasAnnotationTag(&t, "builder")
}
func IsBuilderMatch(match *ComponentMatch) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if match.DockerImage != nil && IsBuilderImage(match.DockerImage) {
		return true
	}
	if match.ImageStream != nil && IsBuilderStreamTag(match.ImageStream, match.ImageTag) {
		return true
	}
	return false
}
func isGeneratorJobImage(image *dockerv10.DockerImage) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if image == nil || image.Config == nil {
		return false
	}
	if image.Config.Labels[labelGenerateJob] == "true" {
		return true
	}
	return false
}
func isGeneratorJobImageStreamTag(stream *imagev1.ImageStream, tag string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if stream == nil {
		return false
	}
	t, hasTag := imageutil.SpecHasTag(stream, tag)
	if !hasTag {
		return false
	}
	return t.Annotations[labelGenerateJob] == "true"
}
func parseGenerateTokenAs(value string) (*TokenInput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	parts := strings.SplitN(value, ":", 2)
	switch parts[0] {
	case "env":
		if len(parts) != 2 {
			return nil, fmt.Errorf("label %s=%s; expected 'env:<NAME>' or not set", labelGenerateTokenAs, value)
		}
		name := strings.TrimSpace(parts[1])
		if len(name) == 0 {
			return nil, fmt.Errorf("label %s=%s; expected 'env:<NAME>' but name was empty", labelGenerateTokenAs, value)
		}
		return &TokenInput{Env: &name}, nil
	case "file":
		if len(parts) != 2 {
			return nil, fmt.Errorf("label %s=%s; expected 'file:<PATH>' or not set", labelGenerateTokenAs, value)
		}
		name := strings.TrimSpace(parts[1])
		if len(name) == 0 {
			return nil, fmt.Errorf("label %s=%s; expected 'file:<PATH>' but path was empty", labelGenerateTokenAs, value)
		}
		return &TokenInput{File: &name}, nil
	case "serviceaccount":
		return &TokenInput{ServiceAccount: true}, nil
	default:
		return nil, fmt.Errorf("unrecognized value for label %s=%s; expected 'env:<NAME>', 'file:<PATH>', or 'serviceaccount'", labelGenerateTokenAs, value)
	}
}

const (
	labelGenerateJob	= "io.openshift.generate.job"
	labelGenerateTokenAs	= "io.openshift.generate.token.as"
)

type TokenInput struct {
	Env		*string
	File		*string
	ServiceAccount	bool
}
type GeneratorInput struct {
	Job	bool
	Token	*TokenInput
}

func GeneratorInputFromMatch(match *ComponentMatch) (GeneratorInput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	input := GeneratorInput{}
	errs := []error{}
	if match.DockerImage != nil && match.DockerImage.Config != nil {
		input.Job = isGeneratorJobImage(match.DockerImage)
		if value, ok := match.DockerImage.Config.Labels[labelGenerateTokenAs]; ok {
			if token, err := parseGenerateTokenAs(value); err != nil {
				errs = append(errs, err)
			} else {
				input.Token = token
			}
		}
	}
	if match.ImageStream != nil {
		input.Job = isGeneratorJobImageStreamTag(match.ImageStream, match.ImageTag)
		if t, hasTag := imageutil.SpecHasTag(match.ImageStream, match.ImageTag); hasTag {
			if value, ok := t.Annotations[labelGenerateTokenAs]; ok {
				if token, err := parseGenerateTokenAs(value); err != nil {
					errs = append(errs, err)
				} else {
					input.Token = token
				}
			}
		}
	}
	return input, errors.NewAggregate(errs)
}
