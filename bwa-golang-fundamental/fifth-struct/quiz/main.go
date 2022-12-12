package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
}

type Group struct {
	ID    int
	Name  string
	Admin User
	Users []User
}

func (group Group) Display() {
	fmt.Printf("Name: %s\n", group.Name)
	fmt.Printf("Admin: %s %s \n", group.Admin.FirstName, group.Admin.LastName)
	fmt.Println("Users: ")
	for _, user := range group.Users {
		fmt.Printf("%s %s \n", user.FirstName, user.FirstName)
	}
}

func main() {
	// Mengubah displayGroup yang tadinya function menjadi method (Group)
	user1 := User{1, "Bambang", "Salim"}
	user2 := User{2, "Sugiyono", "Salim"}
	user3 := User{3, "Sutarjo", "Salim"}

	group1 := Group{1, "Gamer", user1, []User{user1, user2, user3}}

	group1.Display()
}
