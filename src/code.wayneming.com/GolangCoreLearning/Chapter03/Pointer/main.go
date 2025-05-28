package main

import "fmt"

func test_Pointer1() {
	var i int = 10
	fmt.Println("i:", i)
	//fmt.Printf("i变量的内存地址:%v\n", i)
	fmt.Println("i指向的内存地址 &i:", &i)

	var ptr *int = &i
	fmt.Printf("ptr指向的值:%v\n", *ptr)
	fmt.Printf("ptr变量内存地址:%v\n", &ptr)
	fmt.Printf("ptr指向的内存地址:%v\n", ptr)
}

func test_Pointer2() {
	var num int = 100
	fmt.Printf("num:%d\n", num)
	fmt.Printf("num内存地址:%v\n", &num)

	var ptr *int = &num
	*ptr = 102
	fmt.Printf("num:%d\n", num)
	fmt.Printf("ptr:%v\n", *ptr)
	fmt.Printf("ptr内存地址:%v\n", &ptr)
	fmt.Printf("ptr指向内存地址:%v\n", ptr)
}

func main() {
	// test_Pointer1()
	test_Pointer2()
}
