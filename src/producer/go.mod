module github.com/fillipdots/kafka-go-experiment/producer

go 1.18

replace github.com/fillipdots/kafka-go-experiment/util => ../util

require (
	github.com/confluentinc/confluent-kafka-go v1.9.2
	github.com/fillipdots/kafka-go-experiment/util v1.0.0
	github.com/google/uuid v1.3.0 // indirect
)
