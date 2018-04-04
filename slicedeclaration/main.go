package main

import "fmt"

func main() {
	var i = make([]int, 5)
	var f = make([]float64, 9)
	var s = make([]string, 4)

	fmt.Println(len(i), len(f), len(s))
}
