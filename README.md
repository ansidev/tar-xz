# tar-xz

[![Go Reference](https://pkg.go.dev/badge/github.com/ansidev/tar-xz.svg)](https://pkg.go.dev/github.com/ansidev/tar-xz)
[![Go CI](https://github.com/ansidev/tar-xz/actions/workflows/ci.yml/badge.svg)](https://github.com/ansidev/tar-xz/actions/workflows/ci.yml)

## Introduction

A simple Go module for decompressing a tar.xz file.

## Supported platforms

- [x] Linux
- [x] macOS

## Example

```go
package main

import (
	tarXz "github.com/ansidev/tar-xz"
)

func main() {
	err := tarXz.Decompress("archive.tar.xz", "/path/to/output/dir")
	if err != nil {
		log.Fatal(err)
	}
}
```

## Contact

Le Minh Tri [@ansidev](https://ansidev.xyz/about).

## License

This source code is available under the [AGPL-3.0 LICENSE](/LICENSE).
