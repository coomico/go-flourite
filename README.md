# Go Flourite [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE) ![CI](https://github.com/coomico/go-flourite/workflows/CI/badge.svg)

A Go port of TypeScript [Flourite](https://github.com/teknologi-umum/flourite) project (a programming language detector that identifies the language of a given code string).

Just like the original project, this is a pure Go implementation with no external dependencies that is capable of detecting [various languages](https://github.com/teknologi-umum/flourite?tab=readme-ov-file#detectable-languages).

## Installation

```bash
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

## Credits

This is a Go port of [teknologi-umum/flourite](https://github.com/teknologi-umum/flourite), which itself is a fork of [ts95/lang-detector](https://github.com/ts95/lang-detector).

## Support

Contributions are welcome.

## License

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
