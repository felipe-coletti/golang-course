package main

import "fmt"

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

func (square Square) area() float64 {
	return square.side * square.side
}

func (circle Circle) area() float64 {
	return 3.14 * circle.radius * circle.radius
}

type Shape interface {
	area() float64
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	square := Square{
		side: 4.5,
	}

	circle := Circle{
		radius: 3,
	}

	fmt.Println(square.area())
	fmt.Println(circle.area())
	fmt.Println(getArea(square))
	fmt.Println(getArea(circle))
}
