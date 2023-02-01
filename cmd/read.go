/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"os/signal"
	"context"
	"syscall"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/christianlc-highlights/kafka/pkg"
	log "github.com/sirupsen/logrus"

)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Stream kafka topic to stdout",
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
	  	"trace": pkg.Trace("readCmd.Run", "cmd/read"),
	  	"bs": bs,
	  	"topic": topic,
	  })
	  logf.Debug("Enter")
	  defer logf.Debug("Exit")

	  c, err := pkg.Consumer(bs)
	  if err != nil {
	  	logf.WithFields(log.Fields{
	  		"error": err,
	  	}).Fatal("Failed to create consumer")
	  }
	  logf.Info("Created kafka consumer")

	  ever := true
	  for ever {
	  	logf.Debug("Iterate read loop")

	  	select {
	  	case <-ctx.Done():
	  		logf.Debug("Received cancel signal")
	  		ever = false

	  	default:
	  		logf.Debug("Read topic")

	  		msg, err := pkg.Read(c, topic)
	  		if err == nil {
	  			logf.WithFields(log.Fields{
	  				"msg": msg,
	  			}).Info("Received message from topic")
	  			fmt.Println(msg)

	  		} else if !pkg.ErrIsTimeout(err) {
	  			logf.WithFields(log.Fields{
	  				"error": err,
	  			}).Error("Failed to read from topic")
	  			ever = false
	  		}

	  	}
	  }
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
	readCmd.MarkFlagRequired("topic")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
