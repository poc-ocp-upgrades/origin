package registrytest

import (
 "fmt"
 "k8s.io/apiserver/pkg/registry/generic/registry"
 "k8s.io/apiserver/pkg/registry/rest"
 "k8s.io/kubernetes/pkg/util/slice"
)

func ValidateStorageStrategies(storageMap map[string]rest.Storage, exceptions StrategyExceptions) []error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 errs := []error{}
 hasExportExceptionsSeen := []string{}
 for k, storage := range storageMap {
  switch t := storage.(type) {
  case registry.GenericStore:
   if t.GetCreateStrategy() == nil {
    errs = append(errs, fmt.Errorf("store for type [%v] does not have a CreateStrategy", k))
   }
   if t.GetUpdateStrategy() == nil {
    errs = append(errs, fmt.Errorf("store for type [%v] does not have an UpdateStrategy", k))
   }
   if t.GetDeleteStrategy() == nil {
    errs = append(errs, fmt.Errorf("store for type [%v] does not have a DeleteStrategy", k))
   }
   if slice.ContainsString(exceptions.HasExportStrategy, k, nil) {
    hasExportExceptionsSeen = append(hasExportExceptionsSeen, k)
    if t.GetExportStrategy() == nil {
     errs = append(errs, fmt.Errorf("store for type [%v] does not have an ExportStrategy", k))
    }
   } else {
    if t.GetExportStrategy() != nil {
     errs = append(errs, fmt.Errorf("store for type [%v] has an unexpected ExportStrategy", k))
    }
   }
  }
 }
 for _, expKey := range exceptions.HasExportStrategy {
  if !slice.ContainsString(hasExportExceptionsSeen, expKey, nil) {
   errs = append(errs, fmt.Errorf("no generic store seen for expected ExportStrategy: %v", expKey))
  }
 }
 return errs
}

type StrategyExceptions struct{ HasExportStrategy []string }
