package main

import (
	"fmt"

	flourite "go-flourite"
)

func main() {
	snippet := `<!DOCTYPE html>
  <html>
  <head>
    <title>HTML/CSS Quine</title>
    <style type="text/css">
    * { font: 10pt monospace; }
   
    head, style { display: block; }
    style { white-space: pre; }
   
    style:before {
      content:
        "<""!DOCTYPE html>"
        "A<html>A"
        "<head>A"
        "<title>""HTML/CSS Quine""</title>A"
        "<style type="text/css">";
    }
    style:after {
      content:
        "</style>A"
        "</head>A"
        "<""body></body>A"
        "</html>";
    }
    </style>
  </head>
  <body></body>
  </html>`

	detector := flourite.NewDetector()
	res := detector.Detect(snippet)
	fmt.Println(res)
	fmt.Println(res.Best())
}
