package main

import "fmt"

func main() {

	title := "Golang the best language"

	// quiz 1
	/*
		cetak char yang hanya mengandung index genap
	*/
	fmt.Println("quiz 1")
	for index, c := range title {
		if index%2 == 0 {
			fmt.Println("index :", index, ", char :", string(c))
		}
	}

	// quiz 2
	/*
		cetak char yang hanya huruf vokal
	*/
	fmt.Println("quiz 2")
	for _, c := range title {
		char := string(c)
		// if char == "a" || char == "i" || char == "u" || char == "e" || char == "o" {
		// 	fmt.Println("index-char :", c, "char :", string(c))
		// }
		switch char {
		case "a", "i", "u", "e", "o":
			fmt.Println("char :", string(c))
		}
	}
}
