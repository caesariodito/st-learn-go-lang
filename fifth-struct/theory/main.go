package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func main() {
	// parts.CreateStructExample()

	user := User{3, "Agus", "Susanto", "agus@gmail.com", true}
	display := displayUser(user)
	fmt.Println(display)
}

func displayUser(user User) (result string) {
	result = fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
	return
}
