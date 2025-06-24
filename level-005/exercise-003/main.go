package main

import "fmt"

type car struct {
	doorsNumber int
	color       string
}

type pickup struct {
	car
	fourByFour bool
}

type sedan struct {
	car
	deluxe bool
}

func main() {
	car1 := pickup{
		car: car{
			doorsNumber: 2,
			color:       "White",
		},
		fourByFour: true,
	}

	car2 := sedan{
		car: car{
			doorsNumber: 4,
			color:       "Black",
		},
		deluxe: true,
	}

	fmt.Printf("Pickup: %+v\n", car1)
	fmt.Printf("Sedan: %+v\n", car2)
	fmt.Println("Pickup is 4x4:", car1.fourByFour)
	fmt.Println("Sedan is deluxe:", car2.deluxe)
}
