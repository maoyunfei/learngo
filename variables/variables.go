package main

import "fmt"

func GetName() (firstName, lastName, nickName string) {
	return "May", "Chan", "Chibi"
}

func main() {
	_, _, nickName := GetName()
	fmt.Printf(nickName)
}
