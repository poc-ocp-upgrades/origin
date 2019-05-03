package main

import (
	"github.com/openshift/library-go/pkg/serviceability"
	_ "github.com/openshift/origin/pkg/api/install"
	"github.com/openshift/origin/pkg/api/legacy"
	"github.com/openshift/origin/pkg/cmd/openshift"
	"github.com/openshift/origin/pkg/version"
	"k8s.io/apiserver/pkg/util/logs"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	_ "k8s.io/kubernetes/pkg/apis/autoscaling/install"
	_ "k8s.io/kubernetes/pkg/apis/batch/install"
	_ "k8s.io/kubernetes/pkg/apis/core/install"
	_ "k8s.io/kubernetes/pkg/apis/extensions/install"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	logs.InitLogs()
	defer logs.FlushLogs()
	defer serviceability.BehaviorOnPanic(os.Getenv("OPENSHIFT_ON_PANIC"), version.Get())()
	defer serviceability.Profile(os.Getenv("OPENSHIFT_PROFILE")).Stop()
	legacy.InstallInternalLegacyAll(legacyscheme.Scheme)
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	basename := filepath.Base(os.Args[0])
	command := openshift.CommandFor(basename)
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
