package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
	// rootCmd.Flags().BoolVarP(&version, "version", "v", false, "show current version of CLI")
}

var rootCmd = &cobra.Command{
	Use:     "cleanup",
	Short:   "A simple CLI use to cleanup old folder",
	Version: "1.0",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
