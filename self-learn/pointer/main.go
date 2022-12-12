package main

import "fmt"

// pointer -> reference/alamat memori.

// referencing -> mengambil nilai pointer dari suatu variabel biasa,
// caranya adalah menggunakan tanda & sebelum nama variabel (&variable)

// deferencing -> mengambil nilai asli dari variabel pointer,
// caranya adalah menggunakan * sebelum nama variabel (*variable)

func main() {

	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)   :", numberA)  // 4
	fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

	fmt.Println("numberB (value)   :", *numberB) // 4
	fmt.Println("numberB (address) :", numberB)  // 0xc20800a220

	var numberC int = *numberB
	fmt.Println(numberC)

}
