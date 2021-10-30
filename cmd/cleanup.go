package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	str2duration "github.com/xhit/go-str2duration/v2"
)

var (
	age    string
	folder string
)

var cleanupCmd = &cobra.Command{
	Use:   "folder",
	Short: "Remove old folder",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		age = args[0]
		folder = args[1]

		// convert string max age old folder from string to time duration
		maxAge, err := str2duration.ParseDuration(age)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		oldCacheDirs, err := old(folder, maxAge)
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
	cleanupCmd.Flags().StringVarP(&age, "age", "a", age, "Max age old folder to be cleanup example 1s, 1m, 1h, 1d, 1w, 1w2d3s")
	cleanupCmd.Flags().StringVarP(&folder, "folder", folder, "/home/dactoan/upload", "Target folder to be cleanup")
	rootCmd.AddCommand(cleanupCmd)
}

// returns the list of directory.
func listDirs(baseDir string, maxAge time.Duration) ([]os.FileInfo, error) {
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
func olderThan(baseDir string, maxAge time.Duration) ([]os.FileInfo, error) {
	entries, err := listDirs(baseDir, maxAge)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var oldCacheDirs []os.FileInfo
	for _, fi := range entries {
		if !isOld(fi.ModTime(), maxAge) {
			continue
		}
		oldCacheDirs = append(oldCacheDirs, fi)
	}

	return oldCacheDirs, nil
}

// returns a list of folder with a modification time of more
// than 30 days ago.
func old(basedir string, maxAge time.Duration) ([]os.FileInfo, error) {
	return olderThan(basedir, maxAge)
}

// returns true if the timestamp is considered old.
func isOld(t time.Time, maxAge time.Duration) bool {
	oldest := time.Now().Add(-maxAge)
	return t.Before(oldest)
}
