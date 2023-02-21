package pkg

import (
	"time"
	"github.com/confluentinc/confluent-kafka-go/kafka"

	log "github.com/sirupsen/logrus"
)

// export ///////////////////////////////////////

// Create a 
func Administrator(servers string) (*kafka.AdminClient, error) {
  logf := log.WithFields(log.Fields{
  	"trace": Trace("Adminstrator", "pkg/adminstrator"),
  	"bootstrap": servers,
  })
  logf.Debug("Enter")
  defer logf.Debug("Exit")

  return kafka.NewAdminClient(&kafka.ConfigMap{
  	"bootstrap.servers": broker,
  })
}

// Create a kafka topic
func CreateTopic(ctx context.Context, c *kafka.AdminClient, topic string, partitions, replicas int) error {
  logf := log.WithFields(log.Fields{
  	"trace": Trace("CreateTopic", "pkg/adminstrator"),
  	"topic": topic,
  	"partitions": partitions,
  	"replicas": replicas,
  })
  logf.Debug("Enter")
  defer logf.Debug("Exit")

  // https://github.com/confluentinc/confluent-kafka-go/blob/master/examples/admin_create_topic/admin_create_topic.go
	_, err := a.CreateTopics(
    ctx,
    // Multiple topics can be created simultaneously
    // by providing more TopicSpecification structs here.
    []kafka.TopicSpecification{{
      Topic:         topic,
      NumPartitions: partitions,
      ReplicationFactor: replicas,
    }},

    // Admin options
    kafka.SetAdminOperationTimeout(time.ParseDuration("5s"))
  )
  return err

}