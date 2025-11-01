# Go Flourite [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE) ![CI](https://github.com/coomico/go-flourite/workflows/CI/badge.svg)

A Go port of TypeScript [Flourite](https://github.com/teknologi-umum/flourite) project (a programming language detector that identifies the language of a given code string).

Just like the original project, this is a pure Go implementation with no external dependencies that is capable of detecting [various languages](https://github.com/teknologi-umum/flourite?tab=readme-ov-file#detectable-languages).

## Installation

```sh
go get github.com/coomico/go-flourite
```

## Usage

Then you can do something like this:
```go
package main

import (
	"fmt"

	"github.com/coomico/go-flourite"
)

func main() {
	snippet := `fmt.Println("Hello World!")`
	res := flourite.Detect(snippet).Best()
	fmt.Println(res.Language) // Go
}
```

Or you can specify a detection strategy:
```go
strategy := flourite.Strategy{
	IsUnknown: true,
	Heuristic: true,
}
res := strategy.Detect(snippet)
// ...
```

## Performance
All benchmark samples can be found in [testdata/lines](https://github.com/coomico/go-flourite/tree/main/testdata/lines).
```
BenchmarkMultiline/100000_lines-2         1	8146513633 ns/op	14018120 B/op	     36 allocs/op
BenchmarkMultiline/10000_lines-2          2	 883795094 ns/op	 1371616 B/op	     25 allocs/op
BenchmarkMultiline/1000_lines-2           4	 318228596 ns/op	  120690 B/op	     11 allocs/op
BenchmarkMultiline/100_lines-2           25	  46593327 ns/op	    9119 B/op	      5 allocs/op
BenchmarkMultiline/10_lines-2           310	   3853172 ns/op	     889 B/op	      3 allocs/op
```

## Credits

This is a Go port of [teknologi-umum/flourite](https://github.com/teknologi-umum/flourite), which itself is a fork of [ts95/lang-detector](https://github.com/ts95/lang-detector).

## Support

Contributions are welcome.

## License

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
