package rangeallocation

import (
 api "k8s.io/kubernetes/pkg/apis/core"
)

type RangeRegistry interface {
 Get() (*api.RangeAllocation, error)
 CreateOrUpdate(*api.RangeAllocation) error
}
