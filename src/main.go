package main

import "fmt"

func factorial(token int) int {
	if token < 2 {
		return 1
	}
	return token * factorial(token-1)
}

func main() {
	fact := factorial(5)
	fmt.Println(fact)
}
