package basic

import "fmt"

func Condition() {
	/*
		-> if
		-> if else
		-> else if
		-> nested if
	*/

	score := 90
	var grade string

	if score <= 50 {
		grade = "E"
	} else if score <= 60 {
		grade = "D"
	} else if score <= 70 {
		grade = "C"
	} else {
		grade = "NULL"
	}

	fmt.Println(grade)
}
