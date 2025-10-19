package main

import (
	"fmt"

	flourite "go-flourite"
)

func main() {
	snippet := `Future<void> main() async {
  checkVersion();
  print('In main: version is \${await lookUpVersion()}');
}`

	detector := flourite.NewDetector()
	res := detector.Detect(snippet)
	fmt.Println(res)
	fmt.Println(res.Best())
}
