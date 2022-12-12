package main

import "fmt"

func main() {

	/*
		parameter/input
		proses
		output
	*/

	helloWorld()
	greet("Agus")
	result := add(1,2)

	fmt.Println(result)
}

func helloWorld() {
	fmt.Println("Hello, world!")
}

func greet(name string) {
	fmt.Println("Hello,", name)
}

// bisa mempersingkat parameter dengan tipe data variabel yang sama
// func add(a int, b int) int {
func add(a, b int) int {
	return a+b
}