package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id       int
	Username string
}

func main() {
	u1 := User{0, "silva_ana"}
	u2 := User{1, "santos_bruno"}
	u3 := User{2, "oliveira_carla"}

	users := []User{u1, u2, u3}

	fmt.Println(users)

	u, err := json.Marshal(users)

	if err != nil {
		fmt.Println("Erro ao fazer marshal:", err)
		return
	}

	fmt.Println(string(u))
}
