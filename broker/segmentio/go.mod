module github.com/Augustu/go-plugins/broker/segmentio/v2

go 1.16

require (
	github.com/google/uuid v1.1.1
	github.com/Augustu/go-micro/v2 v2.9.3
	github.com/Augustu/go-plugins/broker/kafka/v2 v2.3.0
	github.com/Augustu/go-plugins/codec/segmentio/v2 v2.3.0
	github.com/segmentio/kafka-go v0.3.7
)

replace github.com/Augustu/go-plugins/codec/segmentio/v2 => ../../codec/segmentio
