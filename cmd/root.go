// Copyright © 2017 Zenoss, Inc.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "kafka-helper",
	Short: "A small utility to check on Kafka",
	Long: `kafka-helper is a small, self-contained binary you can use
to perform minimal interactions with Kafka, without requiring
Java or librdkafka.`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringSliceP("broker", "b", []string{"kafka:9092"}, "Kafka broker (e.g. 1.2.3.4:9092). May be specified multiple times")
	viper.BindPFlag("brokers", RootCmd.PersistentFlags().Lookup("broker"))
	viper.SetDefault("brokers", []string{"kafka:9092"})
}

func initConfig() {
	viper.SetEnvPrefix("kafka_helper")
	viper.AutomaticEnv()
}
