package core

import "encoding/json"

var _ = json.Marshaler(&AvoidPods{})
var _ = json.Unmarshaler(&AvoidPods{})

func (AvoidPods) MarshalJSON() ([]byte, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	panic("do not marshal internal struct")
}
func (*AvoidPods) UnmarshalJSON([]byte) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	panic("do not unmarshal to internal struct")
}
