/*
给你一个字符串 date，它的格式为 yyyy-mm-dd，表示一个公历日期。
date 可以重写为二进制表示，只需要将年、月、日分别转换为对应的二进制表示（不带前导零）并遵循 year-month-day 的格式。
返回 date 的 二进制 表示。
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertDateToBinary(date string) string {
	var year int64
	var month int64
	var day int64

	res := strings.Split(date, "-")
	year, _ = strconv.ParseInt(res[0], 10, 64)
	month, _ = strconv.ParseInt(res[1], 10, 64)
	day, _ = strconv.ParseInt(res[2], 10, 64)

	year2 := strconv.FormatInt(year, 2)
	month2 := strconv.FormatInt(month, 2)
	day2 := strconv.FormatInt(day, 2)

	return year2 + "-" + month2 + "-" + day2
}

func convertDateToBinary2(date string) string {
	parts := strings.Split(date, "-")
	if len(parts) != 3 {
		return ""
	}

	binary := make([]string, 3)
	for i, part := range parts {
		num, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			return ""
		}
		binary[i] = strconv.FormatInt(num, 2)
	}

	return strings.Join(binary, "-")
}

func main() {
	var date string = "1900-01-01"
	res := convertDateToBinary2(date)
	fmt.Printf(res)
}
