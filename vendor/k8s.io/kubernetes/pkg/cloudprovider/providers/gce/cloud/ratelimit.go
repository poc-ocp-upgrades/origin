package cloud

import (
	"context"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/meta"
	"time"
)

type RateLimitKey struct {
	ProjectID string
	Operation string
	Version   meta.Version
	Service   string
}
type RateLimiter interface {
	Accept(ctx context.Context, key *RateLimitKey) error
}
type acceptor interface{ Accept() }
type AcceptRateLimiter struct{ Acceptor acceptor }

func (rl *AcceptRateLimiter) Accept(ctx context.Context, key *RateLimitKey) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ch := make(chan struct{})
	go func() {
		rl.Acceptor.Accept()
		close(ch)
	}()
	select {
	case <-ch:
		break
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

type NopRateLimiter struct{}

func (*NopRateLimiter) Accept(ctx context.Context, key *RateLimitKey) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}

type MinimumRateLimiter struct {
	RateLimiter RateLimiter
	Minimum     time.Duration
}

func (m *MinimumRateLimiter) Accept(ctx context.Context, key *RateLimitKey) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	select {
	case <-time.After(m.Minimum):
		return m.RateLimiter.Accept(ctx, key)
	case <-ctx.Done():
		return ctx.Err()
	}
}
