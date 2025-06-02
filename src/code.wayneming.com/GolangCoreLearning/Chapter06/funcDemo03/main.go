package main

import "fmt"

func main() {
	res := func(i int, i2 int) int {
		return i + i2
	}(1, 2)
	fmt.Print(res)
}
