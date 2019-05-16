package main

import (
	"github.com/spf13/cobra"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func MarkdownPostProcessing(cmd *cobra.Command, dir string, processor func(string) string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, c := range cmd.Commands() {
		if !c.IsAvailableCommand() || c.IsAdditionalHelpTopicCommand() {
			continue
		}
		if err := MarkdownPostProcessing(c, dir, processor); err != nil {
			return err
		}
	}
	basename := strings.Replace(cmd.CommandPath(), " ", "_", -1) + ".md"
	filename := filepath.Join(dir, basename)
	markdownBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	processedMarkDown := processor(string(markdownBytes))
	return ioutil.WriteFile(filename, []byte(processedMarkDown), 0644)
}
func cleanupForInclude(md string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	lines := strings.Split(md, "\n")
	cleanMd := ""
	for i, line := range lines {
		if i == 0 {
			continue
		}
		if line == "### SEE ALSO" {
			break
		}
		cleanMd += line
		if i < len(lines)-1 {
			cleanMd += "\n"
		}
	}
	return cleanMd
}
