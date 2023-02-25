/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
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

	  p, err := pkg.Producer(bs)
	  if err != nil {
	  	logf.WithFields(log.Fields{
	  		"error": err,
	  	}).Fatal("Failed to create producer")
	  }
	  logf.Info("Created kafka producer")

	  if err := pkg.DeleteTopic(p, topic); err != nil {
	  	logf.WithFields(log.Fields{
	  		"error": err,
	  	}).Fatal("Failed to delete topic")
	  }
	  logf.Info("Deleted topic")
	},
}

func init() {
	topic.AddCommand(deleteCmd)
}
