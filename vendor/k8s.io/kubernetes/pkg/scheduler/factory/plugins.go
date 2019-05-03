package factory

import (
 "fmt"
 "regexp"
 "sort"
 "strings"
 "sync"
 "k8s.io/apimachinery/pkg/util/sets"
 "k8s.io/kubernetes/pkg/scheduler/algorithm"
 "k8s.io/kubernetes/pkg/scheduler/algorithm/predicates"
 "k8s.io/kubernetes/pkg/scheduler/algorithm/priorities"
 schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
 "k8s.io/kubernetes/pkg/scheduler/volumebinder"
 "k8s.io/klog"
)

type PluginFactoryArgs struct {
 PodLister                      algorithm.PodLister
 ServiceLister                  algorithm.ServiceLister
 ControllerLister               algorithm.ControllerLister
 ReplicaSetLister               algorithm.ReplicaSetLister
 StatefulSetLister              algorithm.StatefulSetLister
 NodeLister                     algorithm.NodeLister
 PDBLister                      algorithm.PDBLister
 NodeInfo                       predicates.NodeInfo
 PVInfo                         predicates.PersistentVolumeInfo
 PVCInfo                        predicates.PersistentVolumeClaimInfo
 StorageClassInfo               predicates.StorageClassInfo
 VolumeBinder                   *volumebinder.VolumeBinder
 HardPodAffinitySymmetricWeight int32
}
type PriorityMetadataProducerFactory func(PluginFactoryArgs) algorithm.PriorityMetadataProducer
type PredicateMetadataProducerFactory func(PluginFactoryArgs) algorithm.PredicateMetadataProducer
type FitPredicateFactory func(PluginFactoryArgs) algorithm.FitPredicate
type PriorityFunctionFactory func(PluginFactoryArgs) algorithm.PriorityFunction
type PriorityFunctionFactory2 func(PluginFactoryArgs) (algorithm.PriorityMapFunction, algorithm.PriorityReduceFunction)
type PriorityConfigFactory struct {
 Function          PriorityFunctionFactory
 MapReduceFunction PriorityFunctionFactory2
 Weight            int
}

var (
 schedulerFactoryMutex     sync.Mutex
 fitPredicateMap           = make(map[string]FitPredicateFactory)
 mandatoryFitPredicates    = sets.NewString()
 priorityFunctionMap       = make(map[string]PriorityConfigFactory)
 algorithmProviderMap      = make(map[string]AlgorithmProviderConfig)
 priorityMetadataProducer  PriorityMetadataProducerFactory
 predicateMetadataProducer PredicateMetadataProducerFactory
)

const (
 DefaultProvider = "DefaultProvider"
)

type AlgorithmProviderConfig struct {
 FitPredicateKeys     sets.String
 PriorityFunctionKeys sets.String
}

func RegisterFitPredicate(name string, predicate algorithm.FitPredicate) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return RegisterFitPredicateFactory(name, func(PluginFactoryArgs) algorithm.FitPredicate {
  return predicate
 })
}
func RemoveFitPredicate(name string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 schedulerFactoryMutex.Lock()
 defer schedulerFactoryMutex.Unlock()
 validateAlgorithmNameOrDie(name)
 delete(fitPredicateMap, name)
 mandatoryFitPredicates.Delete(name)
}
func RemovePredicateKeyFromAlgoProvider(providerName, key string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 schedulerFactoryMutex.Lock()
 defer schedulerFactoryMutex.Unlock()
 validateAlgorithmNameOrDie(providerName)
 provider, ok := algorithmProviderMap[providerName]
 if !ok {
  return fmt.Errorf("plugin %v has not been registered", providerName)
 }
 provider.FitPredicateKeys.Delete(key)
 return nil
}
func RemovePredicateKeyFromAlgorithmProviderMap(key string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 schedulerFactoryMutex.Lock()
 defer schedulerFactoryMutex.Unlock()
 for _, provider := range algorithmProviderMap {
  provider.FitPredicateKeys.Delete(key)
 }
 return
}
func InsertPredicateKeyToAlgoProvider(providerName, key string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 schedulerFactoryMutex.Lock()
 defer schedulerFactoryMutex.Unlock()
 validateAlgorithmNameOrDie(providerName)
 provider, ok := algorithmProviderMap[providerName]
 if !ok {
  return fmt.Errorf("plugin %v has not been registered", providerName)
 }
 provider.FitPredicateKeys.Insert(key)
 return nil
}
func InsertPredicateKeyToAlgorithmProviderMap(key string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 schedulerFactoryMutex.Lock()
 defer schedulerFactoryMutex.Unlock()
 for _, provider := range algorithmProviderMap {
  provider.FitPredicateKeys.Insert(key)
 }
 return
}
func InsertPriorityKeyToAlgorithmProviderMap(key string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 schedulerFactoryMutex.Lock()
 defer schedulerFactoryMutex.Unlock()
 for _, provider := range algorithmProviderMap {
  provider.PriorityFunctionKeys.Insert(key)
 }
 return
}
func RegisterMandatoryFitPredicate(name string, predicate algorithm.FitPredicate) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 schedulerFactoryMutex.Lock()
 defer schedulerFactoryMutex.Unlock()
 validateAlgorithmNameOrDie(name)
 fitPredicateMap[name] = func(PluginFactoryArgs) algorithm.FitPredicate {
  return predicate
 }
 mandatoryFitPredicates.Insert(name)
 return name
}
func RegisterFitPredicateFactory(name string, predicateFactory FitPredicateFactory) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 schedulerFactoryMutex.Lock()
 defer schedulerFactoryMutex.Unlock()
 validateAlgorithmNameOrDie(name)
 fitPredicateMap[name] = predicateFactory
 return name
}
func RegisterCustomFitPredicate(policy schedulerapi.PredicatePolicy) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var predicateFactory FitPredicateFactory
 var ok bool
 validatePredicateOrDie(policy)
 if policy.Argument != nil {
  if policy.Argument.ServiceAffinity != nil {
   predicateFactory = func(args PluginFactoryArgs) algorithm.FitPredicate {
    predicate, precomputationFunction := predicates.NewServiceAffinityPredicate(args.PodLister, args.ServiceLister, args.NodeInfo, policy.Argument.ServiceAffinity.Labels)
    predicates.RegisterPredicateMetadataProducer(policy.Name, precomputationFunction)
    return predicate
   }
  } else if policy.Argument.LabelsPresence != nil {
   predicateFactory = func(args PluginFactoryArgs) algorithm.FitPredicate {
    return predicates.NewNodeLabelPredicate(policy.Argument.LabelsPresence.Labels, policy.Argument.LabelsPresence.Presence)
   }
  }
 } else if predicateFactory, ok = fitPredicateMap[policy.Name]; ok {
  klog.V(2).Infof("Predicate type %s already registered, reusing.", policy.Name)
  return policy.Name
 }
 if predicateFactory == nil {
  klog.Fatalf("Invalid configuration: Predicate type not found for %s", policy.Name)
 }
 return RegisterFitPredicateFactory(policy.Name, predicateFactory)
}
func IsFitPredicateRegistered(name string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 schedulerFactoryMutex.Lock()
 defer schedulerFactoryMutex.Unlock()
 _, ok := fitPredicateMap[name]
 return ok
}
func RegisterPriorityMetadataProducerFactory(factory PriorityMetadataProducerFactory) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 schedulerFactoryMutex.Lock()
 defer schedulerFactoryMutex.Unlock()
 priorityMetadataProducer = factory
}
func RegisterPredicateMetadataProducerFactory(factory PredicateMetadataProducerFactory) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 schedulerFactoryMutex.Lock()
 defer schedulerFactoryMutex.Unlock()
 predicateMetadataProducer = factory
}
func RegisterPriorityFunction(name string, function algorithm.PriorityFunction, weight int) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return RegisterPriorityConfigFactory(name, PriorityConfigFactory{Function: func(PluginFactoryArgs) algorithm.PriorityFunction {
  return function
 }, Weight: weight})
}
func RegisterPriorityFunction2(name string, mapFunction algorithm.PriorityMapFunction, reduceFunction algorithm.PriorityReduceFunction, weight int) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return RegisterPriorityConfigFactory(name, PriorityConfigFactory{MapReduceFunction: func(PluginFactoryArgs) (algorithm.PriorityMapFunction, algorithm.PriorityReduceFunction) {
  return mapFunction, reduceFunction
 }, Weight: weight})
}
func RegisterPriorityConfigFactory(name string, pcf PriorityConfigFactory) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 schedulerFactoryMutex.Lock()
 defer schedulerFactoryMutex.Unlock()
 validateAlgorithmNameOrDie(name)
 priorityFunctionMap[name] = pcf
 return name
}
func RegisterCustomPriorityFunction(policy schedulerapi.PriorityPolicy) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var pcf *PriorityConfigFactory
 validatePriorityOrDie(policy)
 if policy.Argument != nil {
  if policy.Argument.ServiceAntiAffinity != nil {
   pcf = &PriorityConfigFactory{MapReduceFunction: func(args PluginFactoryArgs) (algorithm.PriorityMapFunction, algorithm.PriorityReduceFunction) {
    return priorities.NewServiceAntiAffinityPriority(args.PodLister, args.ServiceLister, policy.Argument.ServiceAntiAffinity.Label)
   }, Weight: policy.Weight}
  } else if policy.Argument.LabelPreference != nil {
   pcf = &PriorityConfigFactory{MapReduceFunction: func(args PluginFactoryArgs) (algorithm.PriorityMapFunction, algorithm.PriorityReduceFunction) {
    return priorities.NewNodeLabelPriority(policy.Argument.LabelPreference.Label, policy.Argument.LabelPreference.Presence)
   }, Weight: policy.Weight}
  } else if policy.Argument.RequestedToCapacityRatioArguments != nil {
   pcf = &PriorityConfigFactory{MapReduceFunction: func(args PluginFactoryArgs) (algorithm.PriorityMapFunction, algorithm.PriorityReduceFunction) {
    scoringFunctionShape := buildScoringFunctionShapeFromRequestedToCapacityRatioArguments(policy.Argument.RequestedToCapacityRatioArguments)
    p := priorities.RequestedToCapacityRatioResourceAllocationPriority(scoringFunctionShape)
    return p.PriorityMap, nil
   }, Weight: policy.Weight}
  }
 } else if existingPcf, ok := priorityFunctionMap[policy.Name]; ok {
  klog.V(2).Infof("Priority type %s already registered, reusing.", policy.Name)
  pcf = &PriorityConfigFactory{Function: existingPcf.Function, MapReduceFunction: existingPcf.MapReduceFunction, Weight: policy.Weight}
 }
 if pcf == nil {
  klog.Fatalf("Invalid configuration: Priority type not found for %s", policy.Name)
 }
 return RegisterPriorityConfigFactory(policy.Name, *pcf)
}
func buildScoringFunctionShapeFromRequestedToCapacityRatioArguments(arguments *schedulerapi.RequestedToCapacityRatioArguments) priorities.FunctionShape {
 _logClusterCodePath()
 defer _logClusterCodePath()
 n := len(arguments.UtilizationShape)
 points := make([]priorities.FunctionShapePoint, 0, n)
 for _, point := range arguments.UtilizationShape {
  points = append(points, priorities.FunctionShapePoint{Utilization: int64(point.Utilization), Score: int64(point.Score)})
 }
 shape, err := priorities.NewFunctionShape(points)
 if err != nil {
  klog.Fatalf("invalid RequestedToCapacityRatioPriority arguments: %s", err.Error())
 }
 return shape
}
func IsPriorityFunctionRegistered(name string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 schedulerFactoryMutex.Lock()
 defer schedulerFactoryMutex.Unlock()
 _, ok := priorityFunctionMap[name]
 return ok
}
func RegisterAlgorithmProvider(name string, predicateKeys, priorityKeys sets.String) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 schedulerFactoryMutex.Lock()
 defer schedulerFactoryMutex.Unlock()
 validateAlgorithmNameOrDie(name)
 algorithmProviderMap[name] = AlgorithmProviderConfig{FitPredicateKeys: predicateKeys, PriorityFunctionKeys: priorityKeys}
 return name
}
func GetAlgorithmProvider(name string) (*AlgorithmProviderConfig, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 schedulerFactoryMutex.Lock()
 defer schedulerFactoryMutex.Unlock()
 provider, ok := algorithmProviderMap[name]
 if !ok {
  return nil, fmt.Errorf("plugin %q has not been registered", name)
 }
 return &provider, nil
}
func getFitPredicateFunctions(names sets.String, args PluginFactoryArgs) (map[string]algorithm.FitPredicate, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 schedulerFactoryMutex.Lock()
 defer schedulerFactoryMutex.Unlock()
 predicates := map[string]algorithm.FitPredicate{}
 for _, name := range names.List() {
  factory, ok := fitPredicateMap[name]
  if !ok {
   return nil, fmt.Errorf("Invalid predicate name %q specified - no corresponding function found", name)
  }
  predicates[name] = factory(args)
 }
 for name := range mandatoryFitPredicates {
  if factory, found := fitPredicateMap[name]; found {
   predicates[name] = factory(args)
  }
 }
 return predicates, nil
}
func getPriorityMetadataProducer(args PluginFactoryArgs) (algorithm.PriorityMetadataProducer, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 schedulerFactoryMutex.Lock()
 defer schedulerFactoryMutex.Unlock()
 if priorityMetadataProducer == nil {
  return algorithm.EmptyPriorityMetadataProducer, nil
 }
 return priorityMetadataProducer(args), nil
}
func getPredicateMetadataProducer(args PluginFactoryArgs) (algorithm.PredicateMetadataProducer, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 schedulerFactoryMutex.Lock()
 defer schedulerFactoryMutex.Unlock()
 if predicateMetadataProducer == nil {
  return algorithm.EmptyPredicateMetadataProducer, nil
 }
 return predicateMetadataProducer(args), nil
}
func getPriorityFunctionConfigs(names sets.String, args PluginFactoryArgs) ([]algorithm.PriorityConfig, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 schedulerFactoryMutex.Lock()
 defer schedulerFactoryMutex.Unlock()
 configs := []algorithm.PriorityConfig{}
 for _, name := range names.List() {
  factory, ok := priorityFunctionMap[name]
  if !ok {
   return nil, fmt.Errorf("Invalid priority name %s specified - no corresponding function found", name)
  }
  if factory.Function != nil {
   configs = append(configs, algorithm.PriorityConfig{Name: name, Function: factory.Function(args), Weight: factory.Weight})
  } else {
   mapFunction, reduceFunction := factory.MapReduceFunction(args)
   configs = append(configs, algorithm.PriorityConfig{Name: name, Map: mapFunction, Reduce: reduceFunction, Weight: factory.Weight})
  }
 }
 if err := validateSelectedConfigs(configs); err != nil {
  return nil, err
 }
 return configs, nil
}
func validateSelectedConfigs(configs []algorithm.PriorityConfig) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var totalPriority int
 for _, config := range configs {
  if config.Weight*schedulerapi.MaxPriority > schedulerapi.MaxTotalPriority-totalPriority {
   return fmt.Errorf("Total priority of priority functions has overflown")
  }
  totalPriority += config.Weight * schedulerapi.MaxPriority
 }
 return nil
}

var validName = regexp.MustCompile("^[a-zA-Z0-9]([-a-zA-Z0-9]*[a-zA-Z0-9])$")

func validateAlgorithmNameOrDie(name string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !validName.MatchString(name) {
  klog.Fatalf("Algorithm name %v does not match the name validation regexp \"%v\".", name, validName)
 }
}
func validatePredicateOrDie(predicate schedulerapi.PredicatePolicy) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if predicate.Argument != nil {
  numArgs := 0
  if predicate.Argument.ServiceAffinity != nil {
   numArgs++
  }
  if predicate.Argument.LabelsPresence != nil {
   numArgs++
  }
  if numArgs != 1 {
   klog.Fatalf("Exactly 1 predicate argument is required, numArgs: %v, Predicate: %s", numArgs, predicate.Name)
  }
 }
}
func validatePriorityOrDie(priority schedulerapi.PriorityPolicy) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if priority.Argument != nil {
  numArgs := 0
  if priority.Argument.ServiceAntiAffinity != nil {
   numArgs++
  }
  if priority.Argument.LabelPreference != nil {
   numArgs++
  }
  if priority.Argument.RequestedToCapacityRatioArguments != nil {
   numArgs++
  }
  if numArgs != 1 {
   klog.Fatalf("Exactly 1 priority argument is required, numArgs: %v, Priority: %s", numArgs, priority.Name)
  }
 }
}
func ListRegisteredFitPredicates() []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 schedulerFactoryMutex.Lock()
 defer schedulerFactoryMutex.Unlock()
 names := []string{}
 for name := range fitPredicateMap {
  names = append(names, name)
 }
 return names
}
func ListRegisteredPriorityFunctions() []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 schedulerFactoryMutex.Lock()
 defer schedulerFactoryMutex.Unlock()
 names := []string{}
 for name := range priorityFunctionMap {
  names = append(names, name)
 }
 return names
}
func ListAlgorithmProviders() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var availableAlgorithmProviders []string
 for name := range algorithmProviderMap {
  availableAlgorithmProviders = append(availableAlgorithmProviders, name)
 }
 sort.Strings(availableAlgorithmProviders)
 return strings.Join(availableAlgorithmProviders, " | ")
}
