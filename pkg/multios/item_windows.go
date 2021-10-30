package multios

func Item(fi fs.FileInfo) time.Time {
	var mtimeLocal time.Time
	if stat, ok := fi.Sys().(*syscall.Win32FileAttributeData); ok {
		mtimeLocal = time.Unix(0, stat.LastWriteTime.Nanoseconds()).UTC()
	}
	return mtimeLocal
}
