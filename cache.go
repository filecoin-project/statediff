package statediff

import (
	"sync"

	lru "github.com/hashicorp/golang-lru"
	"github.com/ipld/go-ipld-prime"
)

type NodeCache struct {
	cache *lru.ARCCache
}

func (nc *NodeCache) Get(key string) ipld.Node {
	v, ok := nc.cache.Get(key)
	if ok {
		return v.(ipld.Node)
	}
	return nil
}

func (nc *NodeCache) Add(key string, n ipld.Node) {
	nc.cache.Add(key, n)
}

func NewNodeCache(size int) *NodeCache {
	c, err := lru.NewARC(size)
	if err != nil {
		return nil
	}
	return &NodeCache{
		cache: c,
	}
}

var singletonNodeCache *NodeCache
var snocsetup sync.Once

func setupSingleton() {
	singletonNodeCache = NewNodeCache(32 * 1024)
}

func GetGlobalNodeCache() *NodeCache {
	snocsetup.Do(setupSingleton)
	return singletonNodeCache
}
