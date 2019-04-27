package verifyimagesignature

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"github.com/containers/image/docker/policyconfiguration"
	"github.com/containers/image/docker/reference"
	"github.com/containers/image/signature"
	sigtypes "github.com/containers/image/types"
	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/util/templates"
	imagev1 "github.com/openshift/api/image/v1"
	imagev1typedclient "github.com/openshift/client-go/image/clientset/versioned/typed/image/v1"
	userv1typedclient "github.com/openshift/client-go/user/clientset/versioned/typed/user/v1"
	imageref "github.com/openshift/library-go/pkg/image/reference"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
)

var (
	verifyImageSignatureLongDesc	= templates.LongDesc(`
	Verifies the image signature of an image imported to internal registry using the local public GPG key.

	This command verifies if the image identity contained in the image signature can be trusted
	by using the public GPG key to verify the signature itself and matching the provided expected identity
	with the identity (pull spec) of the given image.
	By default, this command will use the public GPG keyring located in "$GNUPGHOME/.gnupg/pubring.gpg"

	By default, this command will not save the result of the verification back to the image object, to do so
	user have to specify the "--save" flag. Note that to modify the image signature verification status,
	user have to have permissions to edit an image object (usually an "image-auditor" role).

	Note that using the "--save" flag on already verified image together with invalid GPG
	key or invalid expected identity will cause the saved verification status to be removed
	and the image will become "unverified".

	If this command is outside the cluster, users have to specify the "--registry-url" parameter
	with the public URL of image registry.

	To remove all verifications, users can use the "--remove-all" flag.
	`)
	verifyImageSignatureExample	= templates.Examples(`
	# Verify the image signature and identity using the local GPG keychain
	%[1]s sha256:c841e9b64e4579bd56c794bdd7c36e1c257110fd2404bebbb8b613e4935228c4 \
			--expected-identity=registry.local:5000/foo/bar:v1

	# Verify the image signature and identity using the local GPG keychain and save the status
	%[1]s sha256:c841e9b64e4579bd56c794bdd7c36e1c257110fd2404bebbb8b613e4935228c4 \
			--expected-identity=registry.local:5000/foo/bar:v1 --save

	# Verify the image signature and identity via exposed registry route
	%[1]s sha256:c841e9b64e4579bd56c794bdd7c36e1c257110fd2404bebbb8b613e4935228c4 \
			--expected-identity=registry.local:5000/foo/bar:v1 \
			--registry-url=docker-registry.foo.com

	# Remove all signature verifications from the image
	%[1]s sha256:c841e9b64e4579bd56c794bdd7c36e1c257110fd2404bebbb8b613e4935228c4 --remove-all
	`)
)

const (
	VerifyRecommendedName = "verify-image-signature"
)

type VerifyImageSignatureOptions struct {
	InputImage		string
	ExpectedIdentity	string
	PublicKeyFilename	string
	PublicKey		[]byte
	Save			bool
	RemoveAll		bool
	CurrentUser		string
	CurrentUserToken	string
	RegistryURL		string
	Insecure		bool
	ImageClient		imagev1typedclient.ImageV1Interface
	genericclioptions.IOStreams
}

func NewVerifyImageSignatureOptions(streams genericclioptions.IOStreams) *VerifyImageSignatureOptions {
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
	return &VerifyImageSignatureOptions{PublicKeyFilename: filepath.Join(os.Getenv("GNUPGHOME"), "pubring.gpg"), IOStreams: streams}
}
func NewCmdVerifyImageSignature(name, fullName string, f kcmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
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
	o := NewVerifyImageSignatureOptions(streams)
	cmd := &cobra.Command{Use: fmt.Sprintf("%s IMAGE --expected-identity=EXPECTED_IDENTITY [--save]", VerifyRecommendedName), Short: "Verify the image identity contained in the image signature", Long: verifyImageSignatureLongDesc, Example: fmt.Sprintf(verifyImageSignatureExample, fullName), Run: func(cmd *cobra.Command, args []string) {
		kcmdutil.CheckErr(o.Validate())
		kcmdutil.CheckErr(o.Complete(f, cmd, args))
		kcmdutil.CheckErr(o.Run())
	}}
	cmd.Flags().StringVar(&o.ExpectedIdentity, "expected-identity", o.ExpectedIdentity, "An expected image docker reference to verify (required).")
	cmd.Flags().BoolVar(&o.Save, "save", o.Save, "If true, the result of the verification will be saved to an image object.")
	cmd.Flags().BoolVar(&o.RemoveAll, "remove-all", o.RemoveAll, "If set, all signature verifications will be removed from the given image.")
	cmd.Flags().StringVar(&o.PublicKeyFilename, "public-key", o.PublicKeyFilename, fmt.Sprintf("A path to a public GPG key to be used for verification. (defaults to %q)", o.PublicKeyFilename))
	cmd.Flags().StringVar(&o.RegistryURL, "registry-url", o.RegistryURL, "The address to use when contacting the registry, instead of using the internal cluster address. This is useful if you can't resolve or reach the internal registry address.")
	cmd.Flags().BoolVar(&o.Insecure, "insecure", o.Insecure, "If set, use the insecure protocol for registry communication.")
	return cmd
}
func (o *VerifyImageSignatureOptions) Validate() error {
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
	if !o.RemoveAll {
		if len(o.ExpectedIdentity) == 0 {
			return errors.New("the --expected-identity is required")
		}
		if _, err := imageref.Parse(o.ExpectedIdentity); err != nil {
			return errors.New("the --expected-identity must be valid image reference")
		}
	}
	if o.RemoveAll && len(o.ExpectedIdentity) > 0 {
		return errors.New("the --expected-identity cannot be used when removing all verifications")
	}
	return nil
}
func (o *VerifyImageSignatureOptions) Complete(f kcmdutil.Factory, cmd *cobra.Command, args []string) error {
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
	if len(args) != 1 {
		return kcmdutil.UsageErrorf(cmd, "exactly one image must be specified")
	}
	o.InputImage = args[0]
	var err error
	if len(o.PublicKeyFilename) > 0 {
		if o.PublicKey, err = ioutil.ReadFile(o.PublicKeyFilename); err != nil {
			return fmt.Errorf("unable to read --public-key: %v", err)
		}
	}
	clientConfig, err := f.ToRESTConfig()
	if err != nil {
		return err
	}
	o.ImageClient, err = imagev1typedclient.NewForConfig(clientConfig)
	if err != nil {
		return err
	}
	userClient, err := userv1typedclient.NewForConfig(clientConfig)
	if err != nil {
		return err
	}
	if me, err := userClient.Users().Get("~", metav1.GetOptions{}); err != nil {
		return err
	} else {
		o.CurrentUser = me.Name
		if config, err := f.ToRESTConfig(); err != nil {
			return err
		} else {
			if o.CurrentUserToken = config.BearerToken; len(o.CurrentUserToken) == 0 {
				return fmt.Errorf("no token is currently in use for this session")
			}
		}
	}
	return nil
}
func (o VerifyImageSignatureOptions) Run() error {
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
	img, err := o.ImageClient.Images().Get(o.InputImage, metav1.GetOptions{})
	if err != nil {
		return err
	}
	if len(img.Signatures) == 0 {
		return fmt.Errorf("%s does not have any signature", img.Name)
	}
	pr, err := signature.NewPRSignedByKeyPath(signature.SBKeyTypeGPGKeys, o.PublicKeyFilename, signature.NewPRMMatchRepoDigestOrExact())
	if err != nil {
		return fmt.Errorf("unable to prepare verification policy requirements: %v", err)
	}
	policy := signature.Policy{Default: []signature.PolicyRequirement{pr}}
	pc, err := signature.NewPolicyContext(&policy)
	if err != nil {
		return fmt.Errorf("unable to setup policy: %v", err)
	}
	defer pc.Destroy()
	if o.RemoveAll {
		img.Signatures = []imagev1.ImageSignature{}
	}
	for i, s := range img.Signatures {
		signedBy, err := o.verifySignature(pc, img, s.Content)
		if err != nil {
			fmt.Fprintf(o.ErrOut, "error verifying signature %s for image %s (verification status will be removed): %v\n", img.Signatures[i].Name, o.InputImage, err)
			img.Signatures[i] = imagev1.ImageSignature{}
			continue
		}
		fmt.Fprintf(o.Out, "image %q identity is now confirmed (signed by GPG key %q)\n", o.InputImage, signedBy)
		now := metav1.Now()
		newConditions := []imagev1.SignatureCondition{{Type: imageapi.SignatureTrusted, Status: corev1.ConditionTrue, LastProbeTime: now, LastTransitionTime: now, Reason: "manually verified", Message: fmt.Sprintf("verified by user %q", o.CurrentUser)}, {Type: imageapi.SignatureForImage, Status: corev1.ConditionTrue, LastProbeTime: now, LastTransitionTime: now}}
		img.Signatures[i].Conditions = newConditions
		img.Signatures[i].IssuedBy = &imagev1.SignatureIssuer{}
		img.Signatures[i].IssuedBy.CommonName = signedBy
	}
	if o.Save || o.RemoveAll {
		_, err := o.ImageClient.Images().Update(img)
		return err
	} else {
		fmt.Fprintf(o.Out, "Neither --save nor --remove-all were passed, image %q not updated to %v\n", o.InputImage, img)
	}
	return nil
}
func (o *VerifyImageSignatureOptions) getImageManifest(img *imagev1.Image) ([]byte, error) {
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
	parsed, err := imageapi.ParseDockerImageReference(img.DockerImageReference)
	if err != nil {
		return nil, err
	}
	registryURL := parsed.RegistryURL()
	if len(o.RegistryURL) > 0 {
		registryURL = &url.URL{Host: o.RegistryURL, Scheme: "https"}
		if o.Insecure {
			registryURL.Scheme = ""
		}
	}
	return getImageManifestByIDFromRegistry(registryURL, parsed.RepositoryName(), img.Name, o.CurrentUser, o.CurrentUserToken, o.Insecure)
}
func (o *VerifyImageSignatureOptions) verifySignature(pc *signature.PolicyContext, img *imagev1.Image, sigBlob []byte) (string, error) {
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
	manifest, err := o.getImageManifest(img)
	if err != nil {
		return "", fmt.Errorf("failed to get image %q manifest: %v", img.Name, err)
	}
	allowed, err := pc.IsRunningImageAllowed(newUnparsedImage(o.ExpectedIdentity, sigBlob, manifest))
	if !allowed && err == nil {
		return "", errors.New("signature rejected but no error set")
	}
	if err != nil {
		return "", fmt.Errorf("signature rejected: %v", err)
	}
	if untrustedInfo, err := signature.GetUntrustedSignatureInformationWithoutVerifying(sigBlob); err != nil {
		return "", fmt.Errorf("error getting signing key identity: %v", err)
	} else {
		return untrustedInfo.UntrustedShortKeyIdentifier, nil
	}
}

var dummyDockerTransport = dockerTransport{}

type dockerTransport struct{}

func (t dockerTransport) Name() string {
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
	return "docker"
}
func (t dockerTransport) ParseReference(reference string) (sigtypes.ImageReference, error) {
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
	return parseDockerReference(reference)
}
func (t dockerTransport) ValidatePolicyConfigurationScope(scope string) error {
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
	return nil
}

type dummyDockerReference struct{ ref reference.Named }

func parseDockerReference(refString string) (sigtypes.ImageReference, error) {
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
	if !strings.HasPrefix(refString, "//") {
		return nil, fmt.Errorf("docker: image reference %s does not start with //", refString)
	}
	ref, err := reference.ParseNormalizedNamed(strings.TrimPrefix(refString, "//"))
	if err != nil {
		return nil, err
	}
	ref = reference.TagNameOnly(ref)
	if reference.IsNameOnly(ref) {
		return nil, fmt.Errorf("Docker reference %s has neither a tag nor a digest", reference.FamiliarString(ref))
	}
	_, isTagged := ref.(reference.NamedTagged)
	_, isDigested := ref.(reference.Canonical)
	if isTagged && isDigested {
		return nil, fmt.Errorf("Docker references with both a tag and digest are currently not supported")
	}
	return dummyDockerReference{ref: ref}, nil
}
func (ref dummyDockerReference) Transport() sigtypes.ImageTransport {
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
	return dummyDockerTransport
}
func (ref dummyDockerReference) StringWithinTransport() string {
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
	return "//" + reference.FamiliarString(ref.ref)
}
func (ref dummyDockerReference) DockerReference() reference.Named {
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
	return ref.ref
}
func (ref dummyDockerReference) PolicyConfigurationIdentity() string {
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
	res, err := policyconfiguration.DockerReferenceIdentity(ref.ref)
	if res == "" || err != nil {
		panic(fmt.Sprintf("Internal inconsistency: policyconfiguration.DockerReferenceIdentity returned %#v, %v", res, err))
	}
	return res
}
func (ref dummyDockerReference) PolicyConfigurationNamespaces() []string {
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
	return policyconfiguration.DockerReferenceNamespaces(ref.ref)
}
func (ref dummyDockerReference) NewImage(ctx *sigtypes.SystemContext) (sigtypes.Image, error) {
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
	panic("Unimplemented")
}
func (ref dummyDockerReference) NewImageSource(ctx *sigtypes.SystemContext, requestedManifestMIMETypes []string) (sigtypes.ImageSource, error) {
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
	panic("Unimplemented")
}
func (ref dummyDockerReference) NewImageDestination(ctx *sigtypes.SystemContext) (sigtypes.ImageDestination, error) {
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
	panic("Unimplemented")
}
func (ref dummyDockerReference) DeleteImage(ctx *sigtypes.SystemContext) error {
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
	panic("Unimplemented")
}

type unparsedImage struct {
	ref		sigtypes.ImageReference
	manifest	[]byte
	signature	[]byte
}

func newUnparsedImage(expectedIdentity string, signature, manifest []byte) sigtypes.UnparsedImage {
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
	ref, _ := parseDockerReference("//" + expectedIdentity)
	return &unparsedImage{ref: ref, manifest: manifest, signature: signature}
}
func (ui *unparsedImage) Reference() sigtypes.ImageReference {
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
	return ui.ref
}
func (ui *unparsedImage) Close() error {
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
	return nil
}
func (ui *unparsedImage) Manifest() ([]byte, string, error) {
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
	return ui.manifest, "", nil
}
func (ui *unparsedImage) Signatures(context.Context) ([][]byte, error) {
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
	return [][]byte{ui.signature}, nil
}
