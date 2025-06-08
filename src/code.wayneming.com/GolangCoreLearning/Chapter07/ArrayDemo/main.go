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

func main() {
	ArrayDemo01()
	ArrayDemo02()
	ArrayDemo03()
	//ArrayDemo04()
	ArrayDemo05()
}
