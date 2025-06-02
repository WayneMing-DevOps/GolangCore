package main

import "fmt"

var age int = 10
var Name string = "Wayne"

func test() {
	age := 11
	Name := "Tom"
	fmt.Println("test()", age)
	fmt.Println("test()", Name)
}

func main() {
	test()
}
