package build

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	"time"
)

const (
	BuildAnnotation                           = "openshift.io/build.name"
	BuildConfigAnnotation                     = "openshift.io/build-config.name"
	BuildNumberAnnotation                     = "openshift.io/build.number"
	BuildCloneAnnotation                      = "openshift.io/build.clone-of"
	BuildPodNameAnnotation                    = "openshift.io/build.pod-name"
	BuildJenkinsStatusJSONAnnotation          = "openshift.io/jenkins-status-json"
	BuildJenkinsLogURLAnnotation              = "openshift.io/jenkins-log-url"
	BuildJenkinsConsoleLogURLAnnotation       = "openshift.io/jenkins-console-log-url"
	BuildJenkinsBlueOceanLogURLAnnotation     = "openshift.io/jenkins-blueocean-log-url"
	BuildJenkinsBuildURIAnnotation            = "openshift.io/jenkins-build-uri"
	BuildSourceSecretMatchURIAnnotationPrefix = "build.openshift.io/source-secret-match-uri-"
	BuildLabel                                = "openshift.io/build.name"
	BuildRunPolicyLabel                       = "openshift.io/build.start-policy"
	DefaultDockerLabelNamespace               = "io.openshift."
	AllowedUIDs                               = "ALLOWED_UIDS"
	DropCapabilities                          = "DROP_CAPS"
	BuildConfigLabel                          = "openshift.io/build-config.name"
	BuildConfigLabelDeprecated                = "buildconfig"
	BuildConfigPausedAnnotation               = "openshift.io/build-config.paused"
	BuildAcceptedAnnotation                   = "build.openshift.io/accepted"
	BuildStartedEventReason                   = "BuildStarted"
	BuildStartedEventMessage                  = "Build %s/%s is now running"
	BuildCompletedEventReason                 = "BuildCompleted"
	BuildCompletedEventMessage                = "Build %s/%s completed successfully"
	BuildFailedEventReason                    = "BuildFailed"
	BuildFailedEventMessage                   = "Build %s/%s failed"
	BuildCancelledEventReason                 = "BuildCancelled"
	BuildCancelledEventMessage                = "Build %s/%s has been cancelled"
	DefaultSuccessfulBuildsHistoryLimit       = int32(5)
	DefaultFailedBuildsHistoryLimit           = int32(5)
	WebHookSecretKey                          = "WebHookSecretKey"
)

var (
	WhitelistEnvVarNames = []string{"BUILD_LOGLEVEL", "GIT_SSL_NO_VERIFY"}
)

type Build struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   BuildSpec
	Status BuildStatus
}
type BuildSpec struct {
	CommonSpec
	TriggeredBy []BuildTriggerCause
}
type CommonSpec struct {
	ServiceAccount            string
	Source                    BuildSource
	Revision                  *SourceRevision
	Strategy                  BuildStrategy
	Output                    BuildOutput
	Resources                 kapi.ResourceRequirements
	PostCommit                BuildPostCommitSpec
	CompletionDeadlineSeconds *int64
	NodeSelector              map[string]string
}

const (
	BuildTriggerCauseManualMsg    = "Manually triggered"
	BuildTriggerCauseConfigMsg    = "Build configuration change"
	BuildTriggerCauseImageMsg     = "Image change"
	BuildTriggerCauseGithubMsg    = "GitHub WebHook"
	BuildTriggerCauseGenericMsg   = "Generic WebHook"
	BuildTriggerCauseGitLabMsg    = "GitLab WebHook"
	BuildTriggerCauseBitbucketMsg = "Bitbucket WebHook"
)

type BuildTriggerCause struct {
	Message          string
	GenericWebHook   *GenericWebHookCause
	GitHubWebHook    *GitHubWebHookCause
	ImageChangeBuild *ImageChangeCause
	GitLabWebHook    *GitLabWebHookCause
	BitbucketWebHook *BitbucketWebHookCause
}
type GenericWebHookCause struct {
	Revision *SourceRevision
	Secret   string
}
type GitHubWebHookCause struct {
	Revision *SourceRevision
	Secret   string
}
type CommonWebHookCause struct {
	Revision *SourceRevision
	Secret   string
}
type GitLabWebHookCause struct{ CommonWebHookCause }
type BitbucketWebHookCause struct{ CommonWebHookCause }
type ImageChangeCause struct {
	ImageID string
	FromRef *kapi.ObjectReference
}
type BuildStatus struct {
	Phase                      BuildPhase
	Cancelled                  bool
	Reason                     StatusReason
	Message                    string
	StartTimestamp             *metav1.Time
	CompletionTimestamp        *metav1.Time
	Duration                   time.Duration
	OutputDockerImageReference string
	Config                     *kapi.ObjectReference
	Output                     BuildStatusOutput
	Stages                     []StageInfo
	LogSnippet                 string
}
type StageInfo struct {
	Name                 StageName
	StartTime            metav1.Time
	DurationMilliseconds int64
	Steps                []StepInfo
}
type StageName string

const (
	StageFetchInputs StageName = "FetchInputs"
	StagePullImages  StageName = "PullImages"
	StageBuild       StageName = "Build"
	StagePostCommit  StageName = "PostCommit"
	StagePushImage   StageName = "PushImage"
)

type StepInfo struct {
	Name                 StepName
	StartTime            metav1.Time
	DurationMilliseconds int64
}
type StepName string

const (
	StepExecPostCommitHook StepName = "RunPostCommitHook"
	StepFetchGitSource     StepName = "FetchGitSource"
	StepPullBaseImage      StepName = "PullBaseImage"
	StepPullInputImage     StepName = "PullInputImage"
	StepPushImage          StepName = "PushImage"
	StepPushDockerImage    StepName = "PushDockerImage"
	StepDockerBuild        StepName = "DockerBuild"
)

type BuildPhase string

const (
	BuildPhaseNew       BuildPhase = "New"
	BuildPhasePending   BuildPhase = "Pending"
	BuildPhaseRunning   BuildPhase = "Running"
	BuildPhaseComplete  BuildPhase = "Complete"
	BuildPhaseFailed    BuildPhase = "Failed"
	BuildPhaseError     BuildPhase = "Error"
	BuildPhaseCancelled BuildPhase = "Cancelled"
)

type StatusReason string

const (
	StatusReasonError                           StatusReason = "Error"
	StatusReasonCannotCreateBuildPodSpec        StatusReason = "CannotCreateBuildPodSpec"
	StatusReasonCannotCreateBuildPod            StatusReason = "CannotCreateBuildPod"
	StatusReasonInvalidOutputReference          StatusReason = "InvalidOutputReference"
	StatusReasonInvalidImageReference           StatusReason = "InvalidImageReference"
	StatusReasonCancelBuildFailed               StatusReason = "CancelBuildFailed"
	StatusReasonBuildPodDeleted                 StatusReason = "BuildPodDeleted"
	StatusReasonExceededRetryTimeout            StatusReason = "ExceededRetryTimeout"
	StatusReasonMissingPushSecret               StatusReason = "MissingPushSecret"
	StatusReasonPostCommitHookFailed            StatusReason = "PostCommitHookFailed"
	StatusReasonPushImageToRegistryFailed       StatusReason = "PushImageToRegistryFailed"
	StatusReasonPullBuilderImageFailed          StatusReason = "PullBuilderImageFailed"
	StatusReasonFetchSourceFailed               StatusReason = "FetchSourceFailed"
	StatusReasonInvalidContextDirectory         StatusReason = "InvalidContextDirectory"
	StatusReasonCancelledBuild                  StatusReason = "CancelledBuild"
	StatusReasonDockerBuildFailed               StatusReason = "DockerBuildFailed"
	StatusReasonBuildPodExists                  StatusReason = "BuildPodExists"
	StatusReasonNoBuildContainerStatus          StatusReason = "NoBuildContainerStatus"
	StatusReasonFailedContainer                 StatusReason = "FailedContainer"
	StatusReasonUnresolvableEnvironmentVariable StatusReason = "UnresolvableEnvironmentVariable"
	StatusReasonGenericBuildFailed              StatusReason = "GenericBuildFailed"
	StatusReasonOutOfMemoryKilled               StatusReason = "OutOfMemoryKilled"
	StatusReasonCannotRetrieveServiceAccount    StatusReason = "CannotRetrieveServiceAccount"
)
const (
	StatusMessageCannotCreateBuildPodSpec        = "Failed to create pod spec."
	StatusMessageCannotCreateBuildPod            = "Failed creating build pod."
	StatusMessageInvalidOutputRef                = "Output image could not be resolved."
	StatusMessageInvalidImageRef                 = "Referenced image could not be resolved."
	StatusMessageCancelBuildFailed               = "Failed to cancel build."
	StatusMessageBuildPodDeleted                 = "The pod for this build was deleted before the build completed."
	StatusMessageExceededRetryTimeout            = "Build did not complete and retrying timed out."
	StatusMessageMissingPushSecret               = "Missing push secret."
	StatusMessagePostCommitHookFailed            = "Build failed because of post commit hook."
	StatusMessagePushImageToRegistryFailed       = "Failed to push the image to the registry."
	StatusMessagePullBuilderImageFailed          = "Failed pulling builder image."
	StatusMessageFetchSourceFailed               = "Failed to fetch the input source."
	StatusMessageInvalidContextDirectory         = "The supplied context directory does not exist."
	StatusMessageCancelledBuild                  = "The build was cancelled by the user."
	StatusMessageDockerBuildFailed               = "Docker build strategy has failed."
	StatusMessageBuildPodExists                  = "The pod for this build already exists and is older than the build."
	StatusMessageNoBuildContainerStatus          = "The pod for this build has no container statuses indicating success or failure."
	StatusMessageFailedContainer                 = "The pod for this build has at least one container with a non-zero exit status."
	StatusMessageGenericBuildFailed              = "Generic Build failure - check logs for details."
	StatusMessageOutOfMemoryKilled               = "The build pod was killed due to an out of memory condition."
	StatusMessageUnresolvableEnvironmentVariable = "Unable to resolve build environment variable reference."
	StatusMessageCannotRetrieveServiceAccount    = "Unable to look up the service account secrets for this build."
)

type BuildStatusOutput struct{ To *BuildStatusOutputTo }
type BuildStatusOutputTo struct{ ImageDigest string }
type BuildSource struct {
	Binary       *BinaryBuildSource
	Dockerfile   *string
	Git          *GitBuildSource
	Images       []ImageSource
	ContextDir   string
	SourceSecret *kapi.LocalObjectReference
	Secrets      []SecretBuildSource
	ConfigMaps   []ConfigMapBuildSource
}
type ImageSource struct {
	From       kapi.ObjectReference
	As         []string
	Paths      []ImageSourcePath
	PullSecret *kapi.LocalObjectReference
}
type ImageSourcePath struct {
	SourcePath     string
	DestinationDir string
}
type SecretBuildSource struct {
	Secret         kapi.LocalObjectReference
	DestinationDir string
}
type ConfigMapBuildSource struct {
	ConfigMap      kapi.LocalObjectReference
	DestinationDir string
}
type BinaryBuildSource struct{ AsFile string }
type SourceRevision struct{ Git *GitSourceRevision }
type GitSourceRevision struct {
	Commit    string
	Author    SourceControlUser
	Committer SourceControlUser
	Message   string
}
type ProxyConfig struct {
	HTTPProxy  *string
	HTTPSProxy *string
	NoProxy    *string
}
type GitBuildSource struct {
	URI string
	Ref string
	ProxyConfig
}
type SourceControlUser struct {
	Name  string
	Email string
}
type BuildStrategy struct {
	DockerStrategy          *DockerBuildStrategy
	SourceStrategy          *SourceBuildStrategy
	CustomStrategy          *CustomBuildStrategy
	JenkinsPipelineStrategy *JenkinsPipelineBuildStrategy
}
type BuildStrategyType string

const (
	CustomBuildStrategyBaseImageKey = "OPENSHIFT_CUSTOM_BUILD_BASE_IMAGE"
)

type CustomBuildStrategy struct {
	From               kapi.ObjectReference
	PullSecret         *kapi.LocalObjectReference
	Env                []kapi.EnvVar
	ExposeDockerSocket bool
	ForcePull          bool
	Secrets            []SecretSpec
	BuildAPIVersion    string
}
type ImageOptimizationPolicy string

const (
	ImageOptimizationNone              ImageOptimizationPolicy = "None"
	ImageOptimizationSkipLayers        ImageOptimizationPolicy = "SkipLayers"
	ImageOptimizationSkipLayersAndWarn ImageOptimizationPolicy = "SkipLayersAndWarn"
)

type DockerBuildStrategy struct {
	From                    *kapi.ObjectReference
	PullSecret              *kapi.LocalObjectReference
	NoCache                 bool
	Env                     []kapi.EnvVar
	BuildArgs               []kapi.EnvVar
	ForcePull               bool
	DockerfilePath          string
	ImageOptimizationPolicy *ImageOptimizationPolicy
}
type SourceBuildStrategy struct {
	From        kapi.ObjectReference
	PullSecret  *kapi.LocalObjectReference
	Env         []kapi.EnvVar
	Scripts     string
	Incremental *bool
	ForcePull   bool
}
type JenkinsPipelineBuildStrategy struct {
	JenkinsfilePath string
	Jenkinsfile     string
	Env             []kapi.EnvVar
}
type BuildPostCommitSpec struct {
	Command []string
	Args    []string
	Script  string
}
type BuildOutput struct {
	To          *kapi.ObjectReference
	PushSecret  *kapi.LocalObjectReference
	ImageLabels []ImageLabel
}
type ImageLabel struct {
	Name  string
	Value string
}
type BuildConfig struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   BuildConfigSpec
	Status BuildConfigStatus
}
type BuildConfigSpec struct {
	Triggers  []BuildTriggerPolicy
	RunPolicy BuildRunPolicy
	CommonSpec
	SuccessfulBuildsHistoryLimit *int32
	FailedBuildsHistoryLimit     *int32
}
type BuildRunPolicy string

const (
	BuildRunPolicyParallel         BuildRunPolicy = "Parallel"
	BuildRunPolicySerial           BuildRunPolicy = "Serial"
	BuildRunPolicySerialLatestOnly BuildRunPolicy = "SerialLatestOnly"
)

type BuildConfigStatus struct{ LastVersion int64 }
type SecretLocalReference struct{ Name string }
type WebHookTrigger struct {
	Secret          string
	AllowEnv        bool
	SecretReference *SecretLocalReference
}
type ImageChangeTrigger struct {
	LastTriggeredImageID string
	From                 *kapi.ObjectReference
	Paused               bool
}
type BuildTriggerPolicy struct {
	Type             BuildTriggerType
	GitHubWebHook    *WebHookTrigger
	GenericWebHook   *WebHookTrigger
	ImageChange      *ImageChangeTrigger
	GitLabWebHook    *WebHookTrigger
	BitbucketWebHook *WebHookTrigger
}
type BuildTriggerType string

var KnownTriggerTypes = sets.NewString(string(GitHubWebHookBuildTriggerType), string(GenericWebHookBuildTriggerType), string(ImageChangeBuildTriggerType), string(ConfigChangeBuildTriggerType), string(GitLabWebHookBuildTriggerType), string(BitbucketWebHookBuildTriggerType))

const (
	GitHubWebHookBuildTriggerType            BuildTriggerType = "GitHub"
	GitHubWebHookBuildTriggerTypeDeprecated  BuildTriggerType = "github"
	GenericWebHookBuildTriggerType           BuildTriggerType = "Generic"
	GenericWebHookBuildTriggerTypeDeprecated BuildTriggerType = "generic"
	GitLabWebHookBuildTriggerType            BuildTriggerType = "GitLab"
	BitbucketWebHookBuildTriggerType         BuildTriggerType = "Bitbucket"
	ImageChangeBuildTriggerType              BuildTriggerType = "ImageChange"
	ImageChangeBuildTriggerTypeDeprecated    BuildTriggerType = "imageChange"
	ConfigChangeBuildTriggerType             BuildTriggerType = "ConfigChange"
)

type BuildList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Build
}
type BuildConfigList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []BuildConfig
}
type GenericWebHookEvent struct {
	Git                   *GitInfo
	Env                   []kapi.EnvVar
	DockerStrategyOptions *DockerStrategyOptions
}
type GitInfo struct {
	GitBuildSource
	GitSourceRevision
	Refs []GitRefInfo
}
type GitRefInfo struct {
	GitBuildSource
	GitSourceRevision
}
type BuildLog struct{ metav1.TypeMeta }
type DockerStrategyOptions struct {
	BuildArgs []kapi.EnvVar
	NoCache   *bool
}
type SourceStrategyOptions struct{ Incremental *bool }
type BuildRequest struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Revision              *SourceRevision
	TriggeredByImage      *kapi.ObjectReference
	From                  *kapi.ObjectReference
	Binary                *BinaryBuildSource
	LastVersion           *int64
	Env                   []kapi.EnvVar
	TriggeredBy           []BuildTriggerCause
	DockerStrategyOptions *DockerStrategyOptions
	SourceStrategyOptions *SourceStrategyOptions
}
type BinaryBuildRequestOptions struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	AsFile         string
	Commit         string
	Message        string
	AuthorName     string
	AuthorEmail    string
	CommitterName  string
	CommitterEmail string
}
type BuildLogOptions struct {
	metav1.TypeMeta
	Container    string
	Follow       bool
	Previous     bool
	SinceSeconds *int64
	SinceTime    *metav1.Time
	Timestamps   bool
	TailLines    *int64
	LimitBytes   *int64
	NoWait       bool
	Version      *int64
}
type SecretSpec struct {
	SecretSource kapi.LocalObjectReference
	MountPath    string
}
