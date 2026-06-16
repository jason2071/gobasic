package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	r := Rectangle{Width: 3, Height: 4}
	r.Area()      // 12
	r.Perimeter() // 14

	fmt.Println(r.Area())
	fmt.Println(r.Perimeter())
}
