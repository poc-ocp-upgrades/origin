package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog"
	"strings"
)

const TagNameKubernetesClusterPrefix = "kubernetes.io/cluster/"
const TagNameKubernetesClusterLegacy = "KubernetesCluster"

type ResourceLifecycle string

const (
	ResourceLifecycleOwned  = "owned"
	ResourceLifecycleShared = "shared"
)

type awsTagging struct {
	ClusterID      string
	usesLegacyTags bool
}

func (t *awsTagging) init(legacyClusterID string, clusterID string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if legacyClusterID != "" {
		if clusterID != "" && legacyClusterID != clusterID {
			return fmt.Errorf("ClusterID tags did not match: %q vs %q", clusterID, legacyClusterID)
		}
		t.usesLegacyTags = true
		clusterID = legacyClusterID
	}
	t.ClusterID = clusterID
	if clusterID != "" {
		klog.Infof("AWS cloud filtering on ClusterID: %v", clusterID)
	} else {
		return fmt.Errorf("AWS cloud failed to find ClusterID")
	}
	return nil
}
func (t *awsTagging) initFromTags(tags []*ec2.Tag) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	legacyClusterID, newClusterID, err := findClusterIDs(tags)
	if err != nil {
		return err
	}
	if legacyClusterID == "" && newClusterID == "" {
		klog.Errorf("Tag %q nor %q not found; Kubernetes may behave unexpectedly.", TagNameKubernetesClusterLegacy, TagNameKubernetesClusterPrefix+"...")
	}
	return t.init(legacyClusterID, newClusterID)
}
func findClusterIDs(tags []*ec2.Tag) (string, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	legacyClusterID := ""
	newClusterID := ""
	for _, tag := range tags {
		tagKey := aws.StringValue(tag.Key)
		if strings.HasPrefix(tagKey, TagNameKubernetesClusterPrefix) {
			id := strings.TrimPrefix(tagKey, TagNameKubernetesClusterPrefix)
			if newClusterID != "" {
				return "", "", fmt.Errorf("Found multiple cluster tags with prefix %s (%q and %q)", TagNameKubernetesClusterPrefix, newClusterID, id)
			}
			newClusterID = id
		}
		if tagKey == TagNameKubernetesClusterLegacy {
			id := aws.StringValue(tag.Value)
			if legacyClusterID != "" {
				return "", "", fmt.Errorf("Found multiple %s tags (%q and %q)", TagNameKubernetesClusterLegacy, legacyClusterID, id)
			}
			legacyClusterID = id
		}
	}
	return legacyClusterID, newClusterID, nil
}
func (t *awsTagging) clusterTagKey() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return TagNameKubernetesClusterPrefix + t.ClusterID
}
func (t *awsTagging) hasClusterTag(tags []*ec2.Tag) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(t.ClusterID) == 0 {
		return true
	}
	clusterTagKey := t.clusterTagKey()
	for _, tag := range tags {
		tagKey := aws.StringValue(tag.Key)
		if (tagKey == TagNameKubernetesClusterLegacy) && (aws.StringValue(tag.Value) == t.ClusterID) {
			return true
		}
		if tagKey == clusterTagKey {
			return true
		}
	}
	return false
}
func (t *awsTagging) readRepairClusterTags(client EC2, resourceID string, lifecycle ResourceLifecycle, additionalTags map[string]string, observedTags []*ec2.Tag) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	actualTagMap := make(map[string]string)
	for _, tag := range observedTags {
		actualTagMap[aws.StringValue(tag.Key)] = aws.StringValue(tag.Value)
	}
	expectedTags := t.buildTags(lifecycle, additionalTags)
	addTags := make(map[string]string)
	for k, expected := range expectedTags {
		actual := actualTagMap[k]
		if actual == expected {
			continue
		}
		if actual == "" {
			klog.Warningf("Resource %q was missing expected cluster tag %q.  Will add (with value %q)", resourceID, k, expected)
			addTags[k] = expected
		} else {
			return fmt.Errorf("resource %q has tag belonging to another cluster: %q=%q (expected %q)", resourceID, k, actual, expected)
		}
	}
	if len(addTags) == 0 {
		return nil
	}
	if err := t.createTags(client, resourceID, lifecycle, addTags); err != nil {
		return fmt.Errorf("error adding missing tags to resource %q: %q", resourceID, err)
	}
	return nil
}
func (t *awsTagging) createTags(client EC2, resourceID string, lifecycle ResourceLifecycle, additionalTags map[string]string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	tags := t.buildTags(lifecycle, additionalTags)
	if tags == nil || len(tags) == 0 {
		return nil
	}
	var awsTags []*ec2.Tag
	for k, v := range tags {
		tag := &ec2.Tag{Key: aws.String(k), Value: aws.String(v)}
		awsTags = append(awsTags, tag)
	}
	backoff := wait.Backoff{Duration: createTagInitialDelay, Factor: createTagFactor, Steps: createTagSteps}
	request := &ec2.CreateTagsInput{}
	request.Resources = []*string{&resourceID}
	request.Tags = awsTags
	var lastErr error
	err := wait.ExponentialBackoff(backoff, func() (bool, error) {
		_, err := client.CreateTags(request)
		if err == nil {
			return true, nil
		}
		klog.V(2).Infof("Failed to create tags; will retry.  Error was %q", err)
		lastErr = err
		return false, nil
	})
	if err == wait.ErrWaitTimeout {
		err = lastErr
	}
	return err
}
func (t *awsTagging) addFilters(filters []*ec2.Filter) []*ec2.Filter {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(t.ClusterID) == 0 {
		if len(filters) == 0 {
			return nil
		}
		return filters
	}
	f := newEc2Filter("tag-key", TagNameKubernetesClusterLegacy, t.clusterTagKey())
	filters = append(filters, f)
	return filters
}
func (t *awsTagging) buildTags(lifecycle ResourceLifecycle, additionalTags map[string]string) map[string]string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	tags := make(map[string]string)
	for k, v := range additionalTags {
		tags[k] = v
	}
	if len(t.ClusterID) == 0 {
		return tags
	}
	if t.usesLegacyTags {
		tags[TagNameKubernetesClusterLegacy] = t.ClusterID
	}
	tags[t.clusterTagKey()] = string(lifecycle)
	return tags
}
func (t *awsTagging) clusterID() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return t.ClusterID
}
