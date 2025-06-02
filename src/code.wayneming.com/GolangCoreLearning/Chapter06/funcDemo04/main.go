package main

import (
	"fmt"
	"strings"
)

// 累加器
func AddUpper() func(int) int {
	var n int = 0
	return func(x int) int {
		n += x
		return n
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(fileName string) string {
		if !strings.HasSuffix(fileName, suffix) {
			return fileName + suffix
		}

		return fileName
	}
}

func init() {
	fmt.Println("init()....")
}

func main() {
	fmt.Println("main()....")
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(4))
	fmt.Println(f(8))

	f2 := makeSuffix(".jpg")
	fmt.Println(f2("a.jpg"))
	fmt.Println(f2("a"))
}
