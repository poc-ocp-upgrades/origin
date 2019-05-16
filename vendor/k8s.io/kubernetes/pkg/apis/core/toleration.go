package core

func (t *Toleration) MatchToleration(tolerationToMatch *Toleration) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return t.Key == tolerationToMatch.Key && t.Effect == tolerationToMatch.Effect && t.Operator == tolerationToMatch.Operator && t.Value == tolerationToMatch.Value
}
