package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fillipdots/kafka-go-experiment/util"
	"github.com/fillipdots/kafka-go-experiment/util/event"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	fmt.Println("Starting consumer...")

	topic := "purchases"

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <config-file-path>\n",
			os.Args[0])
		os.Exit(1)
	}

	configFile := os.Args[1]
	config := util.ReadConfig(configFile)
	config["group.id"] = "kafka-go-getting-started"
	config["auto.offset.reset"] = "earliest"

	consumer, err := kafka.NewConsumer(&config)

	if err != nil {
		fmt.Printf("Failed to create consumer: %s", err)
		os.Exit(1)
	}

	err = consumer.SubscribeTopics([]string{topic}, nil)

	if err != nil {
		fmt.Printf("Failed to subscribe to topic %s: %s", topic, err)
		os.Exit(1)
	}

	// Set up a channel for handling Ctrl-C, etc
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Consumer set up, starting to listen...")

	// Process messages
	for run := true; run; {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev, err := consumer.ReadMessage(100 * time.Millisecond)

			if err != nil {
				// Errors are informational and automatically handled by the consumer
				// Error can also just be a timeout, i.e. not a problem
				continue
			}

			var receivedEvent event.Event
			json.Unmarshal(ev.Value, &receivedEvent)

			uuidToColour := util.SimpleUuidToColourInt(receivedEvent.Id.String())

			fmt.Printf("Consumed event: key = \x1b[%dm%-8s\x1b[0m buyer = %-10s item = %-12s price = %-5d\n", uuidToColour, string(ev.Key), receivedEvent.Buyer, receivedEvent.Item, receivedEvent.Price)
		}
	}

	consumer.Close()
}
