package main

import "fmt"

func main() {
	oldSlice := make([]int, 5, 10)
	fmt.Println(oldSlice)

	oldSlice[0] = 1
	oldSlice[1] = 2
	oldSlice[2] = 3
	oldSlice[3] = 4
	fmt.Println(oldSlice)

	//基于数组切片创建数组切片
	newSlice := oldSlice[:6]
	fmt.Println(newSlice)
}
