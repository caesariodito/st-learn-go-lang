package calculation

import "fmt"

func Add(number int, number2 int) int {
	return number + number2
}

func Testing() string{
	return fmt.Sprint("Testing function private:", private("Dito"))
}

func private(name string) string {
	return fmt.Sprintf("Suka %s", name)
}

/*
-> nama function dengan huruf besar menandakan public. ex: "Add" not "add".
*/