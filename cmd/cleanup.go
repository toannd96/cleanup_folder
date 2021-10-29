package cmd

import (
	"cleanup/pkg"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const (
	baseDir = "/home/dactoan/upload"
)

var cleanupCmd = &cobra.Command{
	Use:   "folder",
	Short: "Remove old folder",
	Run: func(cmd *cobra.Command, args []string) {
		oldCacheDirs, err := pkg.Old(baseDir)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Printf("%d old folder found \n", len(oldCacheDirs))

		if len(oldCacheDirs) != 0 {
			for _, item := range oldCacheDirs {
				dir := filepath.Join(baseDir, item.Name())
				err = os.RemoveAll(dir)
				if err != nil {
					fmt.Println(err.Error())
					os.Exit(1)
				}
				fmt.Printf("removing folder %s \n", dir)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
}
