trash-go
========

`trash-go` is a Go library for moving specified files to the trash (Recycle Bin/Trash Can) on Microsoft Windows.

```go doc |
package trash // import "github.com/hymkor/trash-go"

func Throw(filenames ...string) error
```

Non-Windows environments (experimental)
---------------------------------------

`trash.Throw` moves files to "the home trash" of [the FreeDesktop.org Trash specification 1.0][fd1].

[fd1]: https://specifications.freedesktop.org/trash-spec/trashspec-1.0.html

Example
-------

[./example.go](./example.go)

```example.go
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
```

See also
--------

- [hymkor/trash-rs: Move file(s) to trash-box of Microsoft Windows](https://github.com/hymkor/trash-rs)  
    A Rust implementation that can be installed via scoop install trash from [hymkor/bucket](https://github.com/hymkor/scoop-bucket)
