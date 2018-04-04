package main

import "fmt"

func main() {
	var s = []string{"City", "Country", "World"}

	for _, str := range s {
		fmt.Printf(" Hello " + str + "\n")
	}
}
