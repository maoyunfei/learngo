package main

import "fmt"

func main() {
	//通过make()函数创建数组切片

	//初始元素个数为5,元素初始值为0
	mySlice1 := make([]int, 5)
	fmt.Println(mySlice1)

	//初始元素个数为5,元素初始值为0，并预留10个元素的存储空间
	mySlice2 := make([]int, 5, 10)
	fmt.Println(mySlice2)

	//初始化包含5个元素的数组切片
	mySlice3 := []int{1, 2, 3, 4, 5}
	fmt.Println(cap(mySlice3))
	fmt.Println(mySlice3)

}
