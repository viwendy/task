package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.14
}

func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

func (c Circle) Perimeter() float64 {
	return 2 * c.Radius * 3.14
}

func main() {
	r := Rectangle{10, 5}
	c := Circle{5}
	var sr Shape = r
	var sc Shape = c
	fmt.Println(sr.Area())
	fmt.Println(sr.Perimeter())
	fmt.Println(sc.Area())
	fmt.Println(sc.Perimeter())
}
