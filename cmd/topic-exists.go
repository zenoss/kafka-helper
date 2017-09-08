// Copyright © 2017 Zenoss, Inc.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var topicExistsCmd = &cobra.Command{
	Use:   "topic-exists",
	Short: "Check to see if a topic exists",
	Run: func(cmd *cobra.Command, args []string) {
		tocheck := viper.GetString("topic")
		if tocheck == "" {
			fmt.Fprintln(os.Stderr, "Please specify a topic with --topic/-t.")
			os.Exit(1)
		}
		client, err := GetClient()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		topics, err := client.Topics()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _, topic := range topics {
			if topic == tocheck {
				os.Exit(0)
			}
		}
		os.Exit(1)
	},
}

func init() {
	RootCmd.AddCommand(topicExistsCmd)

	topicExistsCmd.Flags().StringP("topic", "t", "", "Topic to check")
	viper.BindPFlag("topic", topicExistsCmd.Flags().Lookup("topic"))
}
