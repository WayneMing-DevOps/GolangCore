package funcDemo05

import "fmt"

func Demo01(n int) {
	if n > 2 {
		n--
		Demo01(n)
	} else {
		fmt.Println(n)
	}
}

func Demo02(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else if n >= 2 {
		return Demo02(n-1) + Demo02(n-2)
	}
	return 0
}
