package main

import "fmt"

func ArrayDemo01() {
	hen1 := 3.0
	hen2 := 4.1
	hen3 := 4.2
	hen4 := 4.3
	hen5 := 4.4

	totalWeight := hen1 + hen2 + hen3 + hen4 + hen5
	fmt.Println(totalWeight)
}

func ArrayDemo02() {
	var hens [7]float64
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 6.0
	hens[3] = 2.0
	hens[4] = 22.1
	hens[5] = 50.0
	hens[6] = 114.1

	totalWeight2 := 0.0

	for i := 0; i < len(hens); i++ {
		totalWeight2 += hens[i]
	}

	fmt.Println(totalWeight2)
}

func ArrayDemo03() {
	var intArr [3]int
	fmt.Println(intArr)

	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	fmt.Println(intArr)
	fmt.Printf("intArr:%p, intArr[0]:%p, intArr[1]:%p, intArr[2]:%p\n",
		&intArr, &intArr[0], &intArr[1], &intArr[2])
}

// ArrayDemo04 初识数组
func ArrayDemo04() {
	var score [5]float64

	for i := 0; i < len(score); i++ {
		fmt.Printf("请输入第%d个元素的值:\n", i+1)
		fmt.Scanln(&score[i])
	}

	for i := 0; i < len(score); i++ {
		fmt.Printf("score[%d]=%v\n", i, score[i])
	}
}

// ArrayDemo05 数组初始化4种方式
func ArrayDemo05() {
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01:", numArr01)

	var numArr02 = [...]int{8, 9, 10}
	fmt.Println("numArr02:", numArr02)

	var numArr03 = [...]int{1: 800, 0: 900, 2: 999}
	fmt.Println("numArr03:", numArr03)

	strArr05 := [...]string{1: "tom", 0: "jack", 2: "mary"}
	fmt.Println("strArr05:", strArr05)
}

// ArrayDemo06 数组遍历
func ArrayDemo06() {
	strArr05 := [...]string{1: "tom", 0: "jack", 2: "mary"}

	// 方式一
	for i := 0; i < len(strArr05); i++ {
		fmt.Printf("index: %v, value: %v\n", i, strArr05[i])
		fmt.Printf("strArr05[%d]: %v\n", i, strArr05[i])
	}

	// 方式二
	for index, value := range strArr05 {
		fmt.Printf("index: %v, value: %v\n", index, value)
		fmt.Printf("strArr05[%d]: %v\n", index, value)
	}

	var i int = 0
	for _, value := range strArr05 {
		fmt.Printf("strArr05[%d]: %v\n", i, value)
		i++
	}
}

// ArrayExercise1 创建一个byte类型的26个元素的数组，分别放置'A'-'Z',使用for循环访问所有元素并打印出来。提示：字符数据运算'A'+1='B'
func ArrayExercise1() {
	var byteArray [26]byte
	byteArray[0] = 'A'
	for i := 1; i < len(byteArray); i++ {
		// 方式一
		// byteArray[i] = byteArray[i-1] + 1

		// 方式二
		byteArray[i] = 'A' + byte(i)
	}

	for i, v := range byteArray {
		fmt.Printf("byteArray[%d]: %c\n", i, v)
	}
}

func main() {
	//ArrayDemo01()
	//ArrayDemo02()
	//ArrayDemo03()
	//ArrayDemo04()
	//ArrayDemo05()
	//ArrayDemo06()
	ArrayExercise1()
}
