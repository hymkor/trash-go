//go:build example

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hymkor/trash-go"
)

func mains(args []string) error {
	filenames := make([]string, 0, len(args))
	for _, arg := range args {
		if matches, err := filepath.Glob(arg); err != nil {
			filenames = append(filenames, arg)
		} else if len(matches) < 1 {
			return fmt.Errorf("%s: %w", arg, os.ErrNotExist)
		} else {
			filenames = append(filenames, matches...)
		}
	}
	return trash.Throw(filenames...)
}

func main() {
	if err := mains(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
