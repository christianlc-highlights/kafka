package pkg

import (
	"time"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/sirupsen/logrus"
)

// create and return pointer to kafka consumer
func Consumer(bootstrap ...string) (*kafka.Consumer, error) {
  servers := strings.Join(bootstrap[1:], ",")
  logf := log.WithFields(
    log.Fields{
    	"trace": Trace("Consumer", "pkg/consumer"),
    	"bootstrap": servers,
    },
  )

  logf.Debug("Enter")
  defer logf.Debug("Exit")

  return kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": servers,
		"group.id": "test",
		"auto.offset.reset": "earliest",
  })

}

// Blocking read against kafka broker
func Read(c *kafka.Consumer, topic string) (string, error) {
  logf := log.WithFields(log.Fields{
  	"trace": Trace("Read", "pkg/consumer"),
  	"topic": topic,
  })

  logf.Debug("Enter")
  defer logf.Debug("Exit")

	msg, err := c.ReadMessage(time.Second)
	if err != nil {
		return "", err
	}

	return string(msg.Value), nil
}

// Determine if error is a read timeout
func ErrIsTimeout(err error) bool {
  logf := log.WithFields(log.Fields{
  	"trace": Trace("ErrIsTimeout", "pkg/consumer"),
  	"error": err,
  })
  logf.Debug("Enter")
  defer logf.Debug("Exit")

  return err.(kafka.Error).Code() == kafka.ErrTimedOut
}