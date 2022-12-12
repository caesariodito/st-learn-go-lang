package main

import "fmt"

// creating struct
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func main() {
	user := User{}
	user.ID = 1
	user.FirstName = "John"
	user.LastName = "Doe"
	user.Email = "john.doe@gmail.com"
	user.IsActive = true
	fmt.Println(user.FirstName)

	user2 := User{
		// bisa diacak
		LastName:  "Bambang",
		ID:        2,
		FirstName: "Zelda",
		IsActive:  true,
		Email:     "bambang@gmail.com",
	}
	fmt.Println(user2.FirstName)

	// gabisa diacak
	user3 := User{3, "Agus", "Susanto", "agus@gmail.com", true}
	fmt.Println(user3)
}
