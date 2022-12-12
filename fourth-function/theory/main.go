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

	// harus ditangkap semua, jika tidak butuh bisa gunakan syntax ini
	// surface, _ := calculate(5,5)
	surface, circumference := calculate(5,5)
	fmt.Println(surface, circumference)

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

// mulitple returns on a function
func calculate(h, w int) (int, int) {
	surface := h * w
	circumference := 2 * (h + w)
	return surface, circumference
}