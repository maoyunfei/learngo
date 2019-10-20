package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	fmt.Println("before copy:")
	fmt.Println(slice1)
	fmt.Println(slice2)
	copy(slice1, slice2)
	fmt.Println("after copy:")
	fmt.Println(slice1)
	fmt.Println(slice2)
}
