package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age  int
}

func main() {
	// methods.Intro()
	// methods.EmbeddedInterface()

	// Empty interface -> interface{}
	// often called dynamic typing, dapat menampung segala jenis data

	var secret interface{}

	secret = "ethan hunt"
	fmt.Println(secret)

	secret = []string{"apple", "manggo", "banana"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)

	// data := map[string]interface{}{
	// 	"name":      "ethan hunt",
	// 	"grade":     2,
	// 	"breakfast": []string{"apple", "manggo", "banana"},
	// }

	// pada go v1.18, interface{} dapat diganti dengan syntax any

	// interface{} dan any sama.
	// data := map[interface{}]interface{}{
	data := map[any]any{
		"name":      "ethan hunt",
		true:        2,
		"breakfast": []string{"apple", "manggo", "banana"},
	}

	fmt.Println(data)

	// untuk melakukan operasi yang membutuhkan nilai asli pada variabel yang bertipe interface{}
	var secret2 any
	secret2 = 2
	var number = secret2.(int) * 10
	fmt.Println(secret2, "multiplied by 10 is :", number)

	secret2 = []string{"apple", "manggo", "banana"}
	var gruits = strings.Join(secret2.([]string), ", ")
	fmt.Println(gruits, "is my favorite fruits")

	// teknik casting pada interface disebut dengan type assertions

	// casting any to object pointer
	var secret3 interface{} = &person{name: "wick", age: 27}
	var name = secret3.(*person).name
	fmt.Println(name)

	// kombinasi slice, map, dan interface{}
	var person = []map[string]interface{}{
		{"name": "Wick", "age": 23},
		{"name": "Ethan", "age": 23},
		{"name": "Bourne", "age": 22},
	}

	for _, each := range person {
		fmt.Println(each["name"], "age is", each["age"])
	}

	var fruits = []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 10},
		[]string{"manggo", "pineapple", "papaya"},
		"orange",
	}

	for _, each := range fruits {
		fmt.Println(each)
	}
}
