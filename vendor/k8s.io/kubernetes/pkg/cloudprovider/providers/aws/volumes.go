package aws

import (
 "fmt"
 "net/url"
 "regexp"
 "strings"
 "github.com/aws/aws-sdk-go/aws"
 "github.com/aws/aws-sdk-go/service/ec2"
 "k8s.io/klog"
 "k8s.io/apimachinery/pkg/types"
)

var awsVolumeRegMatch = regexp.MustCompile("^vol-[^/]*$")

type EBSVolumeID string

func (i EBSVolumeID) awsString() *string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return aws.String(string(i))
}

type KubernetesVolumeID string
type diskInfo struct {
 ec2Instance     *ec2.Instance
 nodeName        types.NodeName
 volumeState     string
 attachmentState string
 hasAttachment   bool
 disk            *awsDisk
}

func (name KubernetesVolumeID) MapToAWSVolumeID() (EBSVolumeID, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 s := string(name)
 if !strings.HasPrefix(s, "aws://") {
  s = "aws://" + "" + "/" + s
 }
 url, err := url.Parse(s)
 if err != nil {
  return "", fmt.Errorf("Invalid disk name (%s): %v", name, err)
 }
 if url.Scheme != "aws" {
  return "", fmt.Errorf("Invalid scheme for AWS volume (%s)", name)
 }
 awsID := url.Path
 awsID = strings.Trim(awsID, "/")
 if !awsVolumeRegMatch.MatchString(awsID) {
  return "", fmt.Errorf("Invalid format for AWS volume (%s)", name)
 }
 return EBSVolumeID(awsID), nil
}
func GetAWSVolumeID(kubeVolumeID string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 kid := KubernetesVolumeID(kubeVolumeID)
 awsID, err := kid.MapToAWSVolumeID()
 return string(awsID), err
}
func (c *Cloud) checkIfAttachedToNode(diskName KubernetesVolumeID, nodeName types.NodeName) (*diskInfo, bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 disk, err := newAWSDisk(c, diskName)
 if err != nil {
  return nil, true, err
 }
 awsDiskInfo := &diskInfo{disk: disk}
 info, err := disk.describeVolume()
 if err != nil {
  klog.Warningf("Error describing volume %s with %v", diskName, err)
  awsDiskInfo.volumeState = "unknown"
  return awsDiskInfo, false, err
 }
 awsDiskInfo.volumeState = aws.StringValue(info.State)
 if len(info.Attachments) > 0 {
  attachment := info.Attachments[0]
  awsDiskInfo.attachmentState = aws.StringValue(attachment.State)
  instanceID := aws.StringValue(attachment.InstanceId)
  instanceInfo, err := c.getInstanceByID(instanceID)
  if err != nil {
   fetchErr := fmt.Errorf("Error fetching instance %s for volume %s", instanceID, diskName)
   klog.Warning(fetchErr)
   return awsDiskInfo, false, fetchErr
  }
  awsDiskInfo.ec2Instance = instanceInfo
  awsDiskInfo.nodeName = mapInstanceToNodeName(instanceInfo)
  awsDiskInfo.hasAttachment = true
  if awsDiskInfo.nodeName == nodeName {
   return awsDiskInfo, true, nil
  }
 }
 return awsDiskInfo, false, nil
}
