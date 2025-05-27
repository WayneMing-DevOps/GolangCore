/*
数据类型转换
*/
package main

import "fmt"

func main() {
	var i int32 = 100
	// i => float
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)
	fmt.Println(n1, n2, n3)

	var n4 int32 = 12
	var n5 int64
	var n6 int8
	n5 = int64(n4) + 20
	// n6 = n4 + 20	// 错误，不允许这么计算，在go中必须转换！
	fmt.Println(n4, n5, n6)

}
