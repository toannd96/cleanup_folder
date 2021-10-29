package pkg

import (
	"errors"
	"fmt"
	"os"
	"time"
)

const (
	// maxAge is the default age (30 days) after which folder
	// are considered old
	maxAge = 24 * time.Hour
)

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

// returns a list of folder with a modification time of more
// than 30 days ago.
func Old(basedir string) ([]os.FileInfo, error) {
	return olderThan(basedir)
}

// returns true if the timestamp is considered old.
func isOld(t time.Time) bool {
	oldest := time.Now().Add(-maxAge)
	return t.Before(oldest)
}
