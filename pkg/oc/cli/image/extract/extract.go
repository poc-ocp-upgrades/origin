package extract

import (
	"archive/tar"
	"bytes"
	"net/http"
	"runtime"
	"context"
	"fmt"
	"io"
	"math"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"github.com/openshift/origin/pkg/image/registryclient"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"github.com/docker/distribution"
	dockerarchive "github.com/docker/docker/pkg/archive"
	digest "github.com/opencontainers/go-digest"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/util/templates"
	"github.com/openshift/origin/pkg/image/apis/image/docker10"
	imagereference "github.com/openshift/origin/pkg/image/apis/image/reference"
	"github.com/openshift/origin/pkg/oc/cli/image/archive"
	imagemanifest "github.com/openshift/origin/pkg/oc/cli/image/manifest"
	"github.com/openshift/origin/pkg/oc/cli/image/workqueue"
)

var (
	desc	= templates.LongDesc(`
		Extract the contents of an image to disk

		Download an image or parts of an image to the filesystem. Allows users to access the
		contents of images without requiring a container runtime engine running.

		Pass images to extract as arguments. The --paths flag allows you to define multiple
		source to destination directory mappings. The source section may be either a file, a
		directory (ends with a '/'), or a file pattern within a directory. The destination
		section	is a directory to extract to. Both source and destination must be specified.

		If the specified image supports multiple operating systems, the image that matches the
		current operating system will be chosen. Otherwise you must pass --filter-by-os to
		select the desired image.

		You may further qualify the image by adding a layer selector to the end of the image
		string to only extract specific layers within an image. The supported selectors are:

		  [<index>] - select the layer at the provided index (zero-indexed)
		  [<from_index>,<to_index>] - select layers by index, exclusive
		  [~<prefix>] - select the layer with the matching digest prefix or return an error

		Negative indices are counted from the end of the list, e.g. [-1] selects the last
		layer.`)
	example	= templates.Examples(`
# Extract the busybox image into the current directory
%[1]s docker.io/library/busybox:latest

# Extract the busybox image to a temp directory (must exist)
%[1]s docker.io/library/busybox:latest --path /:/tmp/busybox

# Extract a single file from the image into the current directory
%[1]s docker.io/library/centos:7 --path /bin/bash:.

# Extract all .repo files from the image's /etc/yum.repos.d/ folder.
%[1]s docker.io/library/centos:7 --path /etc/yum.repos.d/*.repo:.

# Extract the last layer in the image
%[1]s docker.io/library/centos:7[-1]

# Extract the first three layers of the image
%[1]s docker.io/library/centos:7[:3]

# Extract the last three layers of the image
%[1]s docker.io/library/centos:7[-3:]
`)
)

type LayerInfo struct {
	Index		int
	Descriptor	distribution.Descriptor
	Mapping		*Mapping
}
type TarEntryFunc func(*tar.Header, LayerInfo, io.Reader) (cont bool, err error)
type Options struct {
	Mappings		[]Mapping
	Files			[]string
	Paths			[]string
	OnlyFiles		bool
	PreservePermissions	bool
	SecurityOptions		imagemanifest.SecurityOptions
	FilterOptions		imagemanifest.FilterOptions
	ParallelOptions		imagemanifest.ParallelOptions
	Confirm			bool
	DryRun			bool
	genericclioptions.IOStreams
	ImageMetadataCallback	func(m *Mapping, dgst, contentDigest digest.Digest, imageConfig *docker10.DockerImageConfig)
	TarEntryCallback	TarEntryFunc
	AllLayers		bool
}

func NewOptions(streams genericclioptions.IOStreams) *Options {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Options{Paths: []string{}, IOStreams: streams, ParallelOptions: imagemanifest.ParallelOptions{MaxPerRegistry: 1}}
}
func New(name string, streams genericclioptions.IOStreams) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o := NewOptions(streams)
	cmd := &cobra.Command{Use: "extract", Short: "Copy files from an image to the filesystem", Long: desc, Example: fmt.Sprintf(example, name+" extract"), Run: func(c *cobra.Command, args []string) {
		kcmdutil.CheckErr(o.Complete(c, args))
		kcmdutil.CheckErr(o.Validate())
		kcmdutil.CheckErr(o.Run())
	}}
	flag := cmd.Flags()
	o.SecurityOptions.Bind(flag)
	o.FilterOptions.Bind(flag)
	flag.BoolVar(&o.Confirm, "confirm", o.Confirm, "Pass to allow extracting to non-empty directories.")
	flag.BoolVar(&o.DryRun, "dry-run", o.DryRun, "Print the actions that would be taken and exit without writing any contents.")
	flag.StringSliceVar(&o.Files, "file", o.Files, "Extract the specified files to the current directory.")
	flag.StringSliceVar(&o.Paths, "path", o.Paths, "Extract only part of an image. Must be SRC:DST where SRC is the path within the image and DST a local directory. If not specified the default is to extract everything to the current directory.")
	flag.BoolVarP(&o.PreservePermissions, "preserve-ownership", "p", o.PreservePermissions, "Preserve the permissions of extracted files.")
	flag.BoolVar(&o.OnlyFiles, "only-files", o.OnlyFiles, "Only extract regular files and directories from the image.")
	flag.BoolVar(&o.AllLayers, "all-layers", o.AllLayers, "For dry-run mode, process from lowest to highest layer and don't omit duplicate files.")
	return cmd
}

type LayerFilter interface {
	Filter(layers []distribution.Descriptor) ([]distribution.Descriptor, error)
}
type Mapping struct {
	Name		string
	Image		string
	ImageRef	imagereference.DockerImageReference
	LayerFilter	LayerFilter
	From		string
	To		string
	ConditionFn	func(m *Mapping, dgst digest.Digest, imageConfig *docker10.DockerImageConfig) (bool, error)
}

func parseMappings(images, paths, files []string, requireEmpty bool) ([]Mapping, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	layerFilter := regexp.MustCompile(`^(.*)\[([^\]]*)\](.*)$`)
	var mappings []Mapping
	for _, image := range images {
		for _, arg := range files {
			if strings.HasSuffix(arg, "/") {
				return nil, fmt.Errorf("invalid file: %s must not end with a slash", arg)
			}
			mappings = append(mappings, Mapping{Image: image, From: strings.TrimPrefix(arg, "/"), To: "."})
		}
		for _, arg := range paths {
			parts := strings.SplitN(arg, ":", 2)
			var mapping Mapping
			switch len(parts) {
			case 2:
				mapping = Mapping{Image: image, From: parts[0], To: parts[1]}
			default:
				return nil, fmt.Errorf("--paths must be of the form SRC:DST")
			}
			if len(mapping.From) > 0 {
				mapping.From = strings.TrimPrefix(mapping.From, "/")
			}
			if len(mapping.To) > 0 {
				fi, err := os.Stat(mapping.To)
				if os.IsNotExist(err) {
					return nil, fmt.Errorf("destination path does not exist: %s", mapping.To)
				}
				if err != nil {
					return nil, fmt.Errorf("invalid argument: %s", err)
				}
				if !fi.IsDir() {
					return nil, fmt.Errorf("invalid argument: %s is not a directory", arg)
				}
				if requireEmpty {
					f, err := os.Open(mapping.To)
					if err != nil {
						return nil, fmt.Errorf("unable to check directory: %v", err)
					}
					names, err := f.Readdirnames(1)
					f.Close()
					if err != nil && err != io.EOF {
						return nil, fmt.Errorf("could not check for empty directory: %v", err)
					}
					if len(names) > 0 {
						return nil, fmt.Errorf("directory %s must be empty, pass --confirm to overwrite contents of directory", mapping.To)
					}
				}
			}
			mappings = append(mappings, mapping)
		}
	}
	for i := range mappings {
		mapping := &mappings[i]
		if matches := layerFilter.FindStringSubmatch(mapping.Image); len(matches) > 0 {
			if len(matches[1]) == 0 || len(matches[2]) == 0 || len(matches[3]) != 0 {
				return nil, fmt.Errorf("layer selectors must be of the form IMAGE[\\d:\\d]")
			}
			mapping.Image = matches[1]
			var err error
			mapping.LayerFilter, err = parseLayerFilter(matches[2])
			if err != nil {
				return nil, err
			}
		}
		src, err := imagereference.Parse(mapping.Image)
		if err != nil {
			return nil, err
		}
		if len(src.Tag) == 0 && len(src.ID) == 0 {
			return nil, fmt.Errorf("source image must point to an image ID or image tag")
		}
		mapping.ImageRef = src
	}
	return mappings, nil
}
func (o *Options) Complete(cmd *cobra.Command, args []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := o.FilterOptions.Complete(cmd.Flags()); err != nil {
		return err
	}
	if len(args) == 0 {
		return fmt.Errorf("you must specify at least one image to extract as an argument")
	}
	if len(o.Paths) == 0 && len(o.Files) == 0 {
		o.Paths = append(o.Paths, "/:.")
	}
	var err error
	o.Mappings, err = parseMappings(args, o.Paths, o.Files, !o.Confirm && !o.DryRun)
	if err != nil {
		return err
	}
	return nil
}
func (o *Options) Validate() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(o.Mappings) == 0 {
		return fmt.Errorf("you must specify one or more paths or files")
	}
	return o.FilterOptions.Validate()
}
func (o *Options) Run() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx := context.Background()
	fromContext, err := o.SecurityOptions.Context()
	if err != nil {
		return err
	}
	stopCh := make(chan struct{})
	defer close(stopCh)
	q := workqueue.New(o.ParallelOptions.MaxPerRegistry, stopCh)
	return q.Try(func(q workqueue.Try) {
		for i := range o.Mappings {
			mapping := o.Mappings[i]
			from := mapping.ImageRef
			q.Try(func() error {
				repo, err := fromContext.Repository(ctx, from.DockerClientDefaults().RegistryURL(), from.RepositoryName(), o.SecurityOptions.Insecure)
				if err != nil {
					return fmt.Errorf("unable to connect to image repository %s: %v", from.Exact(), err)
				}
				srcManifest, location, err := imagemanifest.FirstManifest(ctx, from, repo, o.FilterOptions.Include)
				if err != nil {
					if imagemanifest.IsImageForbidden(err) {
						var msg string
						if len(o.Mappings) == 1 {
							msg = "image does not exist or you don't have permission to access the repository"
						} else {
							msg = fmt.Sprintf("image %q does not exist or you don't have permission to access the repository", from)
						}
						return imagemanifest.NewImageForbidden(msg, err)
					}
					if imagemanifest.IsImageNotFound(err) {
						var msg string
						if len(o.Mappings) == 1 {
							msg = "image does not exist"
						} else {
							msg = fmt.Sprintf("image %q does not exist", from)
						}
						return imagemanifest.NewImageNotFound(msg, err)
					}
					return fmt.Errorf("unable to read image %s: %v", from, err)
				}
				contentDigest, err := registryclient.ContentDigestForManifest(srcManifest, location.Manifest.Algorithm())
				if err != nil {
					return err
				}
				imageConfig, layers, err := imagemanifest.ManifestToImageConfig(ctx, srcManifest, repo.Blobs(ctx), location)
				if err != nil {
					return fmt.Errorf("unable to parse image %s: %v", from, err)
				}
				if mapping.ConditionFn != nil {
					ok, err := mapping.ConditionFn(&mapping, location.Manifest, imageConfig)
					if err != nil {
						return fmt.Errorf("unable to check whether to include image %s: %v", from, err)
					}
					if !ok {
						klog.V(2).Infof("Filtered out image %s with digest %s from being extracted", from, location.Manifest)
						return nil
					}
				}
				var alter alterations
				if o.OnlyFiles {
					alter = append(alter, filesOnly{})
				}
				if len(mapping.From) > 0 {
					switch {
					case strings.HasSuffix(mapping.From, "/"):
						alter = append(alter, newCopyFromDirectory(mapping.From))
					default:
						name, parent := path.Base(mapping.From), path.Dir(mapping.From)
						if name == "." || parent == "." {
							return fmt.Errorf("unexpected directory from mapping %s", mapping.From)
						}
						alter = append(alter, newCopyFromPattern(parent, name))
					}
				}
				filteredLayers := layers
				if mapping.LayerFilter != nil {
					filteredLayers, err = mapping.LayerFilter.Filter(filteredLayers)
					if err != nil {
						return fmt.Errorf("unable to filter layers for %s: %v", from, err)
					}
				}
				if !o.PreservePermissions {
					alter = append(alter, removePermissions{})
				}
				var byEntry TarEntryFunc = o.TarEntryCallback
				if o.DryRun {
					path := mapping.To
					out := o.Out
					byEntry = func(hdr *tar.Header, layerInfo LayerInfo, r io.Reader) (bool, error) {
						if len(hdr.Name) == 0 {
							return true, nil
						}
						mode := hdr.FileInfo().Mode().String()
						switch hdr.Typeflag {
						case tar.TypeDir:
							fmt.Fprintf(out, "%2d %s %12d %s\n", layerInfo.Index, mode, hdr.Size, filepath.Join(path, hdr.Name))
						case tar.TypeReg, tar.TypeRegA:
							fmt.Fprintf(out, "%2d %s %12d %s\n", layerInfo.Index, mode, hdr.Size, filepath.Join(path, hdr.Name))
						case tar.TypeLink:
							fmt.Fprintf(out, "%2d %s %12d %s -> %s\n", layerInfo.Index, mode, hdr.Size, hdr.Name, filepath.Join(path, hdr.Linkname))
						case tar.TypeSymlink:
							fmt.Fprintf(out, "%2d %s %12d %s -> %s\n", layerInfo.Index, mode, hdr.Size, hdr.Name, filepath.Join(path, hdr.Linkname))
						default:
							fmt.Fprintf(out, "%2d %s %12d %s %x\n", layerInfo.Index, mode, hdr.Size, filepath.Join(path, hdr.Name), hdr.Typeflag)
						}
						return true, nil
					}
				}
				alreadySeen := make(map[string]struct{})
				var layerInfos []LayerInfo
				if byEntry != nil && !o.AllLayers {
					for i := len(filteredLayers) - 1; i >= 0; i-- {
						layerInfos = append(layerInfos, LayerInfo{Index: i, Descriptor: filteredLayers[i], Mapping: &mapping})
					}
				} else {
					for i := range filteredLayers {
						layerInfos = append(layerInfos, LayerInfo{Index: i, Descriptor: filteredLayers[i], Mapping: &mapping})
					}
				}
				for _, info := range layerInfos {
					layer := info.Descriptor
					cont, err := func() (bool, error) {
						fromBlobs := repo.Blobs(ctx)
						klog.V(5).Infof("Extracting from layer: %#v", layer)
						r, err := fromBlobs.Open(ctx, layer.Digest)
						if err != nil {
							return false, fmt.Errorf("unable to access the source layer %s: %v", layer.Digest, err)
						}
						defer r.Close()
						options := &archive.TarOptions{AlterHeaders: alter, Chown: o.PreservePermissions}
						if byEntry != nil {
							cont, err := layerByEntry(r, options, info, byEntry, o.AllLayers, alreadySeen)
							if err != nil {
								err = fmt.Errorf("unable to iterate over layer %s from %s: %v", layer.Digest, from.Exact(), err)
							}
							return cont, err
						}
						klog.V(4).Infof("Extracting layer %s with options %#v", layer.Digest, options)
						if _, err := archive.ApplyLayer(mapping.To, r, options); err != nil {
							return false, fmt.Errorf("unable to extract layer %s from %s: %v", layer.Digest, from.Exact(), err)
						}
						return true, nil
					}()
					if err != nil {
						return err
					}
					if !cont {
						break
					}
				}
				if o.ImageMetadataCallback != nil {
					o.ImageMetadataCallback(&mapping, location.Manifest, contentDigest, imageConfig)
				}
				return nil
			})
		}
	})
}
func layerByEntry(r io.Reader, options *archive.TarOptions, layerInfo LayerInfo, fn TarEntryFunc, allLayers bool, alreadySeen map[string]struct{}) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rc, err := dockerarchive.DecompressStream(r)
	if err != nil {
		return false, err
	}
	defer rc.Close()
	tr := tar.NewReader(rc)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				return true, nil
			}
			return false, err
		}
		klog.V(6).Infof("Printing layer entry %#v", hdr)
		if options.AlterHeaders != nil {
			ok, err := options.AlterHeaders.Alter(hdr)
			if err != nil {
				return false, err
			}
			if !ok {
				klog.V(5).Infof("Exclude entry %s %x %d", hdr.Name, hdr.Typeflag, hdr.Size)
				continue
			}
		}
		if _, ok := alreadySeen[hdr.Name]; ok && !allLayers {
			continue
		}
		alreadySeen[hdr.Name] = struct{}{}
		cont, err := fn(hdr, layerInfo, tr)
		if err != nil {
			return false, err
		}
		if !cont {
			return false, nil
		}
	}
}

type alterations []archive.AlterHeader

func (a alterations) Alter(hdr *tar.Header) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, item := range a {
		ok, err := item.Alter(hdr)
		if err != nil {
			return false, err
		}
		if !ok {
			return false, nil
		}
	}
	return true, nil
}

type removePermissions struct{}

func (_ removePermissions) Alter(hdr *tar.Header) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch hdr.Typeflag {
	case tar.TypeReg, tar.TypeRegA:
		hdr.Mode = int64(os.FileMode(0640))
	default:
		hdr.Mode = int64(os.FileMode(0755))
	}
	return true, nil
}

type writableDirectories struct{}

func (_ writableDirectories) Alter(hdr *tar.Header) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch hdr.Typeflag {
	case tar.TypeDir:
		hdr.Mode = int64(os.FileMode(0600) | os.FileMode(hdr.Mode))
	}
	return true, nil
}

type copyFromDirectory struct{ From string }

func newCopyFromDirectory(from string) archive.AlterHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !strings.HasSuffix(from, "/") {
		from = from + "/"
	}
	return &copyFromDirectory{From: from}
}
func (n *copyFromDirectory) Alter(hdr *tar.Header) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return changeTarEntryParent(hdr, n.From), nil
}

type copyFromPattern struct {
	Base	string
	Name	string
}

func newCopyFromPattern(dir, name string) archive.AlterHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !strings.HasSuffix(dir, "/") {
		dir = dir + "/"
	}
	return &copyFromPattern{Base: dir, Name: name}
}
func (n *copyFromPattern) Alter(hdr *tar.Header) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !changeTarEntryParent(hdr, n.Base) {
		return false, nil
	}
	matchName := hdr.Name
	if i := strings.Index(matchName, "/"); i != -1 {
		matchName = matchName[:i]
	}
	if ok, err := path.Match(n.Name, matchName); !ok || err != nil {
		klog.V(5).Infof("Excluded %s due to filter %s", hdr.Name, n.Name)
		return false, err
	}
	return true, nil
}
func changeTarEntryParent(hdr *tar.Header, from string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !strings.HasPrefix(hdr.Name, from) {
		klog.V(5).Infof("Exclude %s due to missing prefix %s", hdr.Name, from)
		return false
	}
	if len(hdr.Linkname) > 0 {
		if strings.HasPrefix(hdr.Linkname, from) {
			hdr.Linkname = strings.TrimPrefix(hdr.Linkname, from)
			klog.V(5).Infof("Updated link to %s", hdr.Linkname)
		} else {
			klog.V(4).Infof("Name %s won't correctly point to %s outside of %s", hdr.Name, hdr.Linkname, from)
		}
	}
	hdr.Name = strings.TrimPrefix(hdr.Name, from)
	klog.V(5).Infof("Updated name %s", hdr.Name)
	return true
}

type filesOnly struct{}

func (_ filesOnly) Alter(hdr *tar.Header) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch hdr.Typeflag {
	case tar.TypeReg, tar.TypeRegA, tar.TypeDir:
		return true, nil
	default:
		klog.V(6).Infof("Excluded %s because type was not a regular file or directory: %x", hdr.Name, hdr.Typeflag)
		return false, nil
	}
}
func parseLayerFilter(s string) (LayerFilter, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if strings.HasPrefix(s, "~") {
		s = s[1:]
		return &prefixLayerFilter{Prefix: s}, nil
	}
	if strings.Contains(s, ":") {
		l := &indexLayerFilter{From: 0, To: math.MaxInt32}
		parts := strings.SplitN(s, ":", 2)
		if len(parts[0]) > 0 {
			i, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, fmt.Errorf("[from:to] must have valid numbers: %v", err)
			}
			l.From = int32(i)
		}
		if len(parts[1]) > 0 {
			i, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("[from:to] must have valid numbers: %v", err)
			}
			l.To = int32(i)
		}
		if l.To > 0 && l.To < l.From {
			return nil, fmt.Errorf("[from:to] to must be larger than from")
		}
		return l, nil
	}
	if i, err := strconv.Atoi(s); err == nil {
		l := NewPositionLayerFilter(int32(i))
		return l, nil
	}
	return nil, fmt.Errorf("the layer selector [%s] is not valid, must be [from:to], [index], or [~digest]", s)
}

type prefixLayerFilter struct{ Prefix string }

func (s *prefixLayerFilter) Filter(layers []distribution.Descriptor) ([]distribution.Descriptor, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var filtered []distribution.Descriptor
	for _, d := range layers {
		if strings.HasPrefix(d.Digest.String(), s.Prefix) {
			filtered = append(filtered, d)
		}
	}
	if len(filtered) == 0 {
		return nil, fmt.Errorf("no layers start with '%s'", s.Prefix)
	}
	if len(filtered) > 1 {
		return nil, fmt.Errorf("multiple layers start with '%s', you must be more specific", s.Prefix)
	}
	return filtered, nil
}

type indexLayerFilter struct {
	From	int32
	To	int32
}

func (s *indexLayerFilter) Filter(layers []distribution.Descriptor) ([]distribution.Descriptor, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := int32(len(layers))
	from := s.From
	to := s.To
	if from < 0 {
		from = l + from
	}
	if to < 0 {
		to = l + to
	}
	if to > l {
		to = l
	}
	if from < 0 || to < 0 || from >= l {
		if s.To == math.MaxInt32 {
			return nil, fmt.Errorf("tried to select [%d:], but image only has %d layers", s.From, l)
		}
		return nil, fmt.Errorf("tried to select [%d:%d], but image only has %d layers", s.From, s.To, l)
	}
	if to < from {
		to, from = from, to
	}
	return layers[from:to], nil
}

type positionLayerFilter struct{ At int32 }

func NewPositionLayerFilter(at int32) LayerFilter {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &positionLayerFilter{at}
}
func (s *positionLayerFilter) Filter(layers []distribution.Descriptor) ([]distribution.Descriptor, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := int32(len(layers))
	at := s.At
	if at < 0 {
		at = l + s.At
	}
	if at < 0 || at >= l {
		return nil, fmt.Errorf("tried to select layer %d, but image only has %d layers", s.At, l)
	}
	return []distribution.Descriptor{layers[at]}, nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
