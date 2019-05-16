package options

import (
	"github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"sort"
	"strings"
)

const (
	DefaultEtcdPathPrefix = "/registry"
)

type StorageSerializationOptions struct {
	StorageVersions        string
	DefaultStorageVersions string
}

func NewStorageSerializationOptions() *StorageSerializationOptions {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &StorageSerializationOptions{DefaultStorageVersions: ToPreferredVersionString(legacyscheme.Scheme.PreferredVersionAllGroups()), StorageVersions: ToPreferredVersionString(legacyscheme.Scheme.PreferredVersionAllGroups())}
}
func (s *StorageSerializationOptions) StorageGroupsToEncodingVersion() (map[string]schema.GroupVersion, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	storageVersionMap := map[string]schema.GroupVersion{}
	if err := mergeGroupVersionIntoMap(s.DefaultStorageVersions, storageVersionMap); err != nil {
		return nil, err
	}
	if err := mergeGroupVersionIntoMap(s.StorageVersions, storageVersionMap); err != nil {
		return nil, err
	}
	return storageVersionMap, nil
}
func mergeGroupVersionIntoMap(gvList string, dest map[string]schema.GroupVersion) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, gvString := range strings.Split(gvList, ",") {
		if gvString == "" {
			continue
		}
		if !strings.Contains(gvString, "=") {
			gv, err := schema.ParseGroupVersion(gvString)
			if err != nil {
				return err
			}
			dest[gv.Group] = gv
		} else {
			parts := strings.SplitN(gvString, "=", 2)
			gv, err := schema.ParseGroupVersion(parts[1])
			if err != nil {
				return err
			}
			dest[parts[0]] = gv
		}
	}
	return nil
}
func (s *StorageSerializationOptions) AddFlags(fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fs.StringVar(&s.StorageVersions, "storage-versions", s.StorageVersions, ""+"The per-group version to store resources in. "+"Specified in the format \"group1/version1,group2/version2,...\". "+"In the case where objects are moved from one group to the other, "+"you may specify the format \"group1=group2/v1beta1,group3/v1beta1,...\". "+"You only need to pass the groups you wish to change from the defaults. "+"It defaults to a list of preferred versions of all known groups.")
	fs.MarkDeprecated("storage-versions", ""+"Please omit this flag to ensure the default storage versions are used ."+"Otherwise the cluster is not safe to upgrade to a version newer than 1.12. "+"This flag will be removed in 1.13.")
}
func ToPreferredVersionString(versions []schema.GroupVersion) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var defaults []string
	for _, version := range versions {
		defaults = append(defaults, version.String())
	}
	sort.Strings(defaults)
	return strings.Join(defaults, ",")
}
