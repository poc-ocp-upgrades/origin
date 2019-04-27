package migrate

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"k8s.io/klog"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/cli-runtime/pkg/genericclioptions/resource"
)

func TestResourceVisitor_Visit(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var level klog.Level
	origVerbosity := level.Get()
	level.Set("1")
	defer func() {
		level.Set(fmt.Sprintf("%d", origVerbosity))
	}()
	type fields struct {
		Out		mapWriter
		Builder		testBuilder
		SaveFn		*countSaveFn
		PrintFn		MigrateActionFunc
		FilterFn	MigrateFilterFunc
		DryRun		bool
		Workers		int
	}
	type args struct{ fn MigrateVisitFunc }
	tests := []struct {
		name	string
		fields	fields
		args	args
		wantErr	bool
	}{{name: "migrate storage race detection", fields: fields{Out: make(mapWriter), Builder: testBuilder(5000), SaveFn: new(countSaveFn), PrintFn: nil, FilterFn: nil, DryRun: false, Workers: 32 * runtime.NumCPU()}, args: args{fn: AlwaysRequiresMigration}, wantErr: false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &ResourceVisitor{Out: tt.fields.Out, Builder: tt.fields.Builder, SaveFn: tt.fields.SaveFn.save, PrintFn: tt.fields.PrintFn, FilterFn: tt.fields.FilterFn, DryRun: tt.fields.DryRun, Workers: tt.fields.Workers}
			expectedInfos := int(tt.fields.Builder)
			tt.fields.SaveFn.w.Add(expectedInfos)
			if err := o.Visit(tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("ResourceVisitor.Visit() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.fields.SaveFn.w.Wait()
			writes := len(tt.fields.Out) - 1
			saves := tt.fields.SaveFn.n
			if expectedInfos != writes || expectedInfos != saves {
				t.Errorf("ResourceVisitor.Visit() incorrect counts seen, expectedInfos=%d writes=%d saves=%d out=%v", expectedInfos, writes, saves, tt.fields.Out)
			}
		})
	}
}

type mapWriter map[int]string

func (m mapWriter) Write(p []byte) (n int, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(m)
	m[l] = string(p)
	return len(p), nil
}

type countSaveFn struct {
	w	sync.WaitGroup
	m	sync.Mutex
	n	int
}

func (c *countSaveFn) save(_ *resource.Info, _ Reporter) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	go func() {
		c.m.Lock()
		c.n++
		c.m.Unlock()
		c.w.Done()
	}()
	return nil
}

type testBuilder int

func (t testBuilder) Visitor(_ ...resource.ErrMatchFunc) (resource.Visitor, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	infos := make(resource.InfoListVisitor, t)
	for i := range infos {
		infos[i] = &resource.Info{Mapping: &meta.RESTMapping{}}
	}
	return infos, nil
}
