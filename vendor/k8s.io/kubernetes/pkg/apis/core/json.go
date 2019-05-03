package core

import "encoding/json"

var _ = json.Marshaler(&AvoidPods{})
var _ = json.Unmarshaler(&AvoidPods{})

func (AvoidPods) MarshalJSON() ([]byte, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 panic("do not marshal internal struct")
}
func (*AvoidPods) UnmarshalJSON([]byte) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 panic("do not unmarshal to internal struct")
}
