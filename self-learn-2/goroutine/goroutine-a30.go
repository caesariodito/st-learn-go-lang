package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	// menentukan banyak processor yang digunakan untuk melakukan goroutine
	runtime.GOMAXPROCS(2)

	// menggunakan goroutine
	go print(5, "halo")
	print(5, "apa kabar")

	var input string
	//
	fmt.Scanln(&input)

	fmt.Println("Part 2")
	var s1, s2, s3 string
	fmt.Scanln(&s1, &s2, &s3)

	// user inputs: "trafalgar d law"

	fmt.Println(s1) // trafalgar
	fmt.Println(s2) // d
	fmt.Println(s3) // law
}
