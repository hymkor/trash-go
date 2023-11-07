//go:build !windows

package trash

import (
	"github.com/hymkor/trash-go/internal/freedesktop"
)

func throw(filenames ...string) error {
	homeTrash, err := freedesktop.HomeTrash(true)
	if err != nil {
		return err
	}
	for _, fn := range filenames {
		if err := homeTrash.Throw(fn); err != nil {
			return err
		}
	}
	return nil
}
