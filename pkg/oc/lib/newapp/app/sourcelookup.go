package app

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"k8s.io/klog"
	s2iapi "github.com/openshift/source-to-image/pkg/api"
	s2igit "github.com/openshift/source-to-image/pkg/scm/git"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/apis/core/validation"
	buildv1 "github.com/openshift/api/build/v1"
	"github.com/openshift/library-go/pkg/git"
	"github.com/openshift/origin/pkg/oc/lib/newapp"
	"github.com/openshift/origin/pkg/oc/lib/newapp/source"
)

type Dockerfile interface {
	AST() *parser.Node
	Contents() string
}

func NewDockerfileFromFile(path string) (Dockerfile, error) {
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
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("Dockerfile %q is empty", path)
	}
	return NewDockerfile(string(data))
}
func NewDockerfile(contents string) (Dockerfile, error) {
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
	if len(contents) == 0 {
		return nil, errors.New("Dockerfile is empty")
	}
	node, err := parser.Parse(strings.NewReader(contents))
	if err != nil {
		return nil, err
	}
	return dockerfileContents{node.AST, contents}, nil
}

type dockerfileContents struct {
	ast		*parser.Node
	contents	string
}

func (d dockerfileContents) AST() *parser.Node {
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
	return d.ast
}
func (d dockerfileContents) Contents() string {
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
	return d.contents
}
func IsRemoteRepository(s string) (bool, error) {
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
	url, err := s2igit.Parse(s)
	if err != nil {
		klog.V(5).Infof("%s is not a valid url: %v", s, err)
		return false, err
	}
	if url.IsLocal() {
		klog.V(5).Infof("%s is not a valid remote git clone spec", s)
		return false, nil
	}
	gitRepo := git.NewRepository()
	for i := 0; i < 3; i++ {
		_, _, err = gitRepo.ListRemote(url.StringNoFragment())
		if err == nil {
			break
		}
	}
	if err != nil {
		klog.V(5).Infof("could not list git remotes for %s: %v", s, err)
		return false, err
	}
	klog.V(5).Infof("%s is a valid remote git repository", s)
	return true, nil
}

type SourceRepository struct {
	location		string
	url			s2igit.URL
	localDir		string
	remoteURL		*s2igit.URL
	contextDir		string
	secrets			[]buildv1.SecretBuildSource
	configMaps		[]buildv1.ConfigMapBuildSource
	info			*SourceRepositoryInfo
	sourceImage		ComponentReference
	sourceImageFrom		string
	sourceImageTo		string
	usedBy			[]ComponentReference
	strategy		newapp.Strategy
	ignoreRepository	bool
	binary			bool
	forceAddDockerfile	bool
	requiresAuth		bool
}

func NewSourceRepository(s string, strategy newapp.Strategy) (*SourceRepository, error) {
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
	location, err := s2igit.Parse(s)
	if err != nil {
		return nil, err
	}
	return &SourceRepository{location: s, url: *location, strategy: strategy}, nil
}
func NewSourceRepositoryWithDockerfile(s, dockerfilePath string) (*SourceRepository, error) {
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
	r, err := NewSourceRepository(s, newapp.StrategyDocker)
	if err != nil {
		return nil, err
	}
	if len(dockerfilePath) == 0 {
		dockerfilePath = "Dockerfile"
	}
	f, err := NewDockerfileFromFile(filepath.Join(s, dockerfilePath))
	if err != nil {
		return nil, err
	}
	if r.info == nil {
		r.info = &SourceRepositoryInfo{}
	}
	r.info.Dockerfile = f
	return r, nil
}
func NewSourceRepositoryForDockerfile(contents string) (*SourceRepository, error) {
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
	s := &SourceRepository{ignoreRepository: true, strategy: newapp.StrategyDocker}
	err := s.AddDockerfile(contents)
	return s, err
}
func NewBinarySourceRepository(strategy newapp.Strategy) *SourceRepository {
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
	return &SourceRepository{binary: true, ignoreRepository: true, strategy: strategy}
}
func NewImageSourceRepository(compRef ComponentReference, from, to string) *SourceRepository {
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
	return &SourceRepository{sourceImage: compRef, sourceImageFrom: from, sourceImageTo: to, ignoreRepository: true, location: compRef.Input().From, strategy: newapp.StrategySource}
}
func (r *SourceRepository) UsedBy(ref ComponentReference) {
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
	r.usedBy = append(r.usedBy, ref)
}
func (r *SourceRepository) Remote() bool {
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
	return !r.url.IsLocal()
}
func (r *SourceRepository) InUse() bool {
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
	return len(r.usedBy) > 0
}
func (r *SourceRepository) SetStrategy(strategy newapp.Strategy) {
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
	r.strategy = strategy
}
func (r *SourceRepository) GetStrategy() newapp.Strategy {
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
	return r.strategy
}
func (r *SourceRepository) String() string {
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
	return r.location
}
func (r *SourceRepository) Detect(d Detector, dockerStrategy bool) error {
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
	if r.info != nil {
		return nil
	}
	path, err := r.LocalPath()
	if err != nil {
		return err
	}
	r.info, err = d.Detect(path, dockerStrategy)
	if err != nil {
		return err
	}
	if err = r.DetectAuth(); err != nil {
		return err
	}
	return nil
}
func (r *SourceRepository) SetInfo(info *SourceRepositoryInfo) {
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
	r.info = info
}
func (r *SourceRepository) Info() *SourceRepositoryInfo {
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
	return r.info
}
func (r *SourceRepository) LocalPath() (string, error) {
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
	if len(r.localDir) > 0 {
		return r.localDir, nil
	}
	if r.url.IsLocal() {
		r.localDir = filepath.Join(r.url.LocalPath(), r.contextDir)
	} else {
		gitRepo := git.NewRepository()
		var err error
		if r.localDir, err = ioutil.TempDir("", "gen"); err != nil {
			return "", err
		}
		r.localDir, err = CloneAndCheckoutSources(gitRepo, r.url.StringNoFragment(), r.url.URL.Fragment, r.localDir, r.contextDir)
		if err != nil {
			return "", err
		}
	}
	if _, err := os.Stat(r.localDir); os.IsNotExist(err) {
		return "", fmt.Errorf("supplied context directory '%s' does not exist in '%s'", r.contextDir, r.url.String())
	}
	return r.localDir, nil
}
func (r *SourceRepository) DetectAuth() error {
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
	url, ok, err := r.RemoteURL()
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}
	tempHome, err := ioutil.TempDir("", "githome")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempHome)
	tempSrc, err := ioutil.TempDir("", "gen")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempSrc)
	env := []string{fmt.Sprintf("HOME=%s", tempHome), "GIT_SSH=/dev/null", "GIT_CONFIG_NOSYSTEM=true", "GIT_ASKPASS=true"}
	if runtime.GOOS == "windows" {
		env = append(env, fmt.Sprintf("ProgramData=%s", os.Getenv("ProgramData")), fmt.Sprintf("SystemRoot=%s", os.Getenv("SystemRoot")))
	}
	gitRepo := git.NewRepositoryWithEnv(env)
	klog.V(4).Infof("Checking if %v requires authentication", url.StringNoFragment())
	_, _, err = gitRepo.TimedListRemote(10*time.Second, url.StringNoFragment(), "--heads")
	if err != nil {
		r.requiresAuth = true
		fmt.Print("warning: Cannot check if git requires authentication.\n")
	}
	return nil
}
func (r *SourceRepository) RemoteURL() (*s2igit.URL, bool, error) {
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
	if r.remoteURL != nil {
		return r.remoteURL, true, nil
	}
	if r.url.IsLocal() {
		gitRepo := git.NewRepository()
		remote, ok, err := gitRepo.GetOriginURL(r.url.LocalPath())
		if err != nil && err != git.ErrGitNotAvailable {
			return nil, false, err
		}
		if !ok {
			return nil, ok, nil
		}
		ref := gitRepo.GetRef(r.url.LocalPath())
		if len(ref) > 0 {
			remote = fmt.Sprintf("%s#%s", remote, ref)
		}
		if r.remoteURL, err = s2igit.Parse(remote); err != nil {
			return nil, false, err
		}
	} else {
		r.remoteURL = &r.url
	}
	return r.remoteURL, true, nil
}
func (r *SourceRepository) SetContextDir(dir string) {
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
	r.contextDir = dir
}
func (r *SourceRepository) ContextDir() string {
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
	return r.contextDir
}
func (r *SourceRepository) ConfigMaps() []buildv1.ConfigMapBuildSource {
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
	return r.configMaps
}
func (r *SourceRepository) Secrets() []buildv1.SecretBuildSource {
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
	return r.secrets
}
func (r *SourceRepository) SetSourceImage(c ComponentReference) {
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
	r.sourceImage = c
}
func (r *SourceRepository) SetSourceImagePath(source, dest string) {
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
	r.sourceImageFrom = source
	r.sourceImageTo = dest
}
func (r *SourceRepository) AddDockerfile(contents string) error {
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
	dockerfile, err := NewDockerfile(contents)
	if err != nil {
		return err
	}
	if r.info == nil {
		r.info = &SourceRepositoryInfo{}
	}
	r.info.Dockerfile = dockerfile
	r.SetStrategy(newapp.StrategyDocker)
	r.forceAddDockerfile = true
	return nil
}
func (r *SourceRepository) AddBuildConfigMaps(configMaps []string) error {
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
	injections := s2iapi.VolumeList{}
	r.configMaps = []buildv1.ConfigMapBuildSource{}
	for _, in := range configMaps {
		if err := injections.Set(in); err != nil {
			return err
		}
	}
	configMapExists := func(name string) bool {
		for _, c := range r.configMaps {
			if c.ConfigMap.Name == name {
				return true
			}
		}
		return false
	}
	for _, in := range injections {
		if r.GetStrategy() == newapp.StrategyDocker && filepath.IsAbs(in.Destination) {
			return fmt.Errorf("for the docker strategy, the configMap destination directory %q must be a relative path", in.Destination)
		}
		if len(validation.ValidateConfigMapName(in.Source, false)) != 0 {
			return fmt.Errorf("the %q must be a valid configMap name", in.Source)
		}
		if configMapExists(in.Source) {
			return fmt.Errorf("the %q configMap can be used just once", in.Source)
		}
		r.configMaps = append(r.configMaps, buildv1.ConfigMapBuildSource{ConfigMap: corev1.LocalObjectReference{Name: in.Source}, DestinationDir: in.Destination})
	}
	return nil
}
func (r *SourceRepository) AddBuildSecrets(secrets []string) error {
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
	injections := s2iapi.VolumeList{}
	r.secrets = []buildv1.SecretBuildSource{}
	for _, in := range secrets {
		if err := injections.Set(in); err != nil {
			return err
		}
	}
	secretExists := func(name string) bool {
		for _, s := range r.secrets {
			if s.Secret.Name == name {
				return true
			}
		}
		return false
	}
	for _, in := range injections {
		if r.GetStrategy() == newapp.StrategyDocker && filepath.IsAbs(in.Destination) {
			return fmt.Errorf("for the docker strategy, the secret destination directory %q must be a relative path", in.Destination)
		}
		if len(validation.ValidateSecretName(in.Source, false)) != 0 {
			return fmt.Errorf("the %q must be valid secret name", in.Source)
		}
		if secretExists(in.Source) {
			return fmt.Errorf("the %q secret can be used just once", in.Source)
		}
		r.secrets = append(r.secrets, buildv1.SecretBuildSource{Secret: corev1.LocalObjectReference{Name: in.Source}, DestinationDir: in.Destination})
	}
	return nil
}

type SourceRepositories []*SourceRepository

func (rr SourceRepositories) String() string {
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
	repos := []string{}
	for _, r := range rr {
		repos = append(repos, r.String())
	}
	return strings.Join(repos, ",")
}
func (rr SourceRepositories) NotUsed() SourceRepositories {
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
	notUsed := SourceRepositories{}
	for _, r := range rr {
		if !r.InUse() {
			notUsed = append(notUsed, r)
		}
	}
	return notUsed
}

type SourceRepositoryInfo struct {
	Path		string
	Types		[]SourceLanguageType
	Dockerfile	Dockerfile
	Jenkinsfile	bool
}

func (info *SourceRepositoryInfo) Terms() []string {
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
	terms := []string{}
	for i := range info.Types {
		terms = append(terms, info.Types[i].Term())
	}
	return terms
}

type SourceLanguageType struct {
	Platform	string
	Version		string
}

func (t *SourceLanguageType) Term() string {
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
	if len(t.Version) == 0 {
		return t.Platform
	}
	return fmt.Sprintf("%s:%s", t.Platform, t.Version)
}

type Detector interface {
	Detect(dir string, dockerStrategy bool) (*SourceRepositoryInfo, error)
}
type SourceRepositoryEnumerator struct {
	Detectors		source.Detectors
	DockerfileTester	newapp.Tester
	JenkinsfileTester	newapp.Tester
}

func (e SourceRepositoryEnumerator) Detect(dir string, noSourceDetection bool) (*SourceRepositoryInfo, error) {
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
	info := &SourceRepositoryInfo{Path: dir}
	if !noSourceDetection {
		for _, d := range e.Detectors {
			if detected := d(dir); detected != nil {
				info.Types = append(info.Types, SourceLanguageType{Platform: detected.Platform, Version: detected.Version})
			}
		}
	}
	if path, ok, err := e.DockerfileTester.Has(dir); err == nil && ok {
		dockerfile, err := NewDockerfileFromFile(path)
		if err != nil {
			return nil, err
		}
		info.Dockerfile = dockerfile
	}
	if _, ok, err := e.JenkinsfileTester.Has(dir); err == nil && ok {
		info.Jenkinsfile = true
	}
	return info, nil
}
func StrategyAndSourceForRepository(repo *SourceRepository, image *ImageRef) (*BuildStrategyRef, *SourceRef, error) {
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
	strategy := &BuildStrategyRef{Base: image, Strategy: repo.strategy}
	source := &SourceRef{Binary: repo.binary, Secrets: repo.secrets, ConfigMaps: repo.configMaps, RequiresAuth: repo.requiresAuth}
	if repo.sourceImage != nil {
		srcImageRef, err := InputImageFromMatch(repo.sourceImage.Input().ResolvedMatch)
		if err != nil {
			return nil, nil, err
		}
		source.SourceImage = srcImageRef
		source.ImageSourcePath = repo.sourceImageFrom
		source.ImageDestPath = repo.sourceImageTo
	}
	if (repo.ignoreRepository || repo.forceAddDockerfile) && repo.Info() != nil && repo.Info().Dockerfile != nil {
		source.DockerfileContents = repo.Info().Dockerfile.Contents()
	}
	if !repo.ignoreRepository {
		remoteURL, ok, err := repo.RemoteURL()
		if err != nil {
			return nil, nil, fmt.Errorf("cannot obtain remote URL for repository at %s", repo.location)
		}
		if ok {
			source.URL = remoteURL
		} else {
			source.Binary = true
		}
		source.ContextDir = repo.ContextDir()
	}
	return strategy, source, nil
}
func CloneAndCheckoutSources(repo git.Repository, remote, ref, localDir, contextDir string) (string, error) {
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
	if len(ref) == 0 {
		klog.V(5).Infof("No source ref specified, using shallow git clone")
		if err := repo.CloneWithOptions(localDir, remote, git.Shallow, "--recursive"); err != nil {
			return "", fmt.Errorf("shallow cloning repository %q to %q failed: %v", remote, localDir, err)
		}
	} else {
		klog.V(5).Infof("Requested ref %q, performing full git clone and git checkout", ref)
		if err := repo.Clone(localDir, remote); err != nil {
			return "", fmt.Errorf("cloning repository %q to %q failed: %v", remote, localDir, err)
		}
	}
	if len(ref) > 0 {
		if err := repo.Checkout(localDir, ref); err != nil {
			err = repo.PotentialPRRetryAsFetch(localDir, remote, ref, err)
			if err != nil {
				return "", fmt.Errorf("unable to checkout ref %q in %q repository: %v", ref, remote, err)
			}
		}
	}
	if len(contextDir) > 0 {
		klog.V(5).Infof("Using context directory %q. The full source path is %q", contextDir, filepath.Join(localDir, contextDir))
	}
	return filepath.Join(localDir, contextDir), nil
}
