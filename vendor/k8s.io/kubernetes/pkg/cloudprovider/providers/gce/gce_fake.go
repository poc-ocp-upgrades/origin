package gce

import (
 "fmt"
 "net/http"
 compute "google.golang.org/api/compute/v1"
 "k8s.io/client-go/tools/cache"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud"
)

type TestClusterValues struct {
 ProjectID         string
 Region            string
 ZoneName          string
 SecondaryZoneName string
 ClusterID         string
 ClusterName       string
}

func DefaultTestClusterValues() TestClusterValues {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return TestClusterValues{ProjectID: "test-project", Region: "us-central1", ZoneName: "us-central1-b", SecondaryZoneName: "us-central1-c", ClusterID: "test-cluster-id", ClusterName: "Test Cluster Name"}
}

type fakeRoundTripper struct{}

func (*fakeRoundTripper) RoundTrip(*http.Request) (*http.Response, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil, fmt.Errorf("err: test used fake http client")
}
func fakeClusterID(clusterID string) ClusterID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return ClusterID{clusterID: &clusterID, store: cache.NewStore(func(obj interface{}) (string, error) {
  return "", nil
 })}
}
func NewFakeGCECloud(vals TestClusterValues) *Cloud {
 _logClusterCodePath()
 defer _logClusterCodePath()
 client := &http.Client{Transport: &fakeRoundTripper{}}
 service, _ := compute.New(client)
 gce := &Cloud{region: vals.Region, service: service, managedZones: []string{vals.ZoneName}, projectID: vals.ProjectID, networkProjectID: vals.ProjectID, ClusterID: fakeClusterID(vals.ClusterID)}
 c := cloud.NewMockGCE(&gceProjectRouter{gce})
 gce.c = c
 return gce
}
