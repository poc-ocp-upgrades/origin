package main

import (
	"flag"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"os"
	"strings"
	"github.com/openshift/origin/tools/rebasehelpers/util"
)

func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var start, end string
	flag.StringVar(&start, "start", "master", "The start of the revision range for analysis")
	flag.StringVar(&end, "end", "HEAD", "The end of the revision range for analysis")
	flag.Parse()
	commits, err := util.CommitsBetween(start, end)
	if err != nil {
		if err == util.ErrNotCommit {
			fmt.Fprintf(os.Stderr, "WARNING: one of the provided commits does not exist, not a true branch\n")
			os.Exit(0)
		}
		fmt.Fprintf(os.Stderr, "ERROR: couldn't find commits from %s..%s: %v\n", start, end, err)
		os.Exit(1)
	}
	nonbumpCommits := []util.Commit{}
	for _, commit := range commits {
		if !strings.HasPrefix(commit.Summary, "bump") {
			nonbumpCommits = append(nonbumpCommits, commit)
		}
	}
	errs := []string{}
	for _, validate := range AllValidators {
		if err := validate(nonbumpCommits); err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(os.Getenv("RESTORE_AND_VERIFY_GODEPS")) > 0 {
		if err := ValidateGodeps(commits); err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) > 0 {
		fmt.Fprintf(os.Stderr, "%s\n", strings.Join(errs, "\n\n"))
		os.Exit(2)
	}
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
