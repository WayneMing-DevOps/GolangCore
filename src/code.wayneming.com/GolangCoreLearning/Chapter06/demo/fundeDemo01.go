package demo

import "fmt"

func Demo01(n int) {
	if n > 2 {
		n--
		Demo01(n)
	} else {
		fmt.Println(n)
	}
}
