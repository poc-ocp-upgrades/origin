package docker10

import (
	goformat "fmt"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func Convert_DockerV1CompatibilityImage_to_DockerImageConfig(in *DockerV1CompatibilityImage, out *DockerImageConfig) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	*out = DockerImageConfig{ID: in.ID, Parent: in.Parent, Comment: in.Comment, Created: in.Created, Container: in.Container, DockerVersion: in.DockerVersion, Author: in.Author, Architecture: in.Architecture, Size: in.Size, OS: "linux", ContainerConfig: in.ContainerConfig}
	if in.Config != nil {
		out.Config = &DockerConfig{}
		*out.Config = *in.Config
	}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
