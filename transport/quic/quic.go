// Package quic provides a QUIC based transport
package quic

import (
	"github.com/Augustu/go-micro/v2/config/cmd"
	"github.com/Augustu/go-micro/v2/transport"
	"github.com/Augustu/go-micro/v2/transport/quic"
)

func init() {
	cmd.DefaultTransports["quic"] = NewTransport
}

func NewTransport(opts ...transport.Option) transport.Transport {
	return quic.NewTransport(opts...)
}
