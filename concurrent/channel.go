package main

import "fmt"

func Count2(ch chan int) {
	fmt.Println("Counting")
	ch <- 1
}

func main() {
	chs := make([] chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count2(chs[i])
	}

	for _, ch := range chs {
		<-ch
	}
}
