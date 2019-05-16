package rankedset

import (
	goformat "fmt"
	"github.com/google/btree"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type Item interface {
	Key() string
	Rank() int64
}
type RankedSet struct {
	rank *btree.BTree
	set  map[string]*treeItem
}
type StringItem string

func (s StringItem) Key() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return string(s)
}
func (s StringItem) Rank() int64 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return 0
}
func New() *RankedSet {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &RankedSet{rank: btree.New(32), set: make(map[string]*treeItem)}
}
func (s *RankedSet) Insert(item Item) Item {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	old := s.Delete(item)
	key := item.Key()
	value := &treeItem{item: item}
	s.rank.ReplaceOrInsert(value)
	s.set[key] = value
	return old
}
func (s *RankedSet) Delete(item Item) Item {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key := item.Key()
	value, ok := s.set[key]
	if !ok {
		return nil
	}
	s.rank.Delete(value)
	delete(s.set, key)
	return value.item
}
func (s *RankedSet) Min() Item {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if min := s.rank.Min(); min != nil {
		return min.(*treeItem).item
	}
	return nil
}
func (s *RankedSet) Max() Item {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if max := s.rank.Max(); max != nil {
		return max.(*treeItem).item
	}
	return nil
}
func (s *RankedSet) Len() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return len(s.set)
}
func (s *RankedSet) Get(item Item) Item {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if value, ok := s.set[item.Key()]; ok {
		return value.item
	}
	return nil
}
func (s *RankedSet) Has(item Item) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_, ok := s.set[item.Key()]
	return ok
}
func (s *RankedSet) List(delete bool) []Item {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return s.ascend(func(item Item) bool {
		return true
	}, delete)
}
func (s *RankedSet) LessThan(rank int64, delete bool) []Item {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return s.ascend(func(item Item) bool {
		return item.Rank() < rank
	}, delete)
}

type setItemIterator func(item Item) bool

func (s *RankedSet) ascend(iterator setItemIterator, delete bool) []Item {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var items []Item
	s.rank.Ascend(func(i btree.Item) bool {
		item := i.(*treeItem).item
		if !iterator(item) {
			return false
		}
		items = append(items, item)
		return true
	})
	if delete {
		for _, item := range items {
			s.Delete(item)
		}
	}
	return items
}

var _ btree.Item = &treeItem{}

type treeItem struct{ item Item }

func (i *treeItem) Less(than btree.Item) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	other := than.(*treeItem).item
	selfRank := i.item.Rank()
	otherRank := other.Rank()
	if selfRank == otherRank {
		return i.item.Key() < other.Key()
	}
	return selfRank < otherRank
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
