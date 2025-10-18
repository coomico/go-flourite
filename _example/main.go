package main

import (
	"fmt"

	flourite "go-flourite"
)

func main() {
	snippet := `#!/bin/bash
echo "bash"`

	detector := flourite.NewDetector()
	res := detector.Detect(snippet)
	fmt.Println(res)
}
