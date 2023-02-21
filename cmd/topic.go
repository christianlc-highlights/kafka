/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"bufio"

	"github.com/spf13/cobra"
	"github.com/christianlc-highlights/kafka/pkg"
	log "github.com/sirupsen/logrus"

)
// topicCmd represents the topic command
var topicCmd = &cobra.Command{
	Use:   "topic",
	Short: "Manage a kafka topic",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, stop := pkg.InterruptContext(context.Background())
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
	},
}

func init() {
	rootCmd.AddCommand(topicCmd)
}
