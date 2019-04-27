package importer

import (
	"github.com/hashicorp/golang-lru"
)

const (
	DefaultImageStreamLayerCacheSize = 2048
)

type ImageStreamLayerCache struct{ *lru.Cache }

func NewImageStreamLayerCache(size int) (ImageStreamLayerCache, error) {
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
	c, err := lru.New(size)
	if err != nil {
		return ImageStreamLayerCache{}, err
	}
	return ImageStreamLayerCache{Cache: c}, nil
}
