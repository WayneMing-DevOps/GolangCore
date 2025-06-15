package main

import "fmt"

func SliceDemo01() {
	var intArr [5]int = [...]int{1, 22, 33, 55, 66}
	slice := intArr[1 : 3+1]
	fmt.Println("slice:", slice)
	fmt.Println("slice length:", len(slice))
	fmt.Println("slice add:", cap(slice))
}

func SliceDemo2() {
	var slice []float64 = make([]float64, 5, 10)
	slice[1] = 10
	slice[3] = 20
	fmt.Println(slice)
	fmt.Println("slice length:", len(slice))
	fmt.Println("slice cap:", cap(slice))
}

func SliceDemo3() {
	var strSlice []string = []string{"tom", "jack", "mary"}
	fmt.Println("strSlice:", strSlice)
	fmt.Println("strSlice size:", len(strSlice))
	fmt.Println("strSlice cap:", cap(strSlice))
}

func main() {
	//SliceDemo01()
	//SliceDemo2()
	SliceDemo3()
}
