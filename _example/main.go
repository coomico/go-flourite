package main

import (
	"fmt"

	flourite "go-flourite"
)

func main() {
	snippet := `{
    "middlewares": ["./fixtures/middlewares/en", "./fixtures/middlewares/jp"]
  }`

	res := flourite.Detect(snippet)
	fmt.Println(res)
	fmt.Println(res.Best())
}
