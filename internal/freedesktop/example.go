//go:build run

package main

import (
	"os"

	"github.com/hymkor/trash-go/internal/freedesktop"
)

func main() {
	tmpTrash := freedesktop.TmpTrash()

	for _, fn := range os.Args[1:] {
		if err := tmpTrash.Throw(fn); err != nil {
			println(fn, err.Error())
		}
	}
}
