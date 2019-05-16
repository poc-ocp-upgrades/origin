package node

type intSet struct {
	currentGeneration byte
	members           map[int]byte
}

func newIntSet() *intSet {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &intSet{members: map[int]byte{}}
}
func (s *intSet) has(i int) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if s == nil {
		return false
	}
	_, present := s.members[i]
	return present
}
func (s *intSet) startNewGeneration() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	s.currentGeneration++
}
func (s *intSet) mark(i int) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	s.members[i] = s.currentGeneration
}
func (s *intSet) sweep() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for k, v := range s.members {
		if v != s.currentGeneration {
			delete(s.members, k)
		}
	}
}
