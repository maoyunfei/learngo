package main

import "fmt"

func main() {
	myArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mySlice := myArray[5:]
	for i, v := range mySlice {
		fmt.Println(i, v)
	}
}
