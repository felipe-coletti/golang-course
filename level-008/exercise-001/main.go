package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Username string
	Age      int
}

func main() {
	u1 := User{"silva_ana", 28}
	u2 := User{"santos_bruno", 34}
	u3 := User{"oliveira_carla", 31}

	users := []User{u1, u2, u3}

	fmt.Println(users)

	u, err := json.Marshal(users)

	if err != nil {
		fmt.Println("Erro ao fazer marshal:", err)
		return
	}

	fmt.Println(string(u))
}
