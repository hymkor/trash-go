package freedesktop

import (
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

const (
	// the flag for testing
	always_copy_and_move = false
)

func _move(oldPath, newPath string) error {
	fd, err := os.Open(oldPath)
	if err != nil {
		return err
	}
	succeeded := false
	defer func() {
		fd.Close()
		if succeeded {
			os.Remove(oldPath)
		}
	}()

	stat, err := fd.Stat()
	if err != nil {
		return err
	}
	if stat.IsDir() {
		err := os.Mkdir(newPath, stat.Mode()&fs.ModePerm)
		if err != nil {
			return err
		}
		files, err := fd.Readdirnames(-1)
		if err != nil {
			return err
		}
		for _, f := range files {
			if f == "." || f == ".." {
				continue
			}
			err := _move(filepath.Join(oldPath, f), filepath.Join(newPath, f))
			if err != nil {
				return err
			}
		}
		succeeded = true
	} else {
		newFile, err := os.OpenFile(newPath,
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
			stat.Mode()&fs.ModePerm)
		if err != nil {
			return err
		}
		if _, err := io.Copy(newFile, fd); err != nil {
			newFile.Close()
			return err
		}
		if err := newFile.Close(); err != nil {
			return err
		}
		mtime := stat.ModTime()
		atime := time.Now()
		os.Chtimes(newPath, atime, mtime)
		succeeded = true
	}
	return nil
}

func move(oldPath, newPath string) error {
	if !always_copy_and_move {
		err := os.Rename(oldPath, newPath)
		if err == nil {
			return nil
		}
		var osLinkError *os.LinkError
		if !errors.As(err, &osLinkError) {
			return err
		}
	}
	return _move(oldPath, newPath)
}
