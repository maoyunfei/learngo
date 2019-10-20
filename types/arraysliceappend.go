package main

import "fmt"

func main() {
	mySlice := make([]int, 5, 10)
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))
	fmt.Println(mySlice)

	mySlice = append(mySlice, 1, 2, 3)
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))
	fmt.Println(mySlice)

	mySlice2 := []int{8, 9, 10}

	mySlice = append(mySlice, mySlice2...)
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))
	fmt.Println(mySlice)
}
