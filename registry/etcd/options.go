package etcd

import (
	"github.com/Augustu/go-micro/v2/registry"
	"github.com/Augustu/go-micro/v2/registry/etcd"
)

// Auth allows you to specify username/password
func Auth(username, password string) registry.Option {
	return etcd.Auth(username, password)
}
