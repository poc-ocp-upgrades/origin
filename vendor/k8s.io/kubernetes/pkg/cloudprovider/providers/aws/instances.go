package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"k8s.io/api/core/v1"
	"k8s.io/klog"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"time"
)

var awsInstanceRegMatch = regexp.MustCompile("^i-[^/]*$")

type awsInstanceID string

func (i awsInstanceID) awsString() *string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return aws.String(string(i))
}

type kubernetesInstanceID string

func (name kubernetesInstanceID) mapToAWSInstanceID() (awsInstanceID, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	s := string(name)
	if !strings.HasPrefix(s, "aws://") {
		s = "aws://" + "/" + "/" + s
	}
	url, err := url.Parse(s)
	if err != nil {
		return "", fmt.Errorf("Invalid instance name (%s): %v", name, err)
	}
	if url.Scheme != "aws" {
		return "", fmt.Errorf("Invalid scheme for AWS instance (%s)", name)
	}
	awsID := ""
	tokens := strings.Split(strings.Trim(url.Path, "/"), "/")
	if len(tokens) == 1 {
		awsID = tokens[0]
	} else if len(tokens) == 2 {
		awsID = tokens[1]
	}
	if awsID == "" || !awsInstanceRegMatch.MatchString(awsID) {
		return "", fmt.Errorf("Invalid format for AWS instance (%s)", name)
	}
	return awsInstanceID(awsID), nil
}
func mapToAWSInstanceIDs(nodes []*v1.Node) ([]awsInstanceID, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var instanceIDs []awsInstanceID
	for _, node := range nodes {
		if node.Spec.ProviderID == "" {
			return nil, fmt.Errorf("node %q did not have ProviderID set", node.Name)
		}
		instanceID, err := kubernetesInstanceID(node.Spec.ProviderID).mapToAWSInstanceID()
		if err != nil {
			return nil, fmt.Errorf("unable to parse ProviderID %q for node %q", node.Spec.ProviderID, node.Name)
		}
		instanceIDs = append(instanceIDs, instanceID)
	}
	return instanceIDs, nil
}
func mapToAWSInstanceIDsTolerant(nodes []*v1.Node) []awsInstanceID {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var instanceIDs []awsInstanceID
	for _, node := range nodes {
		if node.Spec.ProviderID == "" {
			klog.Warningf("node %q did not have ProviderID set", node.Name)
			continue
		}
		instanceID, err := kubernetesInstanceID(node.Spec.ProviderID).mapToAWSInstanceID()
		if err != nil {
			klog.Warningf("unable to parse ProviderID %q for node %q", node.Spec.ProviderID, node.Name)
			continue
		}
		instanceIDs = append(instanceIDs, instanceID)
	}
	return instanceIDs
}
func describeInstance(ec2Client EC2, instanceID awsInstanceID) (*ec2.Instance, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	request := &ec2.DescribeInstancesInput{InstanceIds: []*string{instanceID.awsString()}}
	instances, err := ec2Client.DescribeInstances(request)
	if err != nil {
		return nil, err
	}
	if len(instances) == 0 {
		return nil, fmt.Errorf("no instances found for instance: %s", instanceID)
	}
	if len(instances) > 1 {
		return nil, fmt.Errorf("multiple instances found for instance: %s", instanceID)
	}
	return instances[0], nil
}

type instanceCache struct {
	cloud    *Cloud
	mutex    sync.Mutex
	snapshot *allInstancesSnapshot
}

func (c *instanceCache) describeAllInstancesUncached() (*allInstancesSnapshot, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	now := time.Now()
	klog.V(4).Infof("EC2 DescribeInstances - fetching all instances")
	filters := []*ec2.Filter{}
	instances, err := c.cloud.describeInstances(filters)
	if err != nil {
		return nil, err
	}
	m := make(map[awsInstanceID]*ec2.Instance)
	for _, i := range instances {
		id := awsInstanceID(aws.StringValue(i.InstanceId))
		m[id] = i
	}
	snapshot := &allInstancesSnapshot{now, m}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.snapshot != nil && snapshot.olderThan(c.snapshot) {
		klog.Infof("Not caching concurrent AWS DescribeInstances results")
	} else {
		c.snapshot = snapshot
	}
	return snapshot, nil
}

type cacheCriteria struct {
	MaxAge       time.Duration
	HasInstances []awsInstanceID
}

func (c *instanceCache) describeAllInstancesCached(criteria cacheCriteria) (*allInstancesSnapshot, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var err error
	snapshot := c.getSnapshot()
	if snapshot != nil && !snapshot.MeetsCriteria(criteria) {
		snapshot = nil
	}
	if snapshot == nil {
		snapshot, err = c.describeAllInstancesUncached()
		if err != nil {
			return nil, err
		}
	} else {
		klog.V(6).Infof("EC2 DescribeInstances - using cached results")
	}
	return snapshot, nil
}
func (c *instanceCache) getSnapshot() *allInstancesSnapshot {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.snapshot
}
func (s *allInstancesSnapshot) olderThan(other *allInstancesSnapshot) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return other.timestamp.After(s.timestamp)
}
func (s *allInstancesSnapshot) MeetsCriteria(criteria cacheCriteria) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if criteria.MaxAge > 0 {
		now := time.Now()
		if now.Sub(s.timestamp) > criteria.MaxAge {
			klog.V(6).Infof("instanceCache snapshot cannot be used as is older than MaxAge=%s", criteria.MaxAge)
			return false
		}
	}
	if len(criteria.HasInstances) != 0 {
		for _, id := range criteria.HasInstances {
			if nil == s.instances[id] {
				klog.V(6).Infof("instanceCache snapshot cannot be used as does not contain instance %s", id)
				return false
			}
		}
	}
	return true
}

type allInstancesSnapshot struct {
	timestamp time.Time
	instances map[awsInstanceID]*ec2.Instance
}

func (s *allInstancesSnapshot) FindInstances(ids []awsInstanceID) map[awsInstanceID]*ec2.Instance {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	m := make(map[awsInstanceID]*ec2.Instance)
	for _, id := range ids {
		instance := s.instances[id]
		if instance != nil {
			m[id] = instance
		}
	}
	return m
}
