package term

import (
	"bufio"
	"fmt"
	goformat "fmt"
	"github.com/docker/docker/pkg/term"
	"io"
	"k8s.io/klog"
	"os"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

func PromptForString(r io.Reader, w io.Writer, format string, a ...interface{}) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w == nil {
		w = os.Stdout
	}
	fmt.Fprintf(w, format, a...)
	return readInput(r)
}
func PromptForPasswordString(r io.Reader, w io.Writer, format string, a ...interface{}) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w == nil {
		w = os.Stdout
	}
	if file, ok := r.(*os.File); ok {
		inFd := file.Fd()
		if term.IsTerminal(inFd) {
			oldState, err := term.SaveState(inFd)
			if err != nil {
				klog.V(3).Infof("Unable to save terminal state")
				return PromptForString(r, w, format, a...)
			}
			fmt.Fprintf(w, format, a...)
			term.DisableEcho(inFd, oldState)
			input := readInput(r)
			defer term.RestoreTerminal(inFd, oldState)
			fmt.Fprintf(w, "\n")
			return input
		}
		klog.V(3).Infof("Stdin is not a terminal")
		return PromptForString(r, w, format, a...)
	}
	return PromptForString(r, w, format, a...)
}
func PromptForBool(r io.Reader, w io.Writer, format string, a ...interface{}) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w == nil {
		w = os.Stdout
	}
	str := PromptForString(r, w, format, a...)
	switch strings.ToLower(str) {
	case "1", "t", "true", "y", "yes":
		return true
	case "0", "f", "false", "n", "no":
		return false
	}
	fmt.Println("Please enter 'yes' or 'no'.")
	return PromptForBool(r, w, format, a...)
}
func PromptForStringWithDefault(r io.Reader, w io.Writer, def string, format string, a ...interface{}) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w == nil {
		w = os.Stdout
	}
	s := PromptForString(r, w, format, a...)
	if len(s) == 0 {
		return def
	}
	return s
}
func readInput(r io.Reader) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if _, isTerminal := term.GetFdInfo(r); isTerminal {
		return readInputFromTerminal(r)
	}
	return readInputFromReader(r)
}
func readInputFromTerminal(r io.Reader) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	reader := bufio.NewReader(r)
	result, _ := reader.ReadString('\n')
	return strings.TrimRight(result, "\r\n")
}
func readInputFromReader(r io.Reader) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var result string
	fmt.Fscan(r, &result)
	return result
}
func IsTerminalReader(r io.Reader) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	file, ok := r.(*os.File)
	return ok && term.IsTerminal(file.Fd())
}
func IsTerminalWriter(w io.Writer) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	file, ok := w.(*os.File)
	return ok && term.IsTerminal(file.Fd())
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
