package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/fillipdots/kafka-go-experiment/util"
	"github.com/fillipdots/kafka-go-experiment/util/event"
)

func main() {

	fmt.Println("Starting producer...")

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <config-file-path>\n",
			os.Args[0])
		os.Exit(1)
	}
	configFile := os.Args[1]
	config := util.ReadConfig(configFile)

	topic := "purchases"
	producer, err := kafka.NewProducer(&config)

	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	// Go-routine to handle message delivery reports
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					var sentEvent event.Event
					json.Unmarshal(ev.Value, &sentEvent)
					uuidToColour := util.SimpleUuidToColourInt(sentEvent.Id.String())

					fmt.Printf("Produced event: key = \x1b[%dm%-10s\x1b[0m value = %s\n", uuidToColour, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()

	// Set up a channel for handling Ctrl-C, etc
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Producer set up, starting to send events...")

	for run := true; run; {
		select {
		case sig := <-sigChannel:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			randomEvent := event.NewRandom()
			eventJson, _ := json.Marshal(randomEvent)

			producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Key:            []byte(randomEvent.Id.String()),
				Value:          []byte(eventJson),
			}, nil)

			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		}
	}

	// Wait for all messages to be delivered
	producer.Flush(15 * 1000)
	producer.Close()
}
