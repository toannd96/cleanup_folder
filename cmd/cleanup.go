package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

const (
	// maxAge is the default age after which folder
	// are considered old
	maxAge = 1 * time.Minute
)

var cleanupCmd = &cobra.Command{
	Use:   "folder [target folder]",
	Short: "Remove old folder",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		folder := args[0]

		oldCacheDirs, err := old(folder)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Printf("%d old folder found \n", len(oldCacheDirs))

		if len(oldCacheDirs) != 0 {
			for _, item := range oldCacheDirs {
				dir := filepath.Join(folder, item.Name())

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
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		result = append(result, entry)
	}

	return result, nil
}

// returns the list of folder older than max.
func olderThan(baseDir string) ([]os.FileInfo, error) {
	entries, err := listDirs(baseDir)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var oldCacheDirs []os.FileInfo
	for _, fi := range entries {
		if !isOld(fi.ModTime()) {
			continue
		}
		oldCacheDirs = append(oldCacheDirs, fi)
	}

	return oldCacheDirs, nil
}

// returns a list of folder with a modification time
func old(basedir string) ([]os.FileInfo, error) {
	return olderThan(basedir)
}

// returns true if the timestamp is considered old.
func isOld(t time.Time) bool {
	oldest := time.Now().Add(-maxAge)
	return t.Before(oldest)
}
