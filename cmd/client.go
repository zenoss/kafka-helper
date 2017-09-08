// Copyright © 2017 Zenoss, Inc.

package cmd

import (
	"fmt"
	"os"

	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
)

func GetClient() (sarama.Client, error) {
	brokers := viper.GetStringSlice("brokers")
	if len(brokers) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify one or more Kafka broker addresses with --broker/-b.")
		os.Exit(1)
	}
	return sarama.NewClient(viper.GetStringSlice("brokers"), nil)
}
