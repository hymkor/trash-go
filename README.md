trash-go
========

**trash-go** is a Go library and command-line tool for moving specified files to the trash (Recycle Bin on Windows, Trash Can on Linux desktop environments).

- ✅ **Windows**: Uses the `SHFileOperationW` API in `Shell32.dll`.
- ⚠️ **Non-Windows** (experimental): Follows the [FreeDesktop.org Trash Specification 1.0][fd1] to move files to the user's "home trash".

[fd1]: https://specifications.freedesktop.org/trash-spec/trashspec-1.0.html

Installation
------------

### As a Library

```bash
go get github.com/hymkor/trash-go
```

### As a Command-Line Tool

```bash
go install github.com/hymkor/trash-go/cmd/trash@latest
```

This will install a `trash` executable in your `$GOBIN` directory.

Usage
-----

### Library

```go
package trash // import "github.com/hymkor/trash-go"

func Throw(filenames ...string) error
```

`Throw` accepts one or more file paths and moves them to the trash.

### Command-Line Tool

```bash
trash [OPTIONS] FILE...
```

Moves the specified files to the trash.

#### Options

* `-from-file FILENAME`
  Read a list of files (one per line) from the given file instead of command-line arguments.
  Use `-` as the filename to read from standard input.

#### Examples

```bash
trash *.log                 # Move matching files to trash
trash -from-file list.txt   # Move files listed in list.txt
find . -name "*.tmp" | trash -from-file -  # Use with standard input
```

Example (as library)
--------------------

See [example.go](./example.go) for a complete usage example.

```go
package main

import (
    "fmt"
    "os"
    "path/filepath"

    "github.com/hymkor/trash-go"
)

func mains(args []string) error {
    var filenames []string
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
    if err := mains(os.Args[1:]); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
```

Notes
-----

- On **Windows**, files are moved to the Recycle Bin using native shell operations.
- On **Linux**, the implementation creates `.Trash` or `.local/share/Trash` directories if they don't exist, following the FreeDesktop specification.

See also
--------

* [hymkor/trash-rs](https://github.com/hymkor/trash-rs)
  A Rust implementation with similar functionality. Available via:
  `scoop install trash` from [hymkor/scoop-bucket](https://github.com/hymkor/scoop-bucket)

Author
------

[hymkor (HAYAMA Kaoru)](https://github.com/hymkor)

License
-------

MIT License
