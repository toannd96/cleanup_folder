package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version bool
)

func init() {
	cobra.OnInitialize()
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "show current version of CLI")
}

var rootCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "A simple CLI use to cleanup old folder",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
