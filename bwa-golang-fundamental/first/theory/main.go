package main

import (
	"fmt"

	// cara import -> "root-module/folder"
	"first/calculation"
)

/*
when and why should we use import?
-> standard library, ex: "fmt" for printing
-> access code on different package
-> codes on github.com or code yang dibikin orang lain
*/

/*
func main() -> untuk menjalankan code (executable), wajib jika kita ingin menjalankan kode
*/
func main() {
// func unmain() {
	fmt.Println("Hello, world!")

	result := calculation.Add(1,2)

	fmt.Println(result)

	fmt.Println(calculation.Testing())

	// sentence := Test()
	// fmt.Println(sentence)
}

/* Types of package:
1. Executable (main)
2. Library
*/