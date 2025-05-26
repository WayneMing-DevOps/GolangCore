package main

import "fmt"

func main() {
	// 变量快速入门案例
	var i int
	i = 10
	fmt.Println("i: ", i)

	/*
		变量声明的三种方式
		1. 指定变量类型
		2. 类型推导自动判断变量数据类型
		3. 省略var，使用:=声明
	*/
	// 1. 使用默认值
	var i1 int
	fmt.Println("i1: ", i1)

	// 2. 类型推导自动判断变量数据类型
	var i2 = 10.11
	fmt.Println("i2: ", i2)
}
