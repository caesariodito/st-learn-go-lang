package basic

import "fmt"

func Loop() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("I'm learning Golang :", i)
	// }

	// while with for
	// i := 1
	// for i <= 100 {
	// 	fmt.Println("I'm learning Golang :", i)
	// 	i++
	// }

	title := "Golang is the best language to learn"

	// for index, char := range title {
	// 	fmt.Println("index :", index, " Char :", string(char))
	// }

	for _, char := range title {
		fmt.Println("Char :", string(char))
	}
}
