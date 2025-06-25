package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

/*
 * @brief InitKafka initializes the Kafka writer for logging URL visits.
 * 
 * This function sets up a Kafka writer with the specified broker and topic for 
 * asynchronously logging URL visits. The writer is configured to use the least 
 * bytes load balancing strategy.
 */
func InitKafka() {
	writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"kafka:9092"},
		Topic:    "url_visits",
		Balancer: &kafka.LeastBytes{},
	})
	log.Println("Kafka writer initialized")
}

/*
 * @brief LogVisit logs a URL visit to Kafka asynchronously.
 *
 * This function publishes a message to the "url_visits" Kafka topic with the
 * given short code and IP address. The message is formatted as a string with
 * the current timestamp and the IP address.
 *
 * @param code The short code of the URL that was visited.
 * @param ip The IP address of the user who visited the URL.
 */
func LogVisit(code string, ip string) {
	msg := kafka.Message{
		Key:   []byte(code),
		Value: []byte(time.Now().Format(time.RFC3339) + " - visited from " + ip),
	}

	err := writer.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Printf("Kafka write error: %v", err)
	} else {
		log.Printf("✅ Kafka write success: %s from %s", code, ip)
	}
}
