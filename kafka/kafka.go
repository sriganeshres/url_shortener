package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

func InitKafka() {
	writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"kafka:9092"},
		Topic:    "url_visits",
		Balancer: &kafka.LeastBytes{},
	})
	log.Println("Kafka writer initialized")
}

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
