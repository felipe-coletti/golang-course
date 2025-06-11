package main

import "fmt"

type Car struct {
	doorsNumber int
	color       string
}

type Pickup struct {
	Car
	fourByFour bool
}

type Sedan struct {
	Car
	deluxe bool
}

func main() {
	car1 := Pickup{
		Car: Car{
			doorsNumber: 2,
			color:       "White",
		},
		fourByFour: true,
	}

	car2 := Sedan{
		Car: Car{
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
