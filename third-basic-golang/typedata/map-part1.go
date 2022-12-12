package typedata

import "fmt"

func MapPart1() {
	// Array -> index/key berupa angka
	// Map -> index/key bebas (seperti dictionary)
	// membuat map dengan nama variabel myMap dengna key bertipe data string
	// dan value bertipe data int
	var myMap map[string]int // deklarasi saja, belum ada nilai default
	myMap = map[string]int{} // mengisi nilai default
	myMap["ruby"] = 9
	myMap["java"] = 12
	myMap["go"] = 8
	// tertimpa
	myMap["go"] = 10

	fmt.Println(myMap["ruby"]) // 8
	fmt.Println(myMap)

	valid := map[string]string{
		"jago": "yes",
		"cupu": "no",
	}
	fmt.Println(valid)
}