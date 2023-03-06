package pkg

import (
	"time"
	"context"

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
  	"bootstrap.servers": servers,
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
	_, err := c.CreateTopics(
    ctx,
    // Multiple topics can be created simultaneously
    // by providing more TopicSpecification structs here.
    []kafka.TopicSpecification{{
      Topic:         topic,
      NumPartitions: partitions,
      ReplicationFactor: replicas,
    }},

    // Admin options
    kafka.SetAdminOperationTimeout(Must(time.ParseDuration("5s"))),
  )
  return err

}

// Send a delete request to broker through producer instance
func DeleteTopic(ctx context.Context, a *kafka.AdminClient, topic string) error {
  logf := log.WithFields(log.Fields{
  	"trace": Trace("DeleteTopic", "pkg/adminstrator"),
  	"topic": topic,
  })
  logf.Debug("Enter")
  defer logf.Debug("Exit")

	results, err := a.DeleteTopics(
		ctx, []string{topic}, nil,
	)
	logf.WithFields(log.Fields{
		"results": results,
	})

	return err
}

// List all topics currently defined in cluster
func ListTopics(ctx context.Context, a *kafka.AdminClient) ([]string, error) {
	var result []string
	var pollerr error

	ever  := true
	datch := make(chan string)
	errch := make(chan error)
  logf := log.WithFields(log.Fields{
  	"trace": Trace("ListTopics", "pkg/adminstrator"),
  	"task": "Retrieve topics",
  })
  logf.Debug("Enter")
  defer logf.Debug("Exit")

  go func() {
  	logf := logf.WithFields(log.Fields{
  		"thread": true,
  	})
  	logf.Info("Start retrieve topics")
  	defer logf.Debug("Stop retrieve topics")

  	m, err := a.GetMetadata(nil, true, 0)
  	if err == nil {
  		logf.Info("Succeeded retrieving metadata")

  		for k, v := range m.Topics {
  			logf.WithFields(log.Fields{
  				"key": k,
  				"topic": v.Topic,
  			}).Debug("Iterate metadata")
  			datch <-v.Topic
  		}
  	}

  	close(datch)
  	close(errch)
  }()

  logf.Debug("Begin polling topic")
  for ever {
  	logf.Debug("Iterate polling topic")

	  select {
	  case t, ok := <-datch:
  		logf.WithFields(log.Fields{
	  		"topic": t,
	  		"ok": ok,
	  	}).Debug("Read topic")

	  	if ok {
		  	logf.WithFields(log.Fields{
		  		"topic": t,
		  	}).Info("Append topic")
	  		result = append(result, t)
	  	}
	  	ever = ok

	  case pollerr = <-errch:
	  	logf.Info("Error polling topic")
	  	ever = false
	  case <-ctx.Done():
	  	logf.Info("Context cancelled")
	  	ever = false
	  }
	}
	logf.Debug("End polling topic")

	return result, pollerr
}
