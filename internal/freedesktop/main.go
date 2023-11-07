package freedesktop

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

type TrashCan struct {
	topDir string
	create bool
}

func (t *TrashCan) mkdir(path string) error {
	if t.create {
		return os.MkdirAll(path, 0755)
	} else {
		return os.Mkdir(path, 0755)
	}
}

func eachDirectorysizes(path string) int64 {
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0
	}
	sum := int64(0)
	for _, e := range entries {
		if e.IsDir() {
			sum += eachDirectorysizes(filepath.Join(path, e.Name()))
		} else {
			info, err := e.Info()
			if err == nil {
				sum += info.Size()
			}
		}
	}
	return sum
}

func (T *TrashCan) updateDirectorySize(newName string, stat os.FileInfo) error {
	size := eachDirectorysizes(newName)

	tmpDirectorySizePath := filepath.Join(T.topDir,
		fmt.Sprintf("directorysizes.%d", os.Getpid()))
	file, err := os.Create(tmpDirectorySizePath)
	if err != nil {
		return err
	}

	directorySizesPath := filepath.Join(T.topDir, "directorysizes")
	orig, err := os.Open(directorySizesPath)
	if err == nil {
		io.Copy(file, orig)
		orig.Close()
	}
	fmt.Fprintf(file, "%d %d %s\n",
		size, stat.ModTime().Unix(), percent(filepath.Base(newName)))
	file.Close()
	return os.Rename(tmpDirectorySizePath, directorySizesPath)
}

func percent(s string) string {
	var buffer strings.Builder
	buffer.Grow(len(s))
	for _, c := range s {
		if c == '%' {
			buffer.WriteString("%%")
		} else if c <= ' ' {
			fmt.Fprintf(&buffer, "%%%02X", c)
		} else {
			buffer.WriteRune(c)
		}
	}
	return buffer.String()
}

func (T *TrashCan) Throw(filename string) error {
	topDir := filepath.Dir(T.topDir)

	oldName, err := filepath.Abs(filename)
	if err != nil {
		return err
	}
	_oldName, err := filepath.Rel(topDir, oldName)
	if err == nil && !strings.HasPrefix(_oldName, ".."+string(os.PathSeparator)) {
		oldName = _oldName
	}

	filesDir := filepath.Join(T.topDir, "files")
	if err := T.mkdir(filesDir); err != nil && !os.IsExist(err) {
		return err
	}
	infoDir := filepath.Join(T.topDir, "info")
	if err := T.mkdir(infoDir); err != nil && !os.IsExist(err) {
		return err
	}

	for try := 0; try < 10; try++ {
		uuid1, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		uniqName := uuid1.String()

		infoPath := filepath.Join(infoDir, uniqName+".trashinfo")
		file, err := os.OpenFile(infoPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0755)
		if err == nil {
			fmt.Fprintf(file, "[Trash Info]\nPath=%s\nDeletionDate=%s\n",
				percent(oldName),
				time.Now().Local().Format("2006-01-02T15:04:05"))
			file.Close()

			newName := filepath.Join(filesDir, uniqName)
			if err := os.Rename(filename, newName); err != nil {
				os.Remove(infoPath)
				return err
			}
			if stat, err := os.Stat(newName); err == nil && stat.IsDir() {
				T.updateDirectorySize(newName, stat)
			}
			return nil
		}
		if !os.IsExist(err) {
			return err
		}
	}
	return errors.New("too many retry to create filename")
}
