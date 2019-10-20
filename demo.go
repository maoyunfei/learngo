package main

import (
	"fmt"
)

func process(uid int, result chan int) {
	result <- uid
	return
}

func main() {
	uids := []int{5, 6, 7}
	channels := make(map[int]chan int, len(uids))
	for _, uid := range uids {
		channels[uid] = make(chan int)
		go process(uid, channels[uid])
	}

	result := map[int]int{}
	for _, uid := range uids {
		result[uid] = <-channels[uid]
	}
	fmt.Println("result", result)

}