package main

import "fmt"

//数组是值类型，
//在函数体中无法修改传入的数组的内容，因为函数内操作的只是所传入数据的一个副本
func modify(array [5]int) {
	array[0] = 10;
	fmt.Println("In modify(), array values:", array)
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	modify(array)
	fmt.Println("In main(), array values:", array)
}
