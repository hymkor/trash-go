//go:build !windows

package trash

import (
	"os"
)

func throw(filenames ...string) error {
	for _, fn := range filenames {
		if err := os.Remove(fn); err != nil {
			return err
		}
	}
	return nil
}
