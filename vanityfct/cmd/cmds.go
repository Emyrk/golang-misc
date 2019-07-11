package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vanityfct",
	Short: "Will generate vanity fct addresses",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var vanity = &cobra.Command{
	Use:   "vanity",
	Short: "target these vanity prefixes",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
