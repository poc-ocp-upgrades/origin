package portallocator

type PortAllocationOperation struct {
	pa              Interface
	allocated       []int
	releaseDeferred []int
	shouldRollback  bool
	dryRun          bool
}

func StartOperation(pa Interface, dryRun bool) *PortAllocationOperation {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	op := &PortAllocationOperation{}
	op.pa = pa
	op.allocated = []int{}
	op.releaseDeferred = []int{}
	op.shouldRollback = true
	op.dryRun = dryRun
	return op
}
func (op *PortAllocationOperation) Finish() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if op.shouldRollback {
		op.Rollback()
	}
}
func (op *PortAllocationOperation) Rollback() []error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if op.dryRun {
		return nil
	}
	errors := []error{}
	for _, allocated := range op.allocated {
		err := op.pa.Release(allocated)
		if err != nil {
			errors = append(errors, err)
		}
	}
	if len(errors) == 0 {
		return nil
	}
	return errors
}
func (op *PortAllocationOperation) Commit() []error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if op.dryRun {
		return nil
	}
	errors := []error{}
	for _, release := range op.releaseDeferred {
		err := op.pa.Release(release)
		if err != nil {
			errors = append(errors, err)
		}
	}
	op.shouldRollback = false
	if len(errors) == 0 {
		return nil
	}
	return errors
}
func (op *PortAllocationOperation) Allocate(port int) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if op.dryRun {
		if op.pa.Has(port) {
			return ErrAllocated
		}
		for _, a := range op.allocated {
			if port == a {
				return ErrAllocated
			}
		}
		op.allocated = append(op.allocated, port)
		return nil
	}
	err := op.pa.Allocate(port)
	if err == nil {
		op.allocated = append(op.allocated, port)
	}
	return err
}
func (op *PortAllocationOperation) AllocateNext() (int, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if op.dryRun {
		var lastPort int
		for _, allocatedPort := range op.allocated {
			if allocatedPort > lastPort {
				lastPort = allocatedPort
			}
		}
		if len(op.allocated) == 0 {
			lastPort = 32768
		}
		for port := lastPort + 1; port < 100; port++ {
			err := op.Allocate(port)
			if err == nil {
				return port, nil
			}
		}
		op.allocated = append(op.allocated, lastPort+1)
		return lastPort + 1, nil
	}
	port, err := op.pa.AllocateNext()
	if err == nil {
		op.allocated = append(op.allocated, port)
	}
	return port, err
}
func (op *PortAllocationOperation) ReleaseDeferred(port int) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	op.releaseDeferred = append(op.releaseDeferred, port)
}
