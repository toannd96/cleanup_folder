package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	multios "cleanup/pkg/multios"
)

var cleanupCmd = &cobra.Command{
	Use:   "folder [max number of hours (integer) the old folder to be cleanup] [target folder]",
	Short: "Remove old folder by max number of hours of exist",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		number, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		maxAge := time.Duration(number) * time.Hour
		folder := args[1]

		oldCacheDirs, err := old(folder, maxAge)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Printf("%d old item found \n", len(oldCacheDirs))

		if len(oldCacheDirs) != 0 {
			for _, item := range oldCacheDirs {
				dir := filepath.Join(folder, item.Name())

				err = os.RemoveAll(dir)
				if err != nil {
					fmt.Println(err.Error())
					os.Exit(1)
				}
				fmt.Printf("remove item %s \n", dir)
			}
			fmt.Printf("cleanup %s completed \n", folder)
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
}

// returns the list of directory.
func listDirs(baseDir string) ([]os.FileInfo, error) {
	f, err := os.Open(baseDir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println(err.Error())
		os.Exit(1)
	}

	entries, err := f.Readdir(-1)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	result := make([]os.FileInfo, 0, len(entries))
	result = append(result, entries...)

	return result, nil
}

// returns the list of folder older than max.
func olderThan(baseDir string, maxAge time.Duration) ([]os.FileInfo, error) {
	entries, err := listDirs(baseDir)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var oldCacheDirs []os.FileInfo
	for _, fi := range entries {
		if !isOld(multios.Item(fi), maxAge) {
			continue
		}
		oldCacheDirs = append(oldCacheDirs, fi)
	}

	return oldCacheDirs, nil
}

// returns a list of folder with a modification time
func old(basedir string, maxAge time.Duration) ([]os.FileInfo, error) {
	return olderThan(basedir, maxAge)
}

// returns true if the timestamp is considered old.
func isOld(t time.Time, maxAge time.Duration) bool {
	oldest := time.Now().Add(-maxAge)
	return t.Before(oldest)
}
