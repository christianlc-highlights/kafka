/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/christianlc-highlights/kafka/pkg"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a topic partition",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, stop := pkg.InterruptContext(context.Background())
		bs    := pkg.Must(cmd.PersistentFlags().GetString("bootstrap-server"))
		topic := pkg.Must(cmd.PersistentFlags().GetString("topic"))
		partitions := pkg.Must(cmd.Flags().GetString("partitions"))
		replicas := pkg.Must(cmd.Flags().GetString("replicas"))

		logf  := log.WithFields(log.Fields{
	  	"trace": pkg.Trace("createCmd.Run", "cmd/create"),
	  	"bs": bs,
	  	"topic": topic,
	  	"partitions": partitions,
	  	"replicas": replicas,
	  })
	  logf.Debug("Enter")
	  defer logf.Debug("Exit")

	  // create administrator client to kafka using bootstrap
	  // servers passed as a required flag to this cmd
	  admin, err := pkg.Administrator(bs)
	  if err != nil { 
	  	logf.WithFields(log.Fields{
	  		"error": err,
	  	}).Fatal("Failed to create administrative client")
	  }
	  logf.Debug("Created administrative client")

	  err := pkg.CreateTopic(
	  	ctx,
	  	admin,
	  	topic,
	  	partitions,
	  	replicas,
	  )
	  if err != nil {
	  	logf.WithFields(log.Fields{
	  		"error": err,
	  	}).Fatal("Failed to create topic")
	  }
	  logf.Info("Created topic")

	},
}

func init() {
	topic.AddCommand(createCmd)
	createCmd.Flags().IntP("partitions", "p", 1, "Number of partitions")
	createCmd.Flags().IntP("replicas", "r", 1, "Number of replicas")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
