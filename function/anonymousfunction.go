package main

import "fmt"

func main() {
	//匿名函数
	func(a, b int) {
		fmt.Println(a + b)
	}(1, 2)

	ret := func(a, b int) int {
		return a + b
	}(1, 2)
	fmt.Println(ret)
}
