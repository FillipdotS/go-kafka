package util

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"os"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ReadConfig(configFile string) kafka.ConfigMap {

	m := make(map[string]kafka.ConfigValue)

	file, err := os.Open(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file: %s", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "#") && len(line) != 0 {
			kv := strings.Split(line, "=")
			parameter := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			m[parameter] = value
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Failed to read file: %s", err)
		os.Exit(1)
	}

	return m
}

// Simple function to convert a UUID to a colour int.
// This is used to colour the uuid output in the logs.
// Not ideal, but it's consistent.
func SimpleUuidToColourInt(uuid string) int {
	h := fnv.New64a()
	h.Write([]byte(uuid))
	hash := h.Sum64()

	// 41 - Red, 42 - Green etc, until 47
	return int(41 + (hash % 6))
}
