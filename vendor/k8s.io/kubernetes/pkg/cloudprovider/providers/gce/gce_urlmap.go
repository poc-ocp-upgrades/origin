package gce

import (
	compute "google.golang.org/api/compute/v1"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/filter"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/meta"
)

func newURLMapMetricContext(request string) *metricContext {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return newGenericMetricContext("urlmap", request, unusedMetricLabel, unusedMetricLabel, computeV1Version)
}
func (g *Cloud) GetURLMap(name string) (*compute.UrlMap, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newURLMapMetricContext("get")
	v, err := g.c.UrlMaps().Get(ctx, meta.GlobalKey(name))
	return v, mc.Observe(err)
}
func (g *Cloud) CreateURLMap(urlMap *compute.UrlMap) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newURLMapMetricContext("create")
	return mc.Observe(g.c.UrlMaps().Insert(ctx, meta.GlobalKey(urlMap.Name), urlMap))
}
func (g *Cloud) UpdateURLMap(urlMap *compute.UrlMap) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newURLMapMetricContext("update")
	return mc.Observe(g.c.UrlMaps().Update(ctx, meta.GlobalKey(urlMap.Name), urlMap))
}
func (g *Cloud) DeleteURLMap(name string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newURLMapMetricContext("delete")
	return mc.Observe(g.c.UrlMaps().Delete(ctx, meta.GlobalKey(name)))
}
func (g *Cloud) ListURLMaps() ([]*compute.UrlMap, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	mc := newURLMapMetricContext("list")
	v, err := g.c.UrlMaps().List(ctx, filter.None)
	return v, mc.Observe(err)
}
