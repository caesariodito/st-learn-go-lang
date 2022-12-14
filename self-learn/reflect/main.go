package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) SetName(name, greet string) {
	s.Name = name
	fmt.Println(greet)
}

// func (s *student) getPropertyInfo() {
// 	fmt.Println()
// 	var reflectValue = reflect.ValueOf(s) // mengambil nilai dari struct

// 	fmt.Println(reflectValue) // pointer

// 	var reflectTypeTest = reflectValue.Type()
// 	fmt.Println(reflectTypeTest)

// 	if reflectValue.Kind() == reflect.Ptr { // mengecek apakah pointer atau tidak
// 		reflectValue = reflectValue.Elem() // mengambil objek reflect aslinya
// 	}

// 	fmt.Println(reflectValue) // value of the address in pointer

// 	var reflectType = reflectValue.Type()
// 	fmt.Println(reflectType)

// 	for i := 0; i < reflectType.NumField(); i++ {
// 		fmt.Println("nama      :", reflectType.Field(i).Name)
// 		fmt.Println("tipe data :", reflectType.Field(i).Type)
// 		fmt.Println("nilai     :", reflectValue.Field(i).Interface())
// 		fmt.Println("")
// 	}

// 	// Pengambilan informasi property,
// 	// selain menggunakan indeks,
// 	// bisa diambil berdasarkan nama field dengan menggunakan method FieldByName()
// }

func main() {

	// reflect, simpelnya adalah "type" di python. Namun bisa lebih? i think.
	// terdapat 2 fungsi yang paling penting yaitu, reflect.ValueOf() dan reflect.TypeOf()

	// 1. ValueOf -> nilai
	// 2. TypeOf -> tipe

	// number := 23
	// reflectValue := reflect.ValueOf(number)

	// fmt.Println("reflectValue:", reflectValue)
	// fmt.Println("reflectType:", reflectValue.Type())
	// fmt.Println("reflectTypeOf:", reflect.TypeOf(number))

	// if reflectValue.Kind() == reflect.Int {
	// 	fmt.Println("nilai variabel :", reflectValue.Int())
	// }

	// fmt.Println(reflectValue.Type())

	// var number2 = 23
	// var reflectValue2 = reflect.ValueOf(number2)

	// fmt.Println("tipe  variabel :", reflectValue2.Type())
	// fmt.Println("nilai variabel :", reflectValue2.Interface())

	// var nilai = reflectValue2.Interface()
	// fmt.Println(reflect.TypeOf(nilai))

	// var s1 = &student{Name: "wick", Grade: 2}
	// s1.getPropertyInfo()

	var s1 = &student{Name: "john wick", Grade: 2}
	fmt.Println("nama :", s1.Name)

	var reflectValue = reflect.ValueOf(s1)
	var method = reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{ // mengisi parameter
		reflect.ValueOf("wick"),
		reflect.ValueOf("Testing"),
	})

	fmt.Println("nama :", s1.Name)

}
