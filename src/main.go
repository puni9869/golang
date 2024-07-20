package main

import "fmt"

func triangle(row int) int {
	if row == 0 {
		return 0
	}
	if row == 1 {
		return 1
	}
	if row == 2 {
		return 3
	}
	return row + triangle(row-1)
}

func main() {
	fact := triangle(5)
	fmt.Println(fact)
}
