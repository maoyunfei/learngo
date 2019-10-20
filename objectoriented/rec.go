package objectoriented

import "fmt"

type Rec struct {
	x      float64
	y      float64
	width  float64
	height float64
}

func NewRec(x, y, width, height float64) *Rec {
	return &Rec{x, y, width, height}
}

func (r Rec) Area() float64 {
	return r.width * r.height
}

func main() {
	r1 := Rec{0, 0, 200, 200}
	fmt.Println(r1.Area())

	r2 := NewRec(0, 0, 100, 100)
	fmt.Println(r2.Area())
}
