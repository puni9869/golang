package main

import "fmt"

func recur(token int) {
	if token < 1 {
		return
	}
	fmt.Println(token, token%2 == 0)
	token--
	recur(token)
}
func main() {
	recur(5)
}
