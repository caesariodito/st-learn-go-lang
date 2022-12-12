package main

import (
	"fmt"
	"try/models"
)

func main() {
	user := models.User{
		ID:        1,
		FirstName: "Anjay",
		LastName:  "Mabar",
	}
	fmt.Println(user)
}
