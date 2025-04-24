v0.3.0
======
Apr 24 2025

- Fixed a bug in the CLI where specifying a non-existent file did not result in an error.
- Added the `-from-file FILENAME` option to read filenames to remove from the specified file (use `-` to read from standard input).
- Properly check the result of `shFileOperationW`, treating `windows.ERROR_SUCCESS` as non-error.

v0.2.0
======
Nov 07 2023

- On not-MSWindows, support "the home trash" of [the FreeDesktop.org Trash 1.0][freedesktop]

[freedesktop]: https://specifications.freedesktop.org/trash-spec/trashspec-1.0.html

v0.1.0
======
Nov 06 2023

- Prototype
    - On not-MSWindows, Trash works as same as os.Remove
