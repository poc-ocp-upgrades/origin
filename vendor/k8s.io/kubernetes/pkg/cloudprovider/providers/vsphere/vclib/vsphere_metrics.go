package vclib

import (
 "time"
 "github.com/prometheus/client_golang/prometheus"
)

const (
 APICreateVolume = "CreateVolume"
 APIDeleteVolume = "DeleteVolume"
 APIAttachVolume = "AttachVolume"
 APIDetachVolume = "DetachVolume"
)
const (
 OperationDeleteVolume                  = "DeleteVolumeOperation"
 OperationAttachVolume                  = "AttachVolumeOperation"
 OperationDetachVolume                  = "DetachVolumeOperation"
 OperationDiskIsAttached                = "DiskIsAttachedOperation"
 OperationDisksAreAttached              = "DisksAreAttachedOperation"
 OperationCreateVolume                  = "CreateVolumeOperation"
 OperationCreateVolumeWithPolicy        = "CreateVolumeWithPolicyOperation"
 OperationCreateVolumeWithRawVSANPolicy = "CreateVolumeWithRawVSANPolicyOperation"
)

var vsphereAPIMetric = prometheus.NewHistogramVec(prometheus.HistogramOpts{Name: "cloudprovider_vsphere_api_request_duration_seconds", Help: "Latency of vsphere api call"}, []string{"request"})
var vsphereAPIErrorMetric = prometheus.NewCounterVec(prometheus.CounterOpts{Name: "cloudprovider_vsphere_api_request_errors", Help: "vsphere Api errors"}, []string{"request"})
var vsphereOperationMetric = prometheus.NewHistogramVec(prometheus.HistogramOpts{Name: "cloudprovider_vsphere_operation_duration_seconds", Help: "Latency of vsphere operation call"}, []string{"operation"})
var vsphereOperationErrorMetric = prometheus.NewCounterVec(prometheus.CounterOpts{Name: "cloudprovider_vsphere_operation_errors", Help: "vsphere operation errors"}, []string{"operation"})

func RegisterMetrics() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 prometheus.MustRegister(vsphereAPIMetric)
 prometheus.MustRegister(vsphereAPIErrorMetric)
 prometheus.MustRegister(vsphereOperationMetric)
 prometheus.MustRegister(vsphereOperationErrorMetric)
}
func RecordvSphereMetric(actionName string, requestTime time.Time, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 switch actionName {
 case APICreateVolume, APIDeleteVolume, APIAttachVolume, APIDetachVolume:
  recordvSphereAPIMetric(actionName, requestTime, err)
 default:
  recordvSphereOperationMetric(actionName, requestTime, err)
 }
}
func recordvSphereAPIMetric(actionName string, requestTime time.Time, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err != nil {
  vsphereAPIErrorMetric.With(prometheus.Labels{"request": actionName}).Inc()
 } else {
  vsphereAPIMetric.With(prometheus.Labels{"request": actionName}).Observe(calculateTimeTaken(requestTime))
 }
}
func recordvSphereOperationMetric(actionName string, requestTime time.Time, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err != nil {
  vsphereOperationErrorMetric.With(prometheus.Labels{"operation": actionName}).Inc()
 } else {
  vsphereOperationMetric.With(prometheus.Labels{"operation": actionName}).Observe(calculateTimeTaken(requestTime))
 }
}
func RecordCreateVolumeMetric(volumeOptions *VolumeOptions, requestTime time.Time, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var actionName string
 if volumeOptions.StoragePolicyName != "" {
  actionName = OperationCreateVolumeWithPolicy
 } else if volumeOptions.VSANStorageProfileData != "" {
  actionName = OperationCreateVolumeWithRawVSANPolicy
 } else {
  actionName = OperationCreateVolume
 }
 RecordvSphereMetric(actionName, requestTime, err)
}
func calculateTimeTaken(requestBeginTime time.Time) (timeTaken float64) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !requestBeginTime.IsZero() {
  timeTaken = time.Since(requestBeginTime).Seconds()
 } else {
  timeTaken = 0
 }
 return timeTaken
}
