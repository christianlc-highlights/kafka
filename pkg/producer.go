package pkg

import (

	"github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/sirupsen/logrus"
)

// Create a producer and return pointer to instance
func Producer(servers string) (*kafka.Producer, error) {
  logf := log.WithFields(
    log.Fields{
    	"trace": Trace("Producer", "pkg/producer"),
    	"bootstrap": servers,
    },
  )

  logf.Debug("Enter")
  defer logf.Debug("Exit")

	return kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": servers,
	})
}

// Send a message to producer instance
func Write(p *kafka.Producer, topic, payload string) error {
	return p.Produce(&kafka.Message{
    TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
    Value: []byte(payload)},
    nil,
	)

}

// Send a delete request to broker through producer instance
func DeleteTopic(p *kafka.Producer, topic string) error {
	_, err := p.DeleteTopics(&kafka.DeleteTopicsRequest{
	  Topics: []string{topic},
	  Timeout: 10 * time.Second,
	})
	return err
}
