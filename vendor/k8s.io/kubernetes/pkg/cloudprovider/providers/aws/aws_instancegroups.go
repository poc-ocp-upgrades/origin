package aws

import (
 "fmt"
 "github.com/aws/aws-sdk-go/aws"
 "github.com/aws/aws-sdk-go/service/autoscaling"
 "k8s.io/klog"
)

var _ InstanceGroups = &Cloud{}

func ResizeInstanceGroup(asg ASG, instanceGroupName string, size int) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 request := &autoscaling.UpdateAutoScalingGroupInput{AutoScalingGroupName: aws.String(instanceGroupName), MinSize: aws.Int64(int64(size)), MaxSize: aws.Int64(int64(size))}
 if _, err := asg.UpdateAutoScalingGroup(request); err != nil {
  return fmt.Errorf("error resizing AWS autoscaling group: %q", err)
 }
 return nil
}
func (c *Cloud) ResizeInstanceGroup(instanceGroupName string, size int) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return ResizeInstanceGroup(c.asg, instanceGroupName, size)
}
func DescribeInstanceGroup(asg ASG, instanceGroupName string) (InstanceGroupInfo, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 request := &autoscaling.DescribeAutoScalingGroupsInput{AutoScalingGroupNames: []*string{aws.String(instanceGroupName)}}
 response, err := asg.DescribeAutoScalingGroups(request)
 if err != nil {
  return nil, fmt.Errorf("error listing AWS autoscaling group (%s): %q", instanceGroupName, err)
 }
 if len(response.AutoScalingGroups) == 0 {
  return nil, nil
 }
 if len(response.AutoScalingGroups) > 1 {
  klog.Warning("AWS returned multiple autoscaling groups with name ", instanceGroupName)
 }
 group := response.AutoScalingGroups[0]
 return &awsInstanceGroup{group: group}, nil
}
func (c *Cloud) DescribeInstanceGroup(instanceGroupName string) (InstanceGroupInfo, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return DescribeInstanceGroup(c.asg, instanceGroupName)
}

var _ InstanceGroupInfo = &awsInstanceGroup{}

type awsInstanceGroup struct{ group *autoscaling.Group }

func (g *awsInstanceGroup) CurrentSize() (int, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return len(g.group.Instances), nil
}
