package main

import "fmt"

func sum(n1 int, n2 int) int {
	res := n1 + n2

	defer fmt.Println("1. n1:", n1)
	defer fmt.Println("2. n2:", n2)

	fmt.Println("3. res:", res)

	return res
}

func main() {

	res := sum(10, 20)
	fmt.Println("res:", res)
}
