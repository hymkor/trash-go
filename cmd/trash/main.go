package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/hymkor/trash-go"
)

var flagFromFile = flag.String("from-file", "", "Read the file list to remove from a file")

func removeFromFile(r io.Reader) error {
	sc := bufio.NewScanner(r)
	list := make([]string, 0, 16)
	for sc.Scan() {
		list = append(list, sc.Text())
	}
	err := sc.Err()
	if err != nil {
		return err
	}
	return trash.Throw(list...)
}

func mains(args []string) error {
	if *flagFromFile == "-" {
		return removeFromFile(os.Stdin)
	} else if *flagFromFile != "" {
		fd, err := os.Open(*flagFromFile)
		if err != nil {
			return err
		}
		err = removeFromFile(fd)
		fd.Close()
		return err
	}

	filenames := make([]string, 0, len(args))
	for _, arg := range args {
		if matches, err := filepath.Glob(arg); err != nil {
			filenames = append(filenames, arg)
		} else if len(matches) == 0 {
			return fmt.Errorf("%s: %w", arg, os.ErrNotExist)
		} else {
			filenames = append(filenames, matches...)
		}
	}
	return trash.Throw(filenames...)
}

func main() {
	flag.Parse()
	if err := mains(flag.Args()); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
