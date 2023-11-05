trash-go
========

The trash-go is the library for golang to move specified files to trashbox (recycle-bin) of Microsoft Windows.

```go doc |
package trash // import "github.com/hymkor/trash-go"

func Throw(filenames ...string) error
```

Caution in NON-Windows environments
-----------------------------------

`Throw` deletes specified files as same as `os.Remove`.
Files given to `Trash` cannot be revived from the trashbox on your desktop.

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
