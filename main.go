package main

import (
	"os"

	"github.com/christianlc-highlights/kafka/pkg"
	log "github.com/sirupsen/logrus"

)

// main /////////////////////////////////////////

func main() {
  logf := log.WithFields(
    log.Fields{
    	"trace": pkg.Trace("main", "main"),
    	"args": os.Args,
    },
  )
  logf.Debug("Enter")
  defer logf.Debug("Exit")

  if len(os.Args) < 2 {
  	logf.Fatal("Failed to pass bootstrap server addresses")
  }

  p, err := pkg.Producer(os.Args...)
  if err != nil {
  	pkg.Logerr(logf, "Failed to create producer", err)
  	panic(err)
  }
  logf.Info("Created kafka producer")

  err = pkg.Write(p, "TopicTest", "The Message Payload")
  if err != nil {
  	pkg.Logerr(logf, "Failed to write to broker", err)
  	panic(err)
  }
  logf.Info("Sent message to producer")

}