package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
	
	"github.com/fillipdots/kafka-experiment/util"
	"github.com/confluentinc/confluent-kafka-go/kafka"
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

	// Go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()

	users := [...]string{"eabara", "jsmith", "sgarcia", "jbernard", "htanaka", "awalther"}
	items := [...]string{"book", "alarm clock", "t-shirts", "gift card", "batteries"}

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
			key := users[rand.Intn(len(users))]
			data := items[rand.Intn(len(items))]
			producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Key:            []byte(key),
				Value:          []byte(data),
			}, nil)

			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		}
	}

	// Wait for all messages to be delivered
	producer.Flush(15 * 1000)
	producer.Close()
}
