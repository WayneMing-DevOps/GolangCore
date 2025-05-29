/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
你可以按任意顺序返回答案
*/

package main

import "fmt"

// 自己写的, Runtime:23ms
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

// 别人写的, Runtime:0ms
func twoSum2(nums []int, target int) []int {
	// 创建哈希表存储元素值到索引的映射
	numMap := make(map[int]int)

	// 遍历数组
	for i, num := range nums {
		// 计算需要的补数
		complement := target - num

		// 检查补数是否已在哈希表中
		if idx, found := numMap[complement]; found {
			return []int{idx, i} // 返回找到的索引对
		}

		// 将当前元素存入哈希表
		numMap[num] = i
	}

	return nil // 无解时返回nil（题目保证有解）
}

func main() {
	var numArr = []int{2, 15, 11, 7}
	var target int = 22
	res := twoSum1(numArr, target)
	res = twoSum2(numArr, target)
	fmt.Println(res)
}
