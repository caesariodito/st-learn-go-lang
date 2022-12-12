package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func main() {
	// parts.CreateStructExample()

	user2 := User{2, "Agus", "Susanto", "agus@gmail.com", true}
	user3 := User{3, "Bambang", "Salim", "bambang@gmail.com", true}
	display := displayUser(user2)
	fmt.Println(display)

	fmt.Println("===")

	// users = []User{user2, user3}
	// group := Group{"Gamer", user2, users, true}
	group := Group{"Gamer", user2, []User{user2, user3}, true}
	displayGroup(group)
}

func displayUser(user User) (result string) {
	result = fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
	return
}

func displayGroup(group Group) {
	fmt.Printf("Name: %s", group.Name)
	fmt.Println()
	fmt.Printf("Member count: %d", len(group.Users))
	fmt.Println("\nUser's name:")

	for _, user := range group.Users {
		fmt.Printf("%s %s", user.FirstName, user.LastName)
		fmt.Println()
	}
}
