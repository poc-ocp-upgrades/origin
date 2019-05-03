package allocator

type Interface interface {
 Allocate(int) (bool, error)
 AllocateNext() (int, bool, error)
 Release(int) error
 ForEach(func(int))
 Has(int) bool
 Free() int
}
type Snapshottable interface {
 Interface
 Snapshot() (string, []byte)
 Restore(string, []byte) error
}
type AllocatorFactory func(max int, rangeSpec string) Interface
