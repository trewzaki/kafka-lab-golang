package main

import (
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var kafkaServer, kafkaTopic string
var producer *kafka.Producer

func init() {
	kafkaServer = readFromENV("KAFKA_BROKER", "pkc-ldvr1.asia-southeast1.gcp.confluent.cloud:9092")
	kafkaTopic = readFromENV("KAFKA_TOPIC", "test")

	fmt.Println("Kafka Broker - ", kafkaServer)
	fmt.Println("Kafka topic - ", kafkaTopic)

	var err error
	producer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaServer,
		"security.protocol": "SASL_SSL",
		"sasl.mechanisms":   "PLAIN",
		"sasl.username":     "J455WLJIMGCBRZEO",
		"sasl.password":     "sH8SgRiEy3ManX9R0HT7CTDWtcB0m1s3oZJ6lnlSMWUBydjHEpx+d+r45qL/l2CT",
		"acks":              "all"})
	if err != nil {
		panic(err)
	}
}
func main() {

	defer producer.Close()
	// Delivery report handler for produced messages
	// go func() {
	// 	for e := range producer.Events() {
	// 		switch ev := e.(type) {
	// 		case *kafka.Message:
	// 			if ev.TopicPartition.Error != nil {
	// 				fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
	// 			} else {
	// 				fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
	// 			}
	// 		}
	// 	}
	// }()
	// Produce messages to topic (asynchronously)
	// topic := "test"

	for i := 0; i < 5; i++ {
		KafkaProducer("test", []byte(fmt.Sprintf("%d", i)))
	}
	time.Sleep(time.Millisecond * 1000)
}
func readFromENV(key, defaultVal string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultVal
	}
	return value
}

// KafkaProducer ..
func KafkaProducer(topic string, msg []byte) {
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msg,
	}, nil)

	// time.Sleep(time.Millisecond * 100)
}
