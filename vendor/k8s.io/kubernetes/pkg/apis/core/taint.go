package core

import "fmt"

func (t *Taint) MatchTaint(taintToMatch Taint) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return t.Key == taintToMatch.Key && t.Effect == taintToMatch.Effect
}
func (t *Taint) ToString() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(t.Value) == 0 {
  return fmt.Sprintf("%v:%v", t.Key, t.Effect)
 }
 return fmt.Sprintf("%v=%v:%v", t.Key, t.Value, t.Effect)
}
