// Package grpc provides a grpc transport
package grpc

import (
	"github.com/Augustu/go-micro/v2/config/cmd"
	"github.com/Augustu/go-micro/v2/transport"
	"github.com/Augustu/go-micro/v2/transport/grpc"
)

func init() {
	cmd.DefaultTransports["grpc"] = NewTransport
}

func NewTransport(opts ...transport.Option) transport.Transport {
	return grpc.NewTransport(opts...)
}
