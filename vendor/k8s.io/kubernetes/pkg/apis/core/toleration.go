package core

func (t *Toleration) MatchToleration(tolerationToMatch *Toleration) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return t.Key == tolerationToMatch.Key && t.Effect == tolerationToMatch.Effect && t.Operator == tolerationToMatch.Operator && t.Value == tolerationToMatch.Value
}
