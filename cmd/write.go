/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"os/signal"
	"context"
	"syscall"
	"bufio"

	"github.com/spf13/cobra"
	"github.com/christianlc-highlights/kafka/pkg"
	log "github.com/sirupsen/logrus"

)

// writeCmd represents the write command
var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Stream stdin to kafka topic",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, stop := signal.NotifyContext(
			context.Background(),
			os.Interrupt,
			syscall.SIGTERM,
			syscall.SIGQUIT,
			syscall.SIGHUP,
		)
		defer stop()

		bs    := pkg.Must(cmd.Flags().GetString("bootstrap-server"))
		topic := pkg.Must(cmd.Flags().GetString("topic"))
		logf  := log.WithFields(log.Fields{
	  	"trace": pkg.Trace("writeCmd.Run", "cmd/write"),
	  	"topic": topic,
	  	"bootstrap": bs,
	  })
	  logf.Debug("Enter")
	  defer logf.Debug("Exit")

	  p, err := pkg.Producer(bs)
	  if err != nil {
	  	logf.WithFields(log.Fields{
	  		"error": err,
	  	}).Fatal("Failed to create producer")
	  }
	  logf.Info("Created kafka producer")

	  ever := true
	  scanner := bufio.NewScanner(os.Stdin)

	  for ever {
	  	logf.Debug("Iterate read loop")

	  	select {
	  	case <-ctx.Done():
	  		logf.Debug("Received cancel signal")
	  		ever = false

	  	default:
	  		logf.Debug("Blocking read on stdin")
	  		ever = scanner.Scan()
	  		logf.WithFields(log.Fields{
	  			"success": ever,
	  		}).Debug("Read token from stdin")

	  		if ever {
	  			token := scanner.Text()
	  			logf.WithFields(log.Fields{
	  				"token": token,
	  			}).Info("Write token to topic")

	  			if err := pkg.Write(p, topic, token); err != nil {
	  				logf.WithFields(log.Fields{
	  					"token": token,
	  					"error": err,
	  				}).Error("Failed write to topic")
	  				ever = false
	  			}
	  		}
	  	}
	  }

	  logf.Info("Exit read loop")
	},
}

func init() {
	rootCmd.AddCommand(writeCmd)
	writeCmd.MarkFlagRequired("topic")

}
