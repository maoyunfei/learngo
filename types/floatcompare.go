package main

import (
	"fmt"
	"math"
)

func IsEqual(f1, f2, p float64) bool {
	return math.Dim(f1, f2) < p
}

func main() {
	fmt.Println(IsEqual(1.1, 1.0, 0.01))
	fmt.Println(IsEqual(1.001, 1.002, 0.01))
}
