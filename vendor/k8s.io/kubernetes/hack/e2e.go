package main

import (
	"flag"
	"fmt"
	goformat "fmt"
	"go/build"
	"log"
	"os"
	goos "os"
	"os/exec"
	"os/signal"
	"path/filepath"
	godefaultruntime "runtime"
	"strings"
	"time"
	gotime "time"
)

type flags struct {
	get  bool
	old  time.Duration
	args []string
}

const (
	getDefault = true
	oldDefault = 24 * time.Hour
)

func parse(args []string) (flags, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fs := flag.NewFlagSet(args[0], flag.ContinueOnError)
	get := fs.Bool("get", getDefault, "go get -u kubetest if old or not installed")
	old := fs.Duration("old", oldDefault, "Consider kubetest old if it exceeds this")
	verbose := fs.Bool("v", true, "this flag is translated to kubetest's --verbose-commands")
	var a []string
	if err := fs.Parse(args[1:]); err == flag.ErrHelp {
		os.Stderr.WriteString("  -- kubetestArgs\n")
		os.Stderr.WriteString("        All flags after -- are passed to the kubetest program\n")
		return flags{}, err
	} else if err != nil {
		log.Print("NOTICE: go run hack/e2e.go is now a shim for test-infra/kubetest")
		log.Printf("  Usage: go run hack/e2e.go [--get=%v] [--old=%v] -- [KUBETEST_ARGS]", getDefault, oldDefault)
		log.Print("  The separator is required to use --get or --old flags")
		log.Print("  The -- flag separator also suppresses this message")
		a = args[len(args)-fs.NArg()-1:]
	} else {
		a = append(a, fmt.Sprintf("--verbose-commands=%t", *verbose))
		a = append(a, fs.Args()...)
	}
	return flags{*get, *old, a}, nil
}
func main() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	f, err := parse(os.Args)
	if err != nil {
		os.Exit(2)
	}
	t := newTester()
	k, err := t.getKubetest(f.get, f.old)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	log.Printf("Calling kubetest %v...", strings.Join(f.args, " "))
	if err = t.wait(k, f.args...); err != nil {
		log.Fatalf("err: %v", err)
	}
	log.Print("Done")
}
func wait(cmd string, args ...string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt)
	c := exec.Command(cmd, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Start(); err != nil {
		return err
	}
	go func() {
		sig := <-sigChannel
		if err := c.Process.Signal(sig); err != nil {
			log.Fatalf("could not send %s signal %s: %v", cmd, sig, err)
		}
	}()
	return c.Wait()
}

type tester struct {
	stat     func(string) (os.FileInfo, error)
	lookPath func(string) (string, error)
	goPath   string
	wait     func(string, ...string) error
}

func newTester() tester {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return tester{os.Stat, exec.LookPath, build.Default.GOPATH, wait}
}
func (t tester) lookKubetest() (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if t.goPath != "" {
		p := filepath.Join(t.goPath, "bin", "kubetest")
		_, err := t.stat(p)
		if err == nil {
			return p, nil
		}
	}
	p, err := t.lookPath("kubetest")
	return p, err
}
func (t tester) getKubetest(get bool, old time.Duration) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	p, err := t.lookKubetest()
	if err == nil && !get {
		return p, nil
	}
	if err == nil {
		if s, err := t.stat(p); err != nil {
			return p, err
		} else if time.Since(s.ModTime()) <= old {
			return p, nil
		} else if t.goPath == "" {
			log.Print("Skipping kubetest upgrade because $GOPATH is empty")
			return p, nil
		}
		log.Printf("The kubetest binary is older than %s.", old)
	}
	if t.goPath == "" {
		return "", fmt.Errorf("Cannot install kubetest until $GOPATH is set")
	}
	log.Print("Updating kubetest binary...")
	cmd := []string{"go", "get", "-u", "k8s.io/test-infra/kubetest"}
	if err = t.wait(cmd[0], cmd[1:]...); err != nil {
		return "", fmt.Errorf("%s: %v", strings.Join(cmd, " "), err)
	}
	if p, err = t.lookKubetest(); err != nil {
		return "", err
	} else if err = t.wait("touch", p); err != nil {
		return "", err
	} else {
		return p, nil
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
