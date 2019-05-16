package util

import (
	"bufio"
	"bytes"
	pkgerrors "github.com/pkg/errors"
	"io"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/util/errors"
	utilyaml "k8s.io/apimachinery/pkg/util/yaml"
	clientsetscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"sigs.k8s.io/yaml"
)

func MarshalToYaml(obj runtime.Object, gv schema.GroupVersion) ([]byte, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return MarshalToYamlForCodecs(obj, gv, clientsetscheme.Codecs)
}
func MarshalToYamlForCodecs(obj runtime.Object, gv schema.GroupVersion, codecs serializer.CodecFactory) ([]byte, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	mediaType := "application/yaml"
	info, ok := runtime.SerializerInfoForMediaType(codecs.SupportedMediaTypes(), mediaType)
	if !ok {
		return []byte{}, pkgerrors.Errorf("unsupported media type %q", mediaType)
	}
	encoder := codecs.EncoderForVersion(info.Serializer, gv)
	return runtime.Encode(encoder, obj)
}
func UnmarshalFromYaml(buffer []byte, gv schema.GroupVersion) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return UnmarshalFromYamlForCodecs(buffer, gv, clientsetscheme.Codecs)
}
func UnmarshalFromYamlForCodecs(buffer []byte, gv schema.GroupVersion, codecs serializer.CodecFactory) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	mediaType := "application/yaml"
	info, ok := runtime.SerializerInfoForMediaType(codecs.SupportedMediaTypes(), mediaType)
	if !ok {
		return nil, pkgerrors.Errorf("unsupported media type %q", mediaType)
	}
	decoder := codecs.DecoderToVersion(info.Serializer, gv)
	return runtime.Decode(decoder, buffer)
}
func SplitYAMLDocuments(yamlBytes []byte) (map[schema.GroupVersionKind][]byte, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	gvkmap := map[schema.GroupVersionKind][]byte{}
	knownKinds := map[string]bool{}
	errs := []error{}
	buf := bytes.NewBuffer(yamlBytes)
	reader := utilyaml.NewYAMLReader(bufio.NewReader(buf))
	for {
		typeMetaInfo := runtime.TypeMeta{}
		b, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		if len(b) == 0 {
			break
		}
		if err := yaml.Unmarshal(b, &typeMetaInfo); err != nil {
			return nil, err
		}
		if len(typeMetaInfo.APIVersion) == 0 || len(typeMetaInfo.Kind) == 0 {
			errs = append(errs, pkgerrors.New("invalid configuration: kind and apiVersion is mandatory information that needs to be specified in all YAML documents"))
			continue
		}
		if known := knownKinds[typeMetaInfo.Kind]; known {
			errs = append(errs, pkgerrors.Errorf("invalid configuration: kind %q is specified twice in YAML file", typeMetaInfo.Kind))
			continue
		}
		knownKinds[typeMetaInfo.Kind] = true
		gv, err := schema.ParseGroupVersion(typeMetaInfo.APIVersion)
		if err != nil {
			errs = append(errs, pkgerrors.Wrap(err, "unable to parse apiVersion"))
			continue
		}
		gvk := gv.WithKind(typeMetaInfo.Kind)
		gvkmap[gvk] = b
	}
	if err := errors.NewAggregate(errs); err != nil {
		return nil, err
	}
	return gvkmap, nil
}
func GroupVersionKindsFromBytes(b []byte) ([]schema.GroupVersionKind, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	gvkmap, err := SplitYAMLDocuments(b)
	if err != nil {
		return nil, err
	}
	gvks := []schema.GroupVersionKind{}
	for gvk := range gvkmap {
		gvks = append(gvks, gvk)
	}
	return gvks, nil
}
func GroupVersionKindsHasKind(gvks []schema.GroupVersionKind, kind string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, gvk := range gvks {
		if gvk.Kind == kind {
			return true
		}
	}
	return false
}
func GroupVersionKindsHasClusterConfiguration(gvks ...schema.GroupVersionKind) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return GroupVersionKindsHasKind(gvks, constants.ClusterConfigurationKind)
}
func GroupVersionKindsHasInitConfiguration(gvks ...schema.GroupVersionKind) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return GroupVersionKindsHasKind(gvks, constants.InitConfigurationKind)
}
func GroupVersionKindsHasJoinConfiguration(gvks ...schema.GroupVersionKind) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return GroupVersionKindsHasKind(gvks, constants.JoinConfigurationKind)
}
