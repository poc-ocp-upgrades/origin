package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"k8s.io/klog"
	"math"
	"sync"
	"time"
)

const (
	decayIntervalSeconds = 20
	decayFraction        = 0.8
	maxDelay             = 60 * time.Second
)

type CrossRequestRetryDelay struct{ backoff Backoff }

func NewCrossRequestRetryDelay() *CrossRequestRetryDelay {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c := &CrossRequestRetryDelay{}
	c.backoff.init(decayIntervalSeconds, decayFraction, maxDelay)
	return c
}
func (c *CrossRequestRetryDelay) BeforeSign(r *request.Request) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	now := time.Now()
	delay := c.backoff.ComputeDelayForRequest(now)
	if delay > 0 {
		klog.Warningf("Inserting delay before AWS request (%s) to avoid RequestLimitExceeded: %s", describeRequest(r), delay.String())
		if sleepFn := r.Config.SleepDelay; sleepFn != nil {
			sleepFn(delay)
		} else if err := aws.SleepWithContext(r.Context(), delay); err != nil {
			r.Error = awserr.New(request.CanceledErrorCode, "request context canceled", err)
			r.Retryable = aws.Bool(false)
			return
		}
		r.Time = now
	}
}
func operationName(r *request.Request) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	name := "?"
	if r.Operation != nil {
		name = r.Operation.Name
	}
	return name
}
func describeRequest(r *request.Request) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	service := r.ClientInfo.ServiceName
	return service + "::" + operationName(r)
}
func (c *CrossRequestRetryDelay) AfterRetry(r *request.Request) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if r.Error == nil {
		return
	}
	awsError, ok := r.Error.(awserr.Error)
	if !ok {
		return
	}
	if awsError.Code() == "RequestLimitExceeded" {
		c.backoff.ReportError()
		recordAWSThrottlesMetric(operationName(r))
		klog.Warningf("Got RequestLimitExceeded error on AWS request (%s)", describeRequest(r))
	}
}

type Backoff struct {
	decayIntervalSeconds    int64
	decayFraction           float64
	maxDelay                time.Duration
	mutex                   sync.Mutex
	countErrorsRequestLimit float32
	countRequests           float32
	lastDecay               int64
}

func (b *Backoff) init(decayIntervalSeconds int, decayFraction float64, maxDelay time.Duration) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	b.lastDecay = time.Now().Unix()
	b.countRequests = 4
	b.decayIntervalSeconds = int64(decayIntervalSeconds)
	b.decayFraction = decayFraction
	b.maxDelay = maxDelay
}
func (b *Backoff) ComputeDelayForRequest(now time.Time) time.Duration {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	b.mutex.Lock()
	defer b.mutex.Unlock()
	timeDeltaSeconds := now.Unix() - b.lastDecay
	if timeDeltaSeconds > b.decayIntervalSeconds {
		intervals := float64(timeDeltaSeconds) / float64(b.decayIntervalSeconds)
		decay := float32(math.Pow(b.decayFraction, intervals))
		b.countErrorsRequestLimit *= decay
		b.countRequests *= decay
		b.lastDecay = now.Unix()
	}
	b.countRequests += 1.0
	errorFraction := float32(0.0)
	if b.countRequests > 0.5 {
		errorFraction = b.countErrorsRequestLimit / b.countRequests
	}
	if errorFraction < 0.1 {
		return time.Duration(0)
	}
	delay := time.Nanosecond * time.Duration(float32(b.maxDelay.Nanoseconds())*errorFraction)
	return time.Second * time.Duration(int(delay.Seconds()))
}
func (b *Backoff) ReportError() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.countErrorsRequestLimit += 1.0
}
