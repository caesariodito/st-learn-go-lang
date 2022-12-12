package models

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

func (user User) Display() string {
	return fmt.Sprintf("%s %s", user.FirstName, user.LastName)
}

func (group Group) Display() {
	fmt.Printf("Name: %s\n", group.Name)
	fmt.Printf("Admin: %s %s \n", group.Admin.FirstName, group.Admin.LastName)
	fmt.Println("Users: ")
	for _, user := range group.Users {
		fmt.Printf("%s %s \n", user.FirstName, user.FirstName)
	}
}
