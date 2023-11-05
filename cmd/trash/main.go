package main

import (
	"fmt"
	"os"

	"github.com/hymkor/trash-go"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		err := trash.Throw(args...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	}
}
