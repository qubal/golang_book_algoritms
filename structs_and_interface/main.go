package main

import (
	"fmt"
	"math"
)

//Circle structure for circle area
type Circle struct {
	x, y, r float64
}

//Rectangle structure for rectangle area
type Rectangle struct {
	x1, x2, y1, y2 float64
}

//Shape areas
type Shape interface {
	area() float64
	perimeter() float64
}

func main() {

	c := Circle{x: 1, y: 1, r: 5}
	fmt.Println(c.area())

	r := Rectangle{0, 10, 0, 10}
	fmt.Println(r.area())

	fmt.Println(totalArea(&c, &r))
	fmt.Println(totalPerimeter(&r))
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (l + w)
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func totalPerimeter(shapes ...Shape) float64 {
	var perimeter float64
	for _, s := range shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}
