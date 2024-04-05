## phtree

[![Release](https://img.shields.io/github/release/0xnu/phtree.svg)](https://github.com/0xnu/phtree/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/0xnu/phtree)](https://goreportcard.com/report/github.com/0xnu/phtree)
[![Go Reference](https://pkg.go.dev/badge/github.com/0xnu/phtree.svg)](https://pkg.go.dev/github.com/0xnu/phtree)
[![License](https://img.shields.io/github/license/0xnu/phtree)](/LICENSE)

[PH-tree](https://en.wikipedia.org/wiki/PH-tree) (Permutation Hierarchical Tree) implementation in Go.

### Features

- Insertion and search operations for multi-dimensional keys
- Removal of key-value pairs from the tree
- Configurable maximum depth of the tree
- Thread-safe concurrent access to the tree
- Efficient handling of high-dimensional and sparse data

### Installation

Install the `phtree` package in your `Go` project by installing it with the following command:

```shell
go get github.com/0xnu/phtree
```

### Usage

Here's a basic example of how to use the `phtree` package:

```go
package main

import (
	"fmt"
	"github.com/0xnu/phtree"
)

func main() {
	tree := phtree.New(phtree.WithMaxDepth(8))

	tree.Insert([]uint64{0, 1, 2}, "apple")
	tree.Insert([]uint64{0, 1, 3}, "banana")

	value := tree.Search([]uint64{0, 1, 2})
	fmt.Println(value) // Output: "apple"

	tree.Remove([]uint64{0, 1, 3})
}
```

Check out the [examples](./examples/) directory for more detailed usage and examples.

### Testing

To run the tests for the PH-Tree package, use the following command:

```sh
go test
```

### Performance

To run the benchmarking code, you can use the following command:

```sh
go run benchmark/benchmark.go
```

### License

This project is licensed under the [MIT License](./LICENSE).

### Copyright

(c) 2024 [Finbarrs Oketunji](https://finbarrs.eu).