/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"

	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"

	"github.com/christianlc-highlights/kafka/pkg"

)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a topic",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, stop := pkg.InterruptContext(context.Background())
		bs    := pkg.Must(cmd.PersistentFlags().GetString("bootstrap-server"))
		topic := pkg.Must(cmd.PersistentFlags().GetString("topic"))
		logf  := log.WithFields(log.Fields{
	  	"trace": pkg.Trace("deleteCmd.Run", "cmd/delete"),
	  	"bs": bs,
	  	"topic": topic,
	  })
	  logf.Debug("Enter")
	  defer logf.Debug("Exit")
	  defer stop()

	  c, err := pkg.Administrator(bs)
	  if err != nil {
	  	logf.WithFields(log.Fields{
	  		"error": err,
	  	}).Fatal("Failed to create administrative client")
	  }
	  logf.Info("Created kafka administrative client")

	  if err := pkg.DeleteTopic(ctx, c, topic); err != nil {
	  	logf.WithFields(log.Fields{
	  		"error": err,
	  	}).Fatal("Failed to delete topic")
	  }
	  logf.Info("Deleted topic")
	},
}

func init() {
	topicCmd.AddCommand(deleteCmd)
}
