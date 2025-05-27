/*
数据类型
*/
package main

import (
	"fmt"
	"unsafe"
)

// 整数类型
func type_int() {
	var i int = 3
	fmt.Println("i:", i)
	fmt.Printf("i 的数据类型:%T,占用的字节数:%d\n", i, unsafe.Sizeof(i))

	var v2 byte = 255
	fmt.Println("b2:", v2)

	var v3 = 1
	fmt.Printf("i 的数据类型:%T,占用的字节数:%d\n", v3, unsafe.Sizeof(v3))
}

// 小数类型/浮点型
func type_float() {
	var price float32 = 89.12
	fmt.Println("pricd:", price)

	var v2 = 89.12
	fmt.Printf("v2 的数据类型:%T,占用的字节数:%d\n", v2, unsafe.Sizeof(v2))

	var num1 float32 = -0.00089
	var num2 float64 = -7809656.09
	fmt.Println(num1, num2)

	var num3 float32 = -123.0000901 // 会造成精度丢失
	var num4 float64 = -123.0000901
	fmt.Println(num3, num4)
}

// 字符类型
func type_char() {
	var c1 byte = 'a'
	var c2 = 'c'
	fmt.Println(c1, c2)                    // 默认输出ASCII对应的数字
	fmt.Printf("c1: %c\nc2: %c\n", c1, c2) // 输出ASCII对应字符
	fmt.Printf("c1 的数据类型:%T,占用的字节数:%d\n", c1, unsafe.Sizeof(c1))

	// var c3 byte = '啊'
	var c3 int = '啊'
	fmt.Printf("c3: %c, %d\n", c3, c3)
	fmt.Printf("c3 的数据类型:%T,占用的字节数:%d\n", c3, unsafe.Sizeof(c3))

	var c4 int = 444
	var c5 int = 111
	var res int = c4 + c5
	fmt.Printf("res: %c, %d, %c\n", res, res, c5)
}

// 布尔类型
func type_bool() {
	var b1 bool = false
	var b2 bool = true
	fmt.Println(b1, b2)
	fmt.Printf("b1 的数据类型:%T,占用的字节数:%d\n", b1, unsafe.Sizeof(b1))
}

// string类型
func type_string() {
	var address string = "在Unicode字符串中不可能由码点数量决定显示它所需要的长度"
	fmt.Println(address)

	var note string = `总体来说，在Unicode
					   符串中不可能由码点数
					   量决定显示它所需要的长度，
						或者显示字符串之后在文本缓冲区
						中光标应该放置的位置；组合字符、
						变宽字体、不可打印字符和从右至左的
						文字都是其归因。
						`
	fmt.Println(note)

	var str string = "Hcc" + "aa"
	str += "112a"
	fmt.Println(str)
}

// 基本数据类型默认值
func type_DefaultValue() {
	var a int     // 0
	var b float32 // 0
	var c float64 // 0
	var d bool    // false
	var e string  // ""
	fmt.Println(a, b, c, d, e)
}

func main() {
	type_int()
	type_float()
	type_char()
	type_bool()
	type_string()
	type_DefaultValue()
}
