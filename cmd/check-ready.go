// Copyright © 2017 Zenoss, Inc.
//

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var checkReadyCmd = &cobra.Command{
	Use:   "check-ready",
	Short: "See if Kafka is ready",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		client.Close()
	},
}

func init() {
	RootCmd.AddCommand(checkReadyCmd)
}
