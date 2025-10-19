package main

import (
	"fmt"

	flourite "go-flourite"
)

func main() {
	snippet := `/**
  * Improve readability when focused and also mouse hovered in all browsers.
  */
 
 a:active,
 a:hover {
   outline: 0;
 }
 
 /**
  * Address styling not present in IE 8/9/10/11, Safari, and Chrome.
  */
 
 abbr[title] {
   border-bottom: 1px dotted;
 }`

	detector := flourite.NewDetector()
	res := detector.Detect(snippet)
	fmt.Println(res)
	fmt.Println(res.Best())
}
