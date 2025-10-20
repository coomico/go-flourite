package main

import (
	"fmt"

	flourite "go-flourite"
)

func main() {
	snippet := `SELECT round ( EXP ( SUM (ln ( ( 1 + SQRT( 5 ) ) / 2)
  ) OVER ( ORDER BY level ) ) / SQRT( 5 ) ) fibo
FROM dual
CONNECT BY level <= 10;`

	res := flourite.Detect(snippet)
	fmt.Println(res)
	fmt.Println(res.Best())
}
