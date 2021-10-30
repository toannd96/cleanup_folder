package multios

import (
	"io/fs"
	"syscall"
	"time"
)

func Item(fi fs.FileInfo) time.Time {
	var mtime time.Time
	if stat, ok := fi.Sys().(*syscall.Stat_t); ok {
		mtime = time.Unix(stat.Mtim.Unix()).UTC()
	}
	return mtime
}
