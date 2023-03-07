/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"context"

	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"

	"github.com/christianlc-highlights/kafka/pkg"

)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List topics",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, stop := pkg.InterruptContext(context.Background())
		bs    := pkg.Must(cmd.Flags().GetString("bootstrap-server"))
		logf  := log.WithFields(log.Fields{
	  	"trace": pkg.Trace("listCmd.Run", "cmd/list"),
	  	"bs": bs,
	  })
	  logf.Debug("Enter")
	  defer logf.Debug("Exit")
	  defer stop()

	  admin, err := pkg.Administrator(bs)
	  if err != nil {
	  	logf.WithFields(log.Fields{
	  		"error": err,
	  	}).Fatal("Failed to create administrative client")
	  }
	  logf.Info("Created administrative client")

	  topics, err := pkg.ListTopics(ctx, admin)
	  if err != nil {
	  	logf.WithFields(log.Fields{
	  		"error": err,
	  	}).Fatal("Failed to list topics")
	  }
	  logf.WithFields(log.Fields{
	  	"topics": topics,
	  }).Info("List topics")

	  for _, v := range topics {
	  	logf.WithFields(log.Fields{
	  		"topic": v,
	  	}).Debug("Iterate topic")
	  	fmt.Println(v)
	  }
	},
}
func init() {
	topicCmd.AddCommand(listCmd)
}
