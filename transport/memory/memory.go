// Package memory is an in-memory transport
package memory

import (
	"github.com/Augustu/go-micro/v2/config/cmd"
	"github.com/Augustu/go-micro/v2/transport"
	"github.com/Augustu/go-micro/v2/transport/memory"
)

func init() {
	cmd.DefaultTransports["memory"] = NewTransport
}

func NewTransport(opts ...transport.Option) transport.Transport {
	return memory.NewTransport(opts...)
}
