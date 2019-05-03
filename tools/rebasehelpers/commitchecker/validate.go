package main

import (
	"bytes"
	"fmt"
	"github.com/openshift/origin/tools/rebasehelpers/util"
	"os/exec"
	"regexp"
	"strings"
	"text/template"
)

var CommitSummaryErrorTemplate = `
The following UPSTREAM commits have invalid summaries:

{{ range .Commits }}  [{{ .Sha }}] {{ .Summary }}
{{ end }}
UPSTREAM commit summaries should look like:

  UPSTREAM: <PR number|carry|drop>: description

UPSTREAM commits which revert previous UPSTREAM commits should look like:

  UPSTREAM: revert: <normal upstream format>

UPSTREAM commits are validated against the following regular expression:

  {{ .Pattern }}

Examples of valid summaries:

  UPSTREAM: 12345: A kube fix
  UPSTREAM: <carry>: A carried kube change
  UPSTREAM: <drop>: A dropped kube change
  UPSTREAM: revert: 12345: A kube revert

`
var AllValidators = []func([]util.Commit) error{ValidateUpstreamCommitSummaries, ValidateUpstreamCommitsWithoutGodepsChanges, ValidateUpstreamCommitModifiesOnlyGodeps, ValidateUpstreamCommitModifiesOnlyKubernetes}

func ValidateUpstreamCommitsWithoutGodepsChanges(commits []util.Commit) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	problemCommits := []util.Commit{}
	for _, commit := range commits {
		if commit.HasVendoredCodeChanges() && !commit.DeclaresUpstreamChange() {
			problemCommits = append(problemCommits, commit)
		}
	}
	if len(problemCommits) > 0 {
		label := "The following commits contain vendor changes but aren't declared as UPSTREAM or bump(*) commits"
		msg := renderGodepFilesError(label, problemCommits, RenderOnlyGodepsFiles)
		return fmt.Errorf(msg)
	}
	return nil
}
func ValidateUpstreamCommitModifiesSingleGodepsRepo(commits []util.Commit) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	problemCommits := []util.Commit{}
	for _, commit := range commits {
		godepsChanges, err := commit.GodepsReposChanged()
		if err != nil {
			return err
		}
		if len(godepsChanges) > 1 {
			problemCommits = append(problemCommits, commit)
		}
	}
	if len(problemCommits) > 0 {
		label := "The following UPSTREAM commits modify more than one repo in their changelist"
		msg := renderGodepFilesError(label, problemCommits, RenderOnlyGodepsFiles)
		return fmt.Errorf(msg)
	}
	return nil
}
func ValidateUpstreamCommitSummaries(commits []util.Commit) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	problemCommits := []util.Commit{}
	for _, commit := range commits {
		if commit.DeclaresUpstreamChange() && !commit.MatchesUpstreamSummaryPattern() {
			problemCommits = append(problemCommits, commit)
		}
	}
	if len(problemCommits) > 0 {
		tmpl, _ := template.New("problems").Parse(CommitSummaryErrorTemplate)
		data := struct {
			Pattern *regexp.Regexp
			Commits []util.Commit
		}{Pattern: util.UpstreamSummaryPattern, Commits: problemCommits}
		buffer := &bytes.Buffer{}
		tmpl.Execute(buffer, data)
		return fmt.Errorf(buffer.String())
	}
	return nil
}
func ValidateUpstreamCommitModifiesOnlyGodeps(commits []util.Commit) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	problemCommits := []util.Commit{}
	for _, commit := range commits {
		if commit.HasVendoredCodeChanges() && commit.HasNonVendoredCodeChanges() {
			problemCommits = append(problemCommits, commit)
		}
	}
	if len(problemCommits) > 0 {
		label := "The following UPSTREAM commits modify files outside vendor"
		msg := renderGodepFilesError(label, problemCommits, RenderAllFiles)
		return fmt.Errorf(msg)
	}
	return nil
}
func ValidateUpstreamCommitModifiesOnlyKubernetes(commits []util.Commit) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	problemCommits := []util.Commit{}
	for _, commit := range commits {
		if commit.DeclaresUpstreamChange() {
			reposChanged, err := commit.GodepsReposChanged()
			if err != nil {
				return err
			}
			for _, changedRepo := range reposChanged {
				if !strings.Contains(changedRepo, "k8s.io/kubernetes") {
					problemCommits = append(problemCommits, commit)
				}
			}
		}
	}
	if len(problemCommits) > 0 {
		label := "The following UPSTREAM commits modify vendored repos other than k8s.io/kubernetes"
		msg := renderGodepFilesError(label, problemCommits, RenderAllFiles)
		return fmt.Errorf(msg)
	}
	return nil
}
func ValidateUpstreamCommitModifiesOnlyDeclaredGodepRepo(commits []util.Commit) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	problemCommits := []util.Commit{}
	for _, commit := range commits {
		if commit.DeclaresUpstreamChange() {
			declaredRepo, err := commit.DeclaredUpstreamRepo()
			if err != nil {
				return err
			}
			reposChanged, err := commit.GodepsReposChanged()
			if err != nil {
				return err
			}
			for _, changedRepo := range reposChanged {
				if !strings.Contains(changedRepo, declaredRepo) {
					problemCommits = append(problemCommits, commit)
				}
			}
		}
	}
	if len(problemCommits) > 0 {
		label := "The following UPSTREAM commits modify Godeps repos other than the repo the commit declares"
		msg := renderGodepFilesError(label, problemCommits, RenderAllFiles)
		return fmt.Errorf(msg)
	}
	return nil
}
func ValidateGodeps(commits []util.Commit) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	runGodepsRestore := false
	for _, commit := range commits {
		if commit.HasVendoredCodeChanges() || commit.HasGodepsChanges() {
			runGodepsRestore = true
			break
		}
	}
	if runGodepsRestore {
		fmt.Println("Running godep-restore")
		cmd := exec.Command("hack/godep-restore.sh")
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("Error running hack/godep-restore.sh: %v\n%s\n%s", err, stderr.String(), stdout.String())
		}
	}
	return nil
}

type CommitFilesRenderOption int

const (
	RenderNoFiles CommitFilesRenderOption = iota
	RenderOnlyGodepsFiles
	RenderOnlyNonGodepsFiles
	RenderAllFiles
)

func renderGodepFilesError(label string, commits []util.Commit, opt CommitFilesRenderOption) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	msg := fmt.Sprintf("%s:\n\n", label)
	for _, commit := range commits {
		msg += fmt.Sprintf("[%s] %s\n", commit.Sha, commit.Summary)
		if opt == RenderNoFiles {
			continue
		}
		for _, file := range commit.Files {
			if opt == RenderAllFiles || (opt == RenderOnlyGodepsFiles && file.HasVendoredCodeChanges()) || (opt == RenderOnlyNonGodepsFiles && !file.HasVendoredCodeChanges()) {
				msg += fmt.Sprintf("  - %s\n", file)
			}
		}
	}
	return msg
}
