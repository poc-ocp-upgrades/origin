package graphview

import (
	"sort"
	"github.com/gonum/graph"
	buildedges "github.com/openshift/origin/pkg/oc/lib/graph/buildgraph"
	buildgraph "github.com/openshift/origin/pkg/oc/lib/graph/buildgraph/nodes"
	osgraph "github.com/openshift/origin/pkg/oc/lib/graph/genericgraph"
	imageedges "github.com/openshift/origin/pkg/oc/lib/graph/imagegraph"
	imagegraph "github.com/openshift/origin/pkg/oc/lib/graph/imagegraph/nodes"
)

type ImagePipeline struct {
	Image			ImageTagLocation
	DestinationResolved	bool
	ScheduledImport		bool
	Build			*buildgraph.BuildConfigNode
	LastSuccessfulBuild	*buildgraph.BuildNode
	LastUnsuccessfulBuild	*buildgraph.BuildNode
	ActiveBuilds		[]*buildgraph.BuildNode
	BaseImage		ImageTagLocation
	BaseBuilds		[]string
	Source			SourceLocation
}
type ImageTagLocation interface {
	ID() int
	ImageSpec() string
	ImageTag() string
}
type SourceLocation interface{ ID() int }

func AllImagePipelinesFromBuildConfig(g osgraph.Graph, excludeNodeIDs IntSet) ([]ImagePipeline, IntSet) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	covered := IntSet{}
	pipelines := []ImagePipeline{}
	for _, uncastNode := range g.NodesByKind(buildgraph.BuildConfigNodeKind) {
		if excludeNodeIDs.Has(uncastNode.ID()) {
			continue
		}
		pipeline, covers := NewImagePipelineFromBuildConfigNode(g, uncastNode.(*buildgraph.BuildConfigNode))
		covered.Insert(covers.List()...)
		pipelines = append(pipelines, pipeline)
	}
	sort.Sort(SortedImagePipelines(pipelines))
	outputImageToBCMap := make(map[string][]string)
	for _, pipeline := range pipelines {
		if pipeline.Image != nil {
			bcs, ok := outputImageToBCMap[pipeline.Image.ImageSpec()]
			if !ok {
				bcs = []string{}
			}
			bcs = append(bcs, pipeline.Build.BuildConfig.Name)
			outputImageToBCMap[pipeline.Image.ImageSpec()] = bcs
		}
	}
	if len(outputImageToBCMap) > 0 {
		for i, pipeline := range pipelines {
			if pipeline.BaseImage != nil {
				baseBCs, ok := outputImageToBCMap[pipeline.BaseImage.ImageSpec()]
				if ok && len(baseBCs) > 0 {
					pipelines[i].BaseBuilds = baseBCs
				}
			}
		}
	}
	return pipelines, covered
}
func NewImagePipelineFromBuildConfigNode(g osgraph.Graph, bcNode *buildgraph.BuildConfigNode) (ImagePipeline, IntSet) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	covered := IntSet{}
	covered.Insert(bcNode.ID())
	flow := ImagePipeline{}
	base, src, coveredInputs, scheduled, _ := findBuildInputs(g, bcNode)
	covered.Insert(coveredInputs.List()...)
	flow.BaseImage = base
	flow.Source = src
	flow.Build = bcNode
	flow.ScheduledImport = scheduled
	flow.LastSuccessfulBuild, flow.LastUnsuccessfulBuild, flow.ActiveBuilds = buildedges.RelevantBuilds(g, flow.Build)
	flow.Image = findBuildOutput(g, bcNode)
	for _, buildOutputNode := range g.SuccessorNodesByEdgeKind(bcNode, buildedges.BuildOutputEdgeKind) {
		for _, input := range g.SuccessorNodesByEdgeKind(buildOutputNode, imageedges.ReferencedImageStreamGraphEdgeKind) {
			imageStreamNode := input.(*imagegraph.ImageStreamNode)
			flow.DestinationResolved = (len(imageStreamNode.Status.DockerImageRepository) != 0)
		}
		for _, input := range g.SuccessorNodesByEdgeKind(buildOutputNode, imageedges.ReferencedImageStreamImageGraphEdgeKind) {
			imageStreamNode := input.(*imagegraph.ImageStreamNode)
			flow.DestinationResolved = (len(imageStreamNode.Status.DockerImageRepository) != 0)
		}
	}
	return flow, covered
}
func NewImagePipelineFromImageTagLocation(g osgraph.Graph, node graph.Node, imageTagLocation ImageTagLocation) (ImagePipeline, IntSet) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	covered := IntSet{}
	covered.Insert(node.ID())
	flow := ImagePipeline{}
	flow.Image = imageTagLocation
	for _, input := range g.PredecessorNodesByEdgeKind(node, buildedges.BuildOutputEdgeKind) {
		covered.Insert(input.ID())
		build := input.(*buildgraph.BuildConfigNode)
		if flow.Build != nil {
		}
		if build.BuildConfig == nil {
			break
		}
		base, src, coveredInputs, scheduled, _ := findBuildInputs(g, build)
		covered.Insert(coveredInputs.List()...)
		flow.BaseImage = base
		flow.Source = src
		flow.Build = build
		flow.ScheduledImport = scheduled
		flow.LastSuccessfulBuild, flow.LastUnsuccessfulBuild, flow.ActiveBuilds = buildedges.RelevantBuilds(g, flow.Build)
	}
	for _, input := range g.SuccessorNodesByEdgeKind(node, imageedges.ReferencedImageStreamGraphEdgeKind) {
		covered.Insert(input.ID())
		imageStreamNode := input.(*imagegraph.ImageStreamNode)
		flow.DestinationResolved = (len(imageStreamNode.Status.DockerImageRepository) != 0)
	}
	for _, input := range g.SuccessorNodesByEdgeKind(node, imageedges.ReferencedImageStreamImageGraphEdgeKind) {
		covered.Insert(input.ID())
		imageStreamNode := input.(*imagegraph.ImageStreamNode)
		flow.DestinationResolved = (len(imageStreamNode.Status.DockerImageRepository) != 0)
	}
	return flow, covered
}
func findBuildInputs(g osgraph.Graph, bcNode *buildgraph.BuildConfigNode) (base ImageTagLocation, source SourceLocation, covered IntSet, scheduled bool, err error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	covered = IntSet{}
	for _, input := range g.PredecessorNodesByEdgeKind(bcNode, buildedges.BuildInputEdgeKind) {
		if source != nil {
		}
		covered.Insert(input.ID())
		source = input.(SourceLocation)
	}
	for _, input := range g.PredecessorNodesByEdgeKind(bcNode, buildedges.BuildInputImageEdgeKind) {
		if base != nil {
		}
		covered.Insert(input.ID())
		base = input.(ImageTagLocation)
		scheduled = imageStreamTagScheduled(g, input, base)
	}
	return
}
func findBuildOutput(g osgraph.Graph, bcNode *buildgraph.BuildConfigNode) (result ImageTagLocation) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, output := range g.SuccessorNodesByEdgeKind(bcNode, buildedges.BuildOutputEdgeKind) {
		result = output.(ImageTagLocation)
		return
	}
	return
}
func imageStreamTagScheduled(g osgraph.Graph, input graph.Node, base ImageTagLocation) (scheduled bool) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, uncastImageStreamNode := range g.SuccessorNodesByEdgeKind(input, imageedges.ReferencedImageStreamGraphEdgeKind) {
		imageStreamNode := uncastImageStreamNode.(*imagegraph.ImageStreamNode)
		if imageStreamNode.ImageStream != nil {
			for _, tag := range imageStreamNode.ImageStream.Spec.Tags {
				if tag.Name == base.ImageTag() {
					scheduled = tag.ImportPolicy.Scheduled
					return
				}
			}
		}
	}
	return
}

type SortedImagePipelines []ImagePipeline

func (m SortedImagePipelines) Len() int {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(m)
}
func (m SortedImagePipelines) Swap(i, j int) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	m[i], m[j] = m[j], m[i]
}
func (m SortedImagePipelines) Less(i, j int) bool {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return CompareImagePipeline(&m[i], &m[j])
}
func CompareImagePipeline(a, b *ImagePipeline) bool {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch {
	case a.Build != nil && b.Build != nil && a.Build.BuildConfig != nil && b.Build.BuildConfig != nil:
		return CompareObjectMeta(&a.Build.BuildConfig.ObjectMeta, &b.Build.BuildConfig.ObjectMeta)
	case a.Build != nil && a.Build.BuildConfig != nil:
		return true
	case b.Build != nil && b.Build.BuildConfig != nil:
		return false
	}
	if a.Image == nil || b.Image == nil {
		return true
	}
	return a.Image.ImageSpec() < b.Image.ImageSpec()
}
