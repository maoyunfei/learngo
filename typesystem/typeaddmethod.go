package main

import "fmt"

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

//只有在需要修改对象的时候才必须用指针
func (a *Integer) Add(b Integer) {
	*a += b
}

func main() {
	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, "Less 2")
	}

	a.Add(2)
	fmt.Println(a)
}
