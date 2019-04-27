package buildlog

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"time"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	genericrest "k8s.io/apiserver/pkg/registry/generic/rest"
	"k8s.io/apiserver/pkg/registry/rest"
	kubetypedclient "k8s.io/client-go/kubernetes/typed/core/v1"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	"github.com/openshift/api/build"
	buildv1 "github.com/openshift/api/build/v1"
	buildtypedclient "github.com/openshift/client-go/build/clientset/versioned/typed/build/v1"
	apiserverrest "github.com/openshift/origin/pkg/apiserver/rest"
	buildapi "github.com/openshift/origin/pkg/build/apis/build"
	buildinternalhelpers "github.com/openshift/origin/pkg/build/apis/build/internal_helpers"
	"github.com/openshift/origin/pkg/build/apis/build/validation"
	buildwait "github.com/openshift/origin/pkg/build/apiserver/registry/wait"
	buildstrategy "github.com/openshift/origin/pkg/build/controller/strategy"
	buildutil "github.com/openshift/origin/pkg/build/util"
)

type REST struct {
	BuildClient	buildtypedclient.BuildsGetter
	PodClient	kubetypedclient.PodsGetter
	Timeout		time.Duration
	getSimpleLogsFn	func(podNamespace, podName string, logOpts *kapi.PodLogOptions) (runtime.Object, error)
}

const defaultTimeout time.Duration = 30 * time.Second

func NewREST(buildClient buildtypedclient.BuildsGetter, podClient kubetypedclient.PodsGetter) *REST {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := &REST{BuildClient: buildClient, PodClient: podClient, Timeout: defaultTimeout}
	r.getSimpleLogsFn = r.getSimpleLogs
	return r
}

var _ = rest.GetterWithOptions(&REST{})

func (r *REST) Get(ctx context.Context, name string, opts runtime.Object) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	buildLogOpts, ok := opts.(*buildapi.BuildLogOptions)
	if !ok {
		return nil, errors.NewBadRequest(fmt.Sprintf("did not get an expected options: %T", opts))
	}
	if errs := validation.ValidateBuildLogOptions(buildLogOpts); len(errs) > 0 {
		return nil, errors.NewInvalid(build.Kind("BuildLogOptions"), "", errs)
	}
	build, err := r.BuildClient.Builds(apirequest.NamespaceValue(ctx)).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	if buildLogOpts.Previous {
		version := versionForBuild(build)
		version--
		previousBuildName := buildutil.BuildNameForConfigVersion(buildutil.ConfigNameForBuild(build), version)
		previous, err := r.BuildClient.Builds(apirequest.NamespaceValue(ctx)).Get(previousBuildName, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
		build = previous
	}
	switch build.Status.Phase {
	case buildv1.BuildPhaseNew, buildv1.BuildPhasePending:
		if buildLogOpts.NoWait {
			klog.V(4).Infof("Build %s/%s is in %s state. No logs to retrieve yet.", build.Namespace, build.Name, build.Status.Phase)
			return &genericrest.LocationStreamer{}, nil
		}
		klog.V(4).Infof("Build %s/%s is in %s state, waiting for Build to start", build.Namespace, build.Name, build.Status.Phase)
		latest, ok, err := buildwait.WaitForRunningBuild(r.BuildClient, build.Namespace, build.Name, r.Timeout)
		if err != nil {
			return nil, errors.NewBadRequest(fmt.Sprintf("unable to wait for build %s to run: %v", build.Name, err))
		}
		switch latest.Status.Phase {
		case buildv1.BuildPhaseError:
			return nil, errors.NewBadRequest(fmt.Sprintf("build %s encountered an error: %s", build.Name, buildutil.NoBuildLogsMessage))
		case buildv1.BuildPhaseCancelled:
			return nil, errors.NewBadRequest(fmt.Sprintf("build %s was cancelled: %s", build.Name, buildutil.NoBuildLogsMessage))
		}
		if !ok {
			return nil, errors.NewTimeoutError(fmt.Sprintf("timed out waiting for build %s to start after %s", build.Name, r.Timeout), 1)
		}
	case buildv1.BuildPhaseCancelled:
		return nil, errors.NewBadRequest(fmt.Sprintf("build %s was cancelled. %s", build.Name, buildutil.NoBuildLogsMessage))
	case buildv1.BuildPhaseError:
		return nil, errors.NewBadRequest(fmt.Sprintf("build %s is in an error state. %s", build.Name, buildutil.NoBuildLogsMessage))
	}
	buildPodName := buildutil.GetBuildPodName(build)
	buildPod, err := r.PodClient.Pods(build.Namespace).Get(buildPodName, metav1.GetOptions{})
	if err != nil {
		return nil, errors.NewBadRequest(err.Error())
	}
	if len(buildPod.Spec.InitContainers) == 0 {
		logOpts := buildinternalhelpers.BuildToPodLogOptions(buildLogOpts)
		return r.getSimpleLogsFn(build.Namespace, buildPodName, logOpts)
	}
	reader, writer := io.Pipe()
	pipeStreamer := PipeStreamer{In: writer, Out: reader, Flush: buildLogOpts.Follow, ContentType: "text/plain"}
	go func() {
		defer pipeStreamer.In.Close()
		doneWithContainer := map[string]bool{}
		waitForInitContainers := true
		initFailed := false
		sleep := true
		for waitForInitContainers {
			select {
			case <-ctx.Done():
				klog.V(4).Infof("timed out while iterating on build init containers for build pod %s/%s", build.Namespace, buildPodName)
				return
			default:
			}
			klog.V(4).Infof("iterating through build init containers for build pod %s/%s", build.Namespace, buildPodName)
			waitForInitContainers = false
			buildPod, err = r.PodClient.Pods(build.Namespace).Get(buildPodName, metav1.GetOptions{})
			if err != nil {
				s := fmt.Sprintf("error retrieving build pod %s/%s : %v", build.Namespace, buildPodName, err.Error())
				pipeStreamer.In.Write([]byte(s))
				return
			}
			for _, status := range buildPod.Status.InitContainerStatuses {
				klog.V(4).Infof("processing build pod: %s/%s container: %s in state %#v", build.Namespace, buildPodName, status.Name, status.State)
				if status.State.Terminated != nil && status.State.Terminated.ExitCode != 0 {
					initFailed = true
					waitForInitContainers = false
					if doneWithContainer[status.Name] {
						break
					}
				}
				if doneWithContainer[status.Name] {
					continue
				}
				if status.State.Waiting != nil {
					waitForInitContainers = true
					continue
				}
				containerLogOpts := buildinternalhelpers.BuildToPodLogOptions(buildLogOpts)
				containerLogOpts.Container = status.Name
				if status.State.Terminated != nil {
					containerLogOpts.Follow = false
				}
				if err := r.pipeLogs(ctx, build.Namespace, buildPodName, containerLogOpts, pipeStreamer.In); err != nil {
					klog.Errorf("error: failed to stream logs for build pod: %s/%s container: %s, due to: %v", build.Namespace, buildPodName, status.Name, err)
					return
				}
				sleep = false
				doneWithContainer[status.Name] = true
				if initFailed {
					break
				}
			}
			if !buildLogOpts.Follow {
				break
			}
			if sleep {
				time.Sleep(time.Second)
			}
		}
		if !initFailed {
			err := wait.PollImmediate(time.Second, 10*time.Minute, func() (bool, error) {
				buildPod, err = r.PodClient.Pods(build.Namespace).Get(buildPodName, metav1.GetOptions{})
				if err != nil {
					s := fmt.Sprintf("error while getting build logs, could not retrieve build pod %s/%s : %v", build.Namespace, buildPodName, err.Error())
					pipeStreamer.In.Write([]byte(s))
					return false, err
				}
				if buildPod.Status.Phase != corev1.PodPending {
					return true, nil
				}
				return false, nil
			})
			if err != nil {
				klog.Errorf("error: failed to stream logs for build pod: %s/%s due to: %v", build.Namespace, buildPodName, err)
				return
			}
			containerLogOpts := buildinternalhelpers.BuildToPodLogOptions(buildLogOpts)
			containerLogOpts.Container = selectBuilderContainer(buildPod.Spec.Containers)
			if containerLogOpts.Container == "" {
				klog.Errorf("error: failed to select a container in build pod: %s/%s", build.Namespace, buildPodName)
			}
			if buildPod.Status.Phase == corev1.PodFailed || buildPod.Status.Phase == corev1.PodSucceeded {
				containerLogOpts.Follow = false
			}
			if err := r.pipeLogs(ctx, build.Namespace, buildPodName, containerLogOpts, pipeStreamer.In); err != nil {
				klog.Errorf("error: failed to stream logs for build pod: %s/%s due to: %v", build.Namespace, buildPodName, err)
				return
			}
		}
	}()
	return &pipeStreamer, nil
}
func (r *REST) NewGetOptions() (runtime.Object, bool, string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &buildapi.BuildLogOptions{}, false, ""
}
func (r *REST) New() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &buildapi.BuildLog{}
}
func (r *REST) pipeLogs(ctx context.Context, namespace, buildPodName string, containerLogOpts *kapi.PodLogOptions, writer io.Writer) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(4).Infof("pulling build pod logs for %s/%s, container %s", namespace, buildPodName, containerLogOpts.Container)
	logRequest := r.PodClient.Pods(namespace).GetLogs(buildPodName, podLogOptionsToV1(containerLogOpts))
	readerCloser, err := logRequest.Stream()
	if err != nil {
		klog.Errorf("error: could not write build log for pod %q to stream due to: %v", buildPodName, err)
		return err
	}
	klog.V(4).Infof("retrieved logs for build pod: %s/%s container: %s", namespace, buildPodName, containerLogOpts.Container)
	_, err = io.Copy(writer, readerCloser)
	return err
}
func podLogOptionsToV1(options *kapi.PodLogOptions) *corev1.PodLogOptions {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	newOptions := &corev1.PodLogOptions{}
	if err := legacyscheme.Scheme.Convert(options, newOptions, nil); err != nil {
		panic(err)
	}
	return newOptions
}
func selectBuilderContainer(containers []corev1.Container) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, c := range containers {
		for _, bcName := range buildstrategy.BuildContainerNames {
			if c.Name == bcName {
				return bcName
			}
		}
	}
	return ""
}
func (r *REST) getSimpleLogs(podNamespace, podName string, logOpts *kapi.PodLogOptions) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	logRequest := r.PodClient.Pods(podNamespace).GetLogs(podName, podLogOptionsToV1(logOpts))
	readerCloser, err := logRequest.Stream()
	if err != nil {
		return nil, err
	}
	return &apiserverrest.PassThroughStreamer{In: readerCloser, Flush: logOpts.Follow, ContentType: "text/plain"}, nil
}
func versionForBuild(build *buildv1.Build) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if build == nil {
		return 0
	}
	versionString := build.Annotations[buildapi.BuildNumberAnnotation]
	version, err := strconv.Atoi(versionString)
	if err != nil {
		return 0
	}
	return version
}
