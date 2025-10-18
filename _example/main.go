package main

import (
	"fmt"

	flourite "go-flourite"
)

func main() {
	snippet := `#include <stdio.h>`

	detector := flourite.NewDetector()
	res := detector.Detect(snippet)
	fmt.Println(res)
	fmt.Println(res.Best())
}
