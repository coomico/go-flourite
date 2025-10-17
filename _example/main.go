package main

import (
	"fmt"

	flourite "go-flourite"
)

func main() {
	snippet := `package main

import "C"
import (
	"fmt"
	"regexp"
)

func main() {
	snippet := ""

	expr := regexp.MustCompile()
	fmt.Println(expr.MatchString(snippet))
}`

	detector := flourite.Detector{}
	res := detector.Detect(snippet)
	fmt.Println(res)
}
