package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	//声明并创建
	var personDB map[string]PersonInfo = make(map[string]PersonInfo)

	//元素赋值
	personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203"}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 101"}

	//元素查找
	id := "1234"
	person, ok := personDB[id]

	if ok {
		fmt.Println("Found person", person.Name, "with ID", id)
	} else {
		fmt.Println("Did not find person with ID ", id)
	}

	//创建并初始化
	myMap := map[string]PersonInfo{
		"1234": PersonInfo{"1234", "Lily", "Room 602"},
	}
	fmt.Println("before delete:", myMap)

	//元素删除
	delete(myMap, "123")
	fmt.Println("after delete 123:", myMap)
	delete(myMap, "1234")
	fmt.Println("after delete 1234:", myMap)

}
