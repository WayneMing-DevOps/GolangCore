/*
	变量快速入门案例
*/

package main

import "fmt"

func main() {
	var i int
	i = 10
	fmt.Println("i:", i)

	/*
		变量声明的三种方式
		1. 指定变量类型
		2. 类型推导自动判断变量数据类型
		3. 省略var，使用:=声明
	*/
	// 1. 指定变量类型
	var i1 int
	fmt.Println("i1:", i1)

	// 2. 类型推导自动判断变量数据类型
	var i2 = 10.11
	fmt.Println("i2:", i2)

	// 3. 省略var, 使用:=
	name := "WayneMing"
	fmt.Println("name:", name)

	// 4. 多声明变量
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)

	var name1, n4, to = "WayneMing", 20, 291.1
	fmt.Println(name1, n4, to)

	n11, n12, n13 := 11, 12, "W"
	fmt.Println(n11, n12, n13)
}
