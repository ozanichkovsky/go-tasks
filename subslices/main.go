package main

import "fmt"

func main() {
	// These are the primes less than 200
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
		31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
		73, 79, 83, 89, 97, 101, 103, 107, 109,
		113, 127, 131, 137, 139, 149, 151, 157,
		163, 167, 173, 179, 181, 191, 193, 197, 199}
	fmt.Println(primes)

	for i := 0; i < len(primes); i++ {
		if primes[i] >= 10 {
			fmt.Println(primes[0:i])
			break
		}
	}

	var lowIndex, highIndex int

	for i := 0; i < len(primes); i++ {
		if lowIndex == 0 && primes[i]/10 >= 1 {
			lowIndex = i
		}
		if highIndex == 0 && primes[i]/10 >= 10 {
			highIndex = i
			break
		}
	}

	fmt.Println(primes[lowIndex:highIndex])
}
