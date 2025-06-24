package main

import "fmt"

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

type Shape interface {
	area() float64
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	s := square{
		side: 4.5,
	}

	c := circle{
		radius: 3,
	}

	fmt.Println(s.area())
	fmt.Println(c.area())
	fmt.Println(getArea(s))
	fmt.Println(getArea(c))
}
