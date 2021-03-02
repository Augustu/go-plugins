// Package service uses the registry service
package service

import (
	"github.com/Augustu/go-micro/v2/config/cmd"
	"github.com/Augustu/go-micro/v2/registry"
	"github.com/Augustu/go-micro/v2/registry/service"
)

func init() {
	cmd.DefaultRegistries["service"] = NewRegistry
}

// NewRegistry returns a new registry service client
func NewRegistry(opts ...registry.Option) registry.Registry {
	return service.NewRegistry(opts...)
}
