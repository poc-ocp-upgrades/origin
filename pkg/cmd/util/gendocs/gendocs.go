package gendocs

import (
	"bytes"
	goformat "fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/cli-runtime/pkg/genericclioptions/printers"
	"os"
	goos "os"
	"path/filepath"
	godefaultruntime "runtime"
	"sort"
	gotime "time"
)

type Examples []*unstructured.Unstructured

func (x Examples) Len() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return len(x)
}
func (x Examples) Swap(i, j int) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	x[i], x[j] = x[j], x[i]
}
func (x Examples) Less(i, j int) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	xi, _ := x[i].Object["fullName"].(string)
	xj, _ := x[j].Object["fullName"].(string)
	return xi < xj
}
func GenDocs(cmd *cobra.Command, filename string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	out := new(bytes.Buffer)
	templateFile, err := filepath.Abs("hack/clibyexample/template")
	if err != nil {
		return err
	}
	template, err := ioutil.ReadFile(templateFile)
	if err != nil {
		return err
	}
	output := &unstructured.UnstructuredList{}
	output.SetGroupVersionKind(corev1.SchemeGroupVersion.WithKind("List"))
	examples := extractExamples(cmd)
	for i := range examples {
		output.Items = append(output.Items, *examples[i])
	}
	printer, err := printers.NewGoTemplatePrinter(template)
	if err != nil {
		return err
	}
	err = printer.PrintObj(output, out)
	if err != nil {
		return err
	}
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()
	_, err = outFile.Write(out.Bytes())
	if err != nil {
		return err
	}
	return nil
}
func extractExamples(cmd *cobra.Command) Examples {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	objs := Examples{}
	for _, c := range cmd.Commands() {
		if len(c.Deprecated) > 0 {
			continue
		}
		objs = append(objs, extractExamples(c)...)
	}
	if cmd.HasExample() {
		o := &unstructured.Unstructured{Object: make(map[string]interface{})}
		o.Object["name"] = cmd.Name()
		o.Object["fullName"] = cmd.CommandPath()
		o.Object["description"] = cmd.Short
		o.Object["examples"] = cmd.Example
		objs = append(objs, o)
	}
	sort.Sort(objs)
	return objs
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
