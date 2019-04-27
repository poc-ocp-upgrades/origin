package image

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/openshift/api/image/docker10"
	public "github.com/openshift/origin/pkg/image/apis/image/docker10"
)

type DockerImage = docker10.DockerImage
type DockerConfig = docker10.DockerConfig

func Convert_public_to_api_DockerImage(in *public.DockerImage, out *docker10.DockerImage) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = docker10.DockerImage{ID: in.ID, Parent: in.Parent, Comment: in.Comment, Created: metav1.Time{Time: in.Created}, Container: in.Container, DockerVersion: in.DockerVersion, Author: in.Author, Architecture: in.Architecture, Size: in.Size}
	if err := Convert_public_to_api_DockerConfig(&in.ContainerConfig, &out.ContainerConfig); err != nil {
		return err
	}
	if in.Config != nil {
		out.Config = &docker10.DockerConfig{}
		if err := Convert_public_to_api_DockerConfig(in.Config, out.Config); err != nil {
			return err
		}
	}
	return nil
}
func Convert_compatibility_to_api_DockerImage(in *public.DockerV1CompatibilityImage, out *docker10.DockerImage) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = docker10.DockerImage{ID: in.ID, Parent: in.Parent, Comment: in.Comment, Created: metav1.Time{Time: in.Created}, Container: in.Container, DockerVersion: in.DockerVersion, Author: in.Author, Architecture: in.Architecture, Size: in.Size}
	if err := Convert_public_to_api_DockerConfig(&in.ContainerConfig, &out.ContainerConfig); err != nil {
		return err
	}
	if in.Config != nil {
		out.Config = &docker10.DockerConfig{}
		if err := Convert_public_to_api_DockerConfig(in.Config, out.Config); err != nil {
			return err
		}
	}
	return nil
}
func Convert_imageconfig_to_api_DockerImage(in *public.DockerImageConfig, out *docker10.DockerImage) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = docker10.DockerImage{ID: in.ID, Parent: in.Parent, Comment: in.Comment, Created: metav1.Time{Time: in.Created}, Container: in.Container, DockerVersion: in.DockerVersion, Author: in.Author, Architecture: in.Architecture, Size: in.Size}
	if err := Convert_public_to_api_DockerConfig(&in.ContainerConfig, &out.ContainerConfig); err != nil {
		return err
	}
	if in.Config != nil {
		out.Config = &docker10.DockerConfig{}
		if err := Convert_public_to_api_DockerConfig(in.Config, out.Config); err != nil {
			return err
		}
	}
	return nil
}
func Convert_api_to_public_DockerImage(in *docker10.DockerImage, out *public.DockerImage) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = public.DockerImage{ID: in.ID, Parent: in.Parent, Comment: in.Comment, Created: in.Created.Time, Container: in.Container, DockerVersion: in.DockerVersion, Author: in.Author, Architecture: in.Architecture, Size: in.Size}
	if err := Convert_api_to_public_DockerConfig(&in.ContainerConfig, &out.ContainerConfig); err != nil {
		return err
	}
	if in.Config != nil {
		out.Config = &public.DockerConfig{}
		if err := Convert_api_to_public_DockerConfig(in.Config, out.Config); err != nil {
			return err
		}
	}
	return nil
}
func Convert_public_to_api_DockerConfig(in *public.DockerConfig, out *docker10.DockerConfig) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = docker10.DockerConfig{Hostname: in.Hostname, Domainname: in.Domainname, User: in.User, Memory: in.Memory, MemorySwap: in.MemorySwap, CPUShares: in.CPUShares, CPUSet: in.CPUSet, AttachStdin: in.AttachStdin, AttachStdout: in.AttachStdout, AttachStderr: in.AttachStderr, PortSpecs: in.PortSpecs, ExposedPorts: in.ExposedPorts, Tty: in.Tty, OpenStdin: in.OpenStdin, StdinOnce: in.StdinOnce, Env: in.Env, Cmd: in.Cmd, DNS: in.DNS, Image: in.Image, Volumes: in.Volumes, VolumesFrom: in.VolumesFrom, WorkingDir: in.WorkingDir, Entrypoint: in.Entrypoint, NetworkDisabled: in.NetworkDisabled, SecurityOpts: in.SecurityOpts, OnBuild: in.OnBuild, Labels: in.Labels}
	return nil
}
func Convert_api_to_public_DockerConfig(in *docker10.DockerConfig, out *public.DockerConfig) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = public.DockerConfig{Hostname: in.Hostname, Domainname: in.Domainname, User: in.User, Memory: in.Memory, MemorySwap: in.MemorySwap, CPUShares: in.CPUShares, CPUSet: in.CPUSet, AttachStdin: in.AttachStdin, AttachStdout: in.AttachStdout, AttachStderr: in.AttachStderr, PortSpecs: in.PortSpecs, ExposedPorts: in.ExposedPorts, Tty: in.Tty, OpenStdin: in.OpenStdin, StdinOnce: in.StdinOnce, Env: in.Env, Cmd: in.Cmd, DNS: in.DNS, Image: in.Image, Volumes: in.Volumes, VolumesFrom: in.VolumesFrom, WorkingDir: in.WorkingDir, Entrypoint: in.Entrypoint, NetworkDisabled: in.NetworkDisabled, SecurityOpts: in.SecurityOpts, OnBuild: in.OnBuild, Labels: in.Labels}
	return nil
}
