package metrics

import (
 "fmt"
)

func GetResourceUtilizationRatio(metrics PodMetricsInfo, requests map[string]int64, targetUtilization int32) (utilizationRatio float64, currentUtilization int32, rawAverageValue int64, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 metricsTotal := int64(0)
 requestsTotal := int64(0)
 numEntries := 0
 for podName, metric := range metrics {
  request, hasRequest := requests[podName]
  if !hasRequest {
   continue
  }
  metricsTotal += metric.Value
  requestsTotal += request
  numEntries++
 }
 if requestsTotal == 0 {
  return 0, 0, 0, fmt.Errorf("no metrics returned matched known pods")
 }
 currentUtilization = int32((metricsTotal * 100) / requestsTotal)
 return float64(currentUtilization) / float64(targetUtilization), currentUtilization, metricsTotal / int64(numEntries), nil
}
func GetMetricUtilizationRatio(metrics PodMetricsInfo, targetUtilization int64) (utilizationRatio float64, currentUtilization int64) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 metricsTotal := int64(0)
 for _, metric := range metrics {
  metricsTotal += metric.Value
 }
 currentUtilization = metricsTotal / int64(len(metrics))
 return float64(currentUtilization) / float64(targetUtilization), currentUtilization
}
