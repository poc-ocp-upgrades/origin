package cloud

import (
	"context"
	"fmt"
	alpha "google.golang.org/api/compute/v0.alpha"
	beta "google.golang.org/api/compute/v0.beta"
	ga "google.golang.org/api/compute/v1"
	"k8s.io/klog"
)

type Service struct {
	GA            *ga.Service
	Alpha         *alpha.Service
	Beta          *beta.Service
	ProjectRouter ProjectRouter
	RateLimiter   RateLimiter
}

func (s *Service) wrapOperation(anyOp interface{}) (operation, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch o := anyOp.(type) {
	case *ga.Operation:
		r, err := ParseResourceURL(o.SelfLink)
		if err != nil {
			return nil, err
		}
		return &gaOperation{s: s, projectID: r.ProjectID, key: r.Key}, nil
	case *alpha.Operation:
		r, err := ParseResourceURL(o.SelfLink)
		if err != nil {
			return nil, err
		}
		return &alphaOperation{s: s, projectID: r.ProjectID, key: r.Key}, nil
	case *beta.Operation:
		r, err := ParseResourceURL(o.SelfLink)
		if err != nil {
			return nil, err
		}
		return &betaOperation{s: s, projectID: r.ProjectID, key: r.Key}, nil
	default:
		return nil, fmt.Errorf("invalid type %T", anyOp)
	}
}
func (s *Service) WaitForCompletion(ctx context.Context, genericOp interface{}) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	op, err := s.wrapOperation(genericOp)
	if err != nil {
		klog.Errorf("wrapOperation(%+v) error: %v", genericOp, err)
		return err
	}
	return s.pollOperation(ctx, op)
}
func (s *Service) pollOperation(ctx context.Context, op operation) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var pollCount int
	for {
		select {
		case <-ctx.Done():
			klog.V(5).Infof("op.pollOperation(%v, %v) not completed, poll count = %d, ctx.Err = %v", ctx, op, pollCount, ctx.Err())
			return ctx.Err()
		default:
		}
		pollCount++
		klog.V(5).Infof("op.isDone(%v) waiting; op = %v, poll count = %d", ctx, op, pollCount)
		s.RateLimiter.Accept(ctx, op.rateLimitKey())
		done, err := op.isDone(ctx)
		if err != nil {
			klog.V(5).Infof("op.isDone(%v) error; op = %v, poll count = %d, err = %v, retrying", ctx, op, pollCount, err)
		}
		if done {
			break
		}
	}
	klog.V(5).Infof("op.isDone(%v) complete; op = %v, poll count = %d, op.err = %v", ctx, op, pollCount, op.error())
	return op.error()
}
