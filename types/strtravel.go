package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Hello, 世界"
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i]
		fmt.Println(i, ch)
	}

	fmt.Println("---------------")
	fmt.Println(utf8.RuneCountInString(str))

	//以Unicode字符方式遍历
	//UTF-8中每个中文字符占据3个字节
	for i, ch := range str {
		fmt.Println(i, string(ch))
	}

}
