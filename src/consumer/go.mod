module github.com/fillipdots/kafka-experiment/consumer

go 1.18

replace github.com/fillipdots/kafka-experiment/util => ../util

require (
	github.com/confluentinc/confluent-kafka-go v1.9.2
	github.com/fillipdots/kafka-experiment/util v1.0.0
)
