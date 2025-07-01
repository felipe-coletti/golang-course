package main

import (
	"fmt"
	"log"
	"math"
)

type sqrtError struct {
	Code  int
	Level string
	Err   error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("Code %d (%s): %v", se.Code, se.Level, se.Err)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		err := fmt.Errorf("valor inválido: %v (não pode ser negativo)", f)

		return 0, sqrtError{
			Code:  400,
			Level: "ERROR",
			Err:   err,
		}
	}
	
	return math.Sqrt(f), nil
}

func main() {
	_, err := sqrt(-9)

	if err != nil {
		log.Println(err)
	}
}
