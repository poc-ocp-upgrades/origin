package dockerfile

import (
	"fmt"
	goformat "fmt"
	"github.com/fsouza/go-dockerclient"
	"github.com/moby/buildkit/frontend/dockerfile/command"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	goos "os"
	"regexp"
	godefaultruntime "runtime"
	"strconv"
	"strings"
	gotime "time"
)

var portRangeRegexp = regexp.MustCompile(`^(\d+)-(\d+)$`)
var argSplitRegexp = regexp.MustCompile(`^([a-zA-Z_]+[a-zA-Z0-9_]*)=(.*)$`)

func FindAll(node *parser.Node, cmd string) []int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if node == nil {
		return nil
	}
	var indices []int
	for i, child := range node.Children {
		if child != nil && child.Value == cmd {
			indices = append(indices, i)
		}
	}
	return indices
}
func InsertInstructions(node *parser.Node, pos int, instructions string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if node == nil {
		return fmt.Errorf("cannot insert instructions in a nil node")
	}
	if pos < 0 || pos > len(node.Children) {
		return fmt.Errorf("pos %d out of range [0, %d]", pos, len(node.Children)-1)
	}
	newChild, err := parser.Parse(strings.NewReader(instructions))
	if err != nil {
		return err
	}
	node.Children = append(node.Children[:pos], append(newChild.AST.Children, node.Children[pos:]...)...)
	return nil
}
func LastBaseImage(node *parser.Node) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	baseImages := baseImages(node)
	if len(baseImages) == 0 {
		return ""
	}
	return baseImages[len(baseImages)-1]
}
func baseImages(node *parser.Node) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var images []string
	for _, pos := range FindAll(node, command.From) {
		images = append(images, nextValues(node.Children[pos])...)
	}
	return images
}
func evalRange(port string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	m, match := match(portRangeRegexp, port)
	if !match {
		return port
	}
	_, err := strconv.ParseUint(m[1], 10, 16)
	if err != nil {
		return port
	}
	return m[1]
}
func evalPorts(exposedPorts []string, node *parser.Node, from, to int) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	shlex := NewShellLex('\\')
	shlex.ProcessWord("w", []string{})
	portsEnv := evalVars(node, from, to, exposedPorts, shlex)
	ports := make([]string, 0, len(portsEnv))
	for _, p := range portsEnv {
		dp := docker.Port(p)
		port := dp.Port()
		port = evalRange(port)
		if strings.Contains(p, `/`) {
			p = port + `/` + dp.Proto()
		} else {
			p = port
		}
		ports = append(ports, p)
	}
	return ports
}
func LastExposedPorts(node *parser.Node) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	exposedPorts, exposeIndices := exposedPorts(node)
	if len(exposedPorts) == 0 || len(exposeIndices) == 0 {
		return nil
	}
	lastExposed := exposedPorts[len(exposedPorts)-1]
	froms := FindAll(node, command.From)
	from := froms[len(froms)-1]
	to := exposeIndices[len(exposeIndices)-1]
	return evalPorts(lastExposed, node, from, to)
}
func exposedPorts(node *parser.Node) ([][]string, []int) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var allPorts [][]string
	var ports []string
	froms := FindAll(node, command.From)
	exposes := FindAll(node, command.Expose)
	for i, j := len(froms)-1, len(exposes)-1; i >= 0; i-- {
		for ; j >= 0 && exposes[j] > froms[i]; j-- {
			ports = append(nextValues(node.Children[exposes[j]]), ports...)
		}
		allPorts = append([][]string{ports}, allPorts...)
		ports = nil
	}
	return allPorts, exposes
}
func nextValues(node *parser.Node) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if node == nil {
		return nil
	}
	var values []string
	for next := node.Next; next != nil; next = next.Next {
		values = append(values, next.Value)
	}
	return values
}
func match(r *regexp.Regexp, str string) ([]string, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	m := r.FindStringSubmatch(str)
	return m, len(m) == r.NumSubexp()+1
}
func containsVars(ports []string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, p := range ports {
		if strings.Contains(p, `$`) {
			return true
		}
	}
	return false
}
func evalVars(n *parser.Node, from, to int, ports []string, shlex *ShellLex) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	envs := make([]string, 0)
	if !containsVars(ports) {
		return ports
	}
	evaledPorts := make([]string, 0)
	for i := from; i <= to; i++ {
		switch n.Children[i].Value {
		case command.Expose:
			args := nextValues(n.Children[i])
			for _, arg := range args {
				if processed, err := shlex.ProcessWord(arg, envs); err == nil {
					evaledPorts = append(evaledPorts, processed)
				} else {
					evaledPorts = append(evaledPorts, arg)
				}
			}
		case command.Arg:
			args := nextValues(n.Children[i])
			if len(args) == 1 {
				if _, match := match(argSplitRegexp, args[0]); match {
					if processed, err := shlex.ProcessWord(args[0], envs); err == nil {
						envs = append([]string{processed}, envs...)
					}
				}
			}
		case command.Env:
			args := nextValues(n.Children[i])
			currentEnvs := make([]string, 0)
			for j := 0; j < len(args)-1; j += 2 {
				if processed, err := shlex.ProcessWord(args[j+1], envs); err == nil {
					currentEnvs = append(currentEnvs, args[j]+"="+processed)
				}
			}
			envs = append(currentEnvs, envs...)
		}
	}
	return evaledPorts
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
