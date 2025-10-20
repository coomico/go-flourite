package main

import (
	"fmt"

	flourite "go-flourite"
)

func main() {
	snippet := `const length: number = 100

   const numberArrays: number[] = Array.from({length: length}, (_, i: number) => i)

    numberArrays.forEach((item: number): void => {
        if (item % 15 === 0) console.log('FizzBuzz' as string)
        else if(item % 3 === 0) console.log('Fizz' as string)
        else if (item % 5 === 0) console.log('Buzz' as string)
    })`

	res := flourite.Detect(snippet)
	fmt.Println(res)
	fmt.Println(res.Best())
}
