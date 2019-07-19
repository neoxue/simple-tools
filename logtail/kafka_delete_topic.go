package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	//if len(os.Args) < 3 {
	//	fmt.Fprintf(os.Stderr,
	//		"Usage: %s <broker> <topic1> <topic2> ..\n",
	//		os.Args[0])
	//	os.Exit(1)
	//}

	//broker := os.Args[1]
	//topics := os.Args[2:]
	broker := "10.83.0.44:9092"
	topics := []string{"logtail_esdoc", "logtail_list", "logtail_default", "logtail_kubernetes"}
	topics = []string{"pub_day_report_import"}

	// Create a new AdminClient.
	// AdminClient can also be instantiated using an existing
	// Producer or Consumer instance, see NewAdminClientFromProducer and
	// NewAdminClientFromConsumer.
	a, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		fmt.Printf("Failed to create Admin client: %s\n", err)
		os.Exit(1)
	}

	// Contexts are used to abort or limit the amount of time
	// the Admin call blocks waiting for a result.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Delete topics on cluster
	// Set Admin options to wait for the operation to finish (or at most 60s)
	maxDur, err := time.ParseDuration("60s")
	if err != nil {
		panic("ParseDuration(60s)")
	}
	topic := "logtail_esdoc"
	meta, err := a.GetMetadata(&topic, true, 100)
	fmt.Println(err)
	fmt.Println(meta)

	results, err := a.DeleteTopics(ctx, topics, kafka.SetAdminOperationTimeout(maxDur))
	fmt.Println(results)
	fmt.Println(err)
	if err != nil {
		fmt.Printf("Failed to delete topics: %v\n", err)
		os.Exit(1)
	}

	// Print results
	for _, result := range results {
		fmt.Printf("%s\n", result)
	}

	a.Close()
}
