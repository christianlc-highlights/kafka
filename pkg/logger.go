package pkg

import (
	"os"
	log "github.com/sirupsen/logrus"
)

// func /////////////////////////////////////////

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stderr)
}

func Logerr(logger *log.Entry, msg string, err error) {
	logger.WithFields(log.Fields{
  	"err": err,
  }).Error(msg)
}