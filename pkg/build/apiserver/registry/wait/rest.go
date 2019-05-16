package wait

import (
	"context"
	"fmt"
	goformat "fmt"
	buildv1 "github.com/openshift/api/build/v1"
	buildtypedclient "github.com/openshift/client-go/build/clientset/versioned/typed/build/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	watchtools "k8s.io/client-go/tools/watch"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

var (
	ErrUnknownBuildPhase = fmt.Errorf("unknown build phase")
	ErrBuildDeleted      = fmt.Errorf("build was deleted")
)

type ErrWatchError struct{ error }

func WaitForRunningBuild(buildClient buildtypedclient.BuildsGetter, buildNamespace, buildName string, timeout time.Duration) (*buildv1.Build, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fieldSelector := fields.OneTermEqualSelector("metadata.name", buildName)
	options := metav1.ListOptions{FieldSelector: fieldSelector.String(), ResourceVersion: "0"}
	done := make(chan interface{}, 1)
	var resultBuild *buildv1.Build
	var success bool
	var resultErr error
	deadline := time.Now().Add(timeout)
	go func() {
		defer close(done)
		defer utilruntime.HandleCrash()
		for time.Now().Before(deadline) {
			_, err := buildClient.Builds(buildNamespace).Get(buildName, metav1.GetOptions{})
			if err != nil {
				resultErr = err
				if errors.IsNotFound(err) {
					resultErr = ErrBuildDeleted
				}
				return
			}
			w, err := buildClient.Builds(buildNamespace).Watch(options)
			if err != nil {
				resultErr = err
				return
			}
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			_, err = watchtools.UntilWithoutRetry(ctx, w, func(event watch.Event) (bool, error) {
				if event.Type == watch.Error {
					return false, ErrWatchError{fmt.Errorf("watch event type error: %v", event)}
				}
				obj, ok := event.Object.(*buildv1.Build)
				if !ok {
					return false, fmt.Errorf("received unknown object while watching for builds: %T", event.Object)
				}
				if event.Type == watch.Deleted {
					return false, ErrBuildDeleted
				}
				switch obj.Status.Phase {
				case buildv1.BuildPhaseRunning, buildv1.BuildPhaseComplete, buildv1.BuildPhaseFailed, buildv1.BuildPhaseError, buildv1.BuildPhaseCancelled:
					resultBuild = obj
					return true, nil
				case buildv1.BuildPhaseNew, buildv1.BuildPhasePending:
				default:
					return false, ErrUnknownBuildPhase
				}
				return false, nil
			})
			if err != nil {
				if _, ok := err.(ErrWatchError); ok {
					continue
				}
				resultErr = err
				success = false
				resultBuild = nil
				return
			}
			success = true
			return
		}
	}()
	select {
	case <-time.After(timeout):
		return nil, false, wait.ErrWaitTimeout
	case <-done:
		return resultBuild, success, resultErr
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
