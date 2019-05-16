package workflow

import (
	goformat "fmt"
	"github.com/spf13/pflag"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type Phase struct {
	Name           string
	Aliases        []string
	Short          string
	Long           string
	Example        string
	Hidden         bool
	Phases         []Phase
	RunAllSiblings bool
	Run            func(data RunData) error
	RunIf          func(data RunData) (bool, error)
	InheritFlags   []string
	LocalFlags     *pflag.FlagSet
}

func (t *Phase) AppendPhase(phase Phase) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	t.Phases = append(t.Phases, phase)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
