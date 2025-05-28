/*
数据类型转换
*/
package main

import (
	"fmt"
	"unsafe"
	"strconv"
)

// 基本数据类型转换string
func conversion_String() {
	var num1 int = 99
	// var num2 float64 = 11.11
	// var b bool = true
	// var myChar byte = 'A'
	var str1 string
	
	// 方式一：使用fmt.Sprintf()函数
	str1 = fmt.Sprintf("%d", num1)
	fmt.Printf("str1 的数据类型:%T,占用的字节数:%d\n", str1, unsafe.Sizeof(str1))
	
	// 方式二：使用strconv.Format函数
	str1 = strconv.FormatInt(int64(num1), 10)
	fmt.Printf("str1 的数据类型:%T,占用的字节数:%d\n", str1, unsafe.Sizeof(str1))
}

// string转换基本数据类型
func stringConversionBase() {
	var str string = "12345678"
	var n1 int64
	n1, _ = strconv.ParseInt(str, 10, 64)
	fmt.Printf("n1 的数据类型:%T,占用的字节数:%d\n", n1, unsafe.Sizeof(n1))
	
}

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

	conversion_String()
	stringConversionBase()
}
