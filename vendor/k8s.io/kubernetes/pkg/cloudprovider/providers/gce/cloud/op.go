package cloud

import (
 "context"
 "fmt"
 "k8s.io/klog"
 alpha "google.golang.org/api/compute/v0.alpha"
 beta "google.golang.org/api/compute/v0.beta"
 ga "google.golang.org/api/compute/v1"
 "google.golang.org/api/googleapi"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/meta"
)

const (
 operationStatusDone = "DONE"
)

type operation interface {
 isDone(ctx context.Context) (bool, error)
 error() error
 rateLimitKey() *RateLimitKey
}
type gaOperation struct {
 s         *Service
 projectID string
 key       *meta.Key
 err       error
}

func (o *gaOperation) String() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("gaOperation{%q, %v}", o.projectID, o.key)
}
func (o *gaOperation) isDone(ctx context.Context) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var (
  op  *ga.Operation
  err error
 )
 switch o.key.Type() {
 case meta.Regional:
  op, err = o.s.GA.RegionOperations.Get(o.projectID, o.key.Region, o.key.Name).Context(ctx).Do()
  klog.V(5).Infof("GA.RegionOperations.Get(%v, %v, %v) = %+v, %v; ctx = %v", o.projectID, o.key.Region, o.key.Name, op, err, ctx)
 case meta.Zonal:
  op, err = o.s.GA.ZoneOperations.Get(o.projectID, o.key.Zone, o.key.Name).Context(ctx).Do()
  klog.V(5).Infof("GA.ZoneOperations.Get(%v, %v, %v) = %+v, %v; ctx = %v", o.projectID, o.key.Zone, o.key.Name, op, err, ctx)
 case meta.Global:
  op, err = o.s.GA.GlobalOperations.Get(o.projectID, o.key.Name).Context(ctx).Do()
  klog.V(5).Infof("GA.GlobalOperations.Get(%v, %v) = %+v, %v; ctx = %v", o.projectID, o.key.Name, op, err, ctx)
 default:
  return false, fmt.Errorf("invalid key type: %#v", o.key)
 }
 if err != nil {
  return false, err
 }
 if op == nil || op.Status != operationStatusDone {
  return false, nil
 }
 if op.Error != nil && len(op.Error.Errors) > 0 && op.Error.Errors[0] != nil {
  e := op.Error.Errors[0]
  o.err = &googleapi.Error{Code: int(op.HttpErrorStatusCode), Message: fmt.Sprintf("%v - %v", e.Code, e.Message)}
 }
 return true, nil
}
func (o *gaOperation) rateLimitKey() *RateLimitKey {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &RateLimitKey{ProjectID: o.projectID, Operation: "Get", Service: "Operations", Version: meta.VersionGA}
}
func (o *gaOperation) error() error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return o.err
}

type alphaOperation struct {
 s         *Service
 projectID string
 key       *meta.Key
 err       error
}

func (o *alphaOperation) String() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("alphaOperation{%q, %v}", o.projectID, o.key)
}
func (o *alphaOperation) isDone(ctx context.Context) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var (
  op  *alpha.Operation
  err error
 )
 switch o.key.Type() {
 case meta.Regional:
  op, err = o.s.Alpha.RegionOperations.Get(o.projectID, o.key.Region, o.key.Name).Context(ctx).Do()
  klog.V(5).Infof("Alpha.RegionOperations.Get(%v, %v, %v) = %+v, %v; ctx = %v", o.projectID, o.key.Region, o.key.Name, op, err, ctx)
 case meta.Zonal:
  op, err = o.s.Alpha.ZoneOperations.Get(o.projectID, o.key.Zone, o.key.Name).Context(ctx).Do()
  klog.V(5).Infof("Alpha.ZoneOperations.Get(%v, %v, %v) = %+v, %v; ctx = %v", o.projectID, o.key.Zone, o.key.Name, op, err, ctx)
 case meta.Global:
  op, err = o.s.Alpha.GlobalOperations.Get(o.projectID, o.key.Name).Context(ctx).Do()
  klog.V(5).Infof("Alpha.GlobalOperations.Get(%v, %v) = %+v, %v; ctx = %v", o.projectID, o.key.Name, op, err, ctx)
 default:
  return false, fmt.Errorf("invalid key type: %#v", o.key)
 }
 if err != nil {
  return false, err
 }
 if op == nil || op.Status != operationStatusDone {
  return false, nil
 }
 if op.Error != nil && len(op.Error.Errors) > 0 && op.Error.Errors[0] != nil {
  e := op.Error.Errors[0]
  o.err = &googleapi.Error{Code: int(op.HttpErrorStatusCode), Message: fmt.Sprintf("%v - %v", e.Code, e.Message)}
 }
 return true, nil
}
func (o *alphaOperation) rateLimitKey() *RateLimitKey {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &RateLimitKey{ProjectID: o.projectID, Operation: "Get", Service: "Operations", Version: meta.VersionAlpha}
}
func (o *alphaOperation) error() error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return o.err
}

type betaOperation struct {
 s         *Service
 projectID string
 key       *meta.Key
 err       error
}

func (o *betaOperation) String() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("betaOperation{%q, %v}", o.projectID, o.key)
}
func (o *betaOperation) isDone(ctx context.Context) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var (
  op  *beta.Operation
  err error
 )
 switch o.key.Type() {
 case meta.Regional:
  op, err = o.s.Beta.RegionOperations.Get(o.projectID, o.key.Region, o.key.Name).Context(ctx).Do()
  klog.V(5).Infof("Beta.RegionOperations.Get(%v, %v, %v) = %+v, %v; ctx = %v", o.projectID, o.key.Region, o.key.Name, op, err, ctx)
 case meta.Zonal:
  op, err = o.s.Beta.ZoneOperations.Get(o.projectID, o.key.Zone, o.key.Name).Context(ctx).Do()
  klog.V(5).Infof("Beta.ZoneOperations.Get(%v, %v, %v) = %+v, %v; ctx = %v", o.projectID, o.key.Zone, o.key.Name, op, err, ctx)
 case meta.Global:
  op, err = o.s.Beta.GlobalOperations.Get(o.projectID, o.key.Name).Context(ctx).Do()
  klog.V(5).Infof("Beta.GlobalOperations.Get(%v, %v) = %+v, %v; ctx = %v", o.projectID, o.key.Name, op, err, ctx)
 default:
  return false, fmt.Errorf("invalid key type: %#v", o.key)
 }
 if err != nil {
  return false, err
 }
 if op == nil || op.Status != operationStatusDone {
  return false, nil
 }
 if op.Error != nil && len(op.Error.Errors) > 0 && op.Error.Errors[0] != nil {
  e := op.Error.Errors[0]
  o.err = &googleapi.Error{Code: int(op.HttpErrorStatusCode), Message: fmt.Sprintf("%v - %v", e.Code, e.Message)}
 }
 return true, nil
}
func (o *betaOperation) rateLimitKey() *RateLimitKey {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &RateLimitKey{ProjectID: o.projectID, Operation: "Get", Service: "Operations", Version: meta.VersionBeta}
}
func (o *betaOperation) error() error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return o.err
}
