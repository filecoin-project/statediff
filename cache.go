package statediff

import (
	"fmt"
	"os"
	"strconv"
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
	rcsn := 32 * 1024
	rcs := os.Getenv("REALIZED_CACHE_SIZE")
	if rcs != "" {
		var err error
		rcsn, err = strconv.Atoi(rcs)
		if err != nil {
			fmt.Printf("Failed to parse cache size: %v\n", err)
			rcsn = 32 * 1024
		}
	}
	singletonNodeCache = NewNodeCache(rcsn)
}

func GetGlobalNodeCache() *NodeCache {
	snocsetup.Do(setupSingleton)
	return singletonNodeCache
}
