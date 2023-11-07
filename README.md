trash-go
========

The trash-go is the library for golang to move specified files to trashbox (recycle-bin/trash-can) of Microsoft Windows.

```go doc |
package trash // import "github.com/hymkor/trash-go"

func Throw(filenames ...string) error
```

in NON-Windows environments (experimental)
------------------------------------------

`trash.Throw` moves files to "the home trash" of [the FreeDesktop.org Trash specification 1.0][fd1].

[fd1]: https://specifications.freedesktop.org/trash-spec/trashspec-1.0.html

Sample
------

[cmd/trash/main.go](cmd/trash/main.go)

```cmd/trash/main.go
package main

import (
    "fmt"
    "os"
    "path/filepath"

    "github.com/hymkor/trash-go"
)

func main() {
    args := os.Args[1:]
    if len(args) > 0 {
        filenames := make([]string, 0, len(args))
        for _, arg := range args {
            if matches, err := filepath.Glob(arg); err != nil {
                filenames = append(filenames, arg)
            } else {
                filenames = append(filenames, matches...)
            }
        }
        err := trash.Throw(filenames...)
        if err != nil {
            fmt.Fprintln(os.Stderr, err.Error())
            os.Exit(1)
        }
    }
}
```
