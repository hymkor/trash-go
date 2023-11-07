package freedesktop

import (
	"os"
	"path/filepath"
)

// HomeTrash is the trashcan using XDG_DATA_HOME/Trash
func HomeTrash(create bool) (*TrashCan, error) {
	if xdgDataHome := os.Getenv("XDG_DATA_HOME"); xdgDataHome != "" {
		return &TrashCan{topDir: filepath.Join(xdgDataHome, "Trash"), create: create}, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	return &TrashCan{topDir: filepath.Join(home, ".local", "share", "Trash"), create: create}, nil
}

// TmpTrash is the trashcan for test
func TmpTrash() *TrashCan {
	return &TrashCan{topDir: os.TempDir()}
}
