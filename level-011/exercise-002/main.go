package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)

	if err != nil {
		return nil, fmt.Errorf("falha ao converter para JSON: %w", err)
	}

	return bs, nil
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)

	if err != nil {
		log.Fatalf("Erro: %v", err)
	}

	fmt.Println(string(bs))

}
