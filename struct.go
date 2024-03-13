package main

import (
	"fmt"
)

type Circle struct {
	x float64
	y float64
	r float64
}

type Rectangle struct {
	w float64
	h float64
}

type Shape interface {
	area(a int) float64
}

func circleArea(c *Circle) float64 {
	return 3.14 * c.r * c.r
}

func (c *Circle) area() float64 {
	return 3.14 * c.r * c.r
}

func (r *Rectangle) area() float64 {
	return r.w * r.h
}

func sumArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}

	return area
}

func main() {
	c := new(Circle)
	c.x = 1
	c.y = 2
	c.r = 5
	c2 := Circle{3, 2, 5}

	r := Rectangle{
		w: 1,
		h: 2,
	}

	fmt.Println(*c, c2)
	fmt.Println(c, &c2)
	fmt.Println(circleArea(c))
	fmt.Println((*c).area())
	fmt.Println(c2.area())
	fmt.Println(sumArea(c, &r))
}