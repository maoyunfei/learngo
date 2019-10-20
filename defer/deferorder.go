package main

import "fmt"

func main() {
	//defer按照先进后出的顺序执行
	defer fmt.Println(1)
	defer fmt.Println(2)
}
