package main

import (
	"fmt"

	flourite "go-flourite"
)

func main() {
	snippet := "```bash\nmake build-debug```"

	res := flourite.Detect(snippet)
	fmt.Println(res)
	fmt.Println(res.Best())
}
