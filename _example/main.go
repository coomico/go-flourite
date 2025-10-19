package main

import (
	"fmt"

	flourite "go-flourite"
)

func main() {
	snippet := `defmodule Sort do
  def qsort([]), do: []
  def qsort([h | t]) do
    {lesser, greater} = Enum.split_with(t, &(&1 < h))
    qsort(lesser) ++ [h] ++ qsort(greater)
  end
end`

	detector := flourite.NewDetector()
	res := detector.Detect(snippet)
	fmt.Println(res)
	fmt.Println(res.Best())
}
