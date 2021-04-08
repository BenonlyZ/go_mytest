package main

import (
	"log"

	"github.com/Shopify/sarama"
)

var KAFKA_BROKERS = []string{"10.0.0.0:9093", "10.0.0.1:9093", "10.0.0.2:9093"}

func main() {
	if checkKafka() {
		log.Printf("Check kafka status success\n")
	} else {
		log.Printf("Check kafka status failure\n")
	}
}

func checkKafka() bool {
	config := sarama.NewConfig()
	config.Net.SASL.Enable = true
	config.Net.SASL.Mechanism = "PLAIN"
	config.Net.SASL.User = "User"
	config.Net.SASL.Password = "123"
	client, err := sarama.NewClient(KAFKA_BROKERS, config)
	if err != nil {
		log.Printf("ERROR: Unable to create kafka client, err=[%v]", err)
		return false
	}
	defer client.Close()

	topics, err := client.Topics()
	if err != nil {
		log.Printf("ERROR: Unable to list kafka topics, err=[%v]", err)
		return false
	}

	log.Printf("Get total topics count=[%d]\n", len(topics))
	for i, topic := range topics {
		log.Printf("\tTopic[%d]: [%s]\n", i, topic)
	}

	return true
}
