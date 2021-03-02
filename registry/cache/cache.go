// Package cache provides a registry cache
package cache

import (
	"github.com/Augustu/go-micro/v2/registry"
	"github.com/Augustu/go-micro/v2/registry/cache"
)

// New returns a new cache
func New(r registry.Registry, opts ...cache.Option) cache.Cache {
	return cache.New(r, opts...)
}
