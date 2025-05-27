/*
2894. 分类求和并作差
给你两个正整数 n 和 m 。

现定义两个整数 num1 和 num2 ，如下所示：
	- num1：范围 [1, n] 内所有 无法被 m 整除 的整数之和。
	- num2：范围 [1, n] 内所有 能够被 m 整除 的整数之和。
返回整数 num1 - num2 。
*/

package main

import "fmt"

func differenceOfSums(n int, m int) int {
	var num1, num2, temp int

	for i := 1; i <= n; i++ {
		temp = i % m

		if temp != 0 {
			// num1：范围 [1, n] 内所有 无法被 m 整除 的整数之和。
			num1 += i
		} else {
			// num2：范围 [1, n] 内所有 能够被 m 整除 的整数之和。
			num2 += i
		}
	}

	return num1 - num2
}

func main() {
	var res int = differenceOfSums(5, 1)
	fmt.Println(res)
}
