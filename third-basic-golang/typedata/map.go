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

func MapPartTwo(){
	// Map part two
	myMap := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	// fmt.Println(myMap)
	
	// iterate
	for key, value := range myMap {
		fmt.Println(key, value)
	}

	fmt.Println("=====")
	// delete
	delete(myMap, "key3")

	// iterate
	for key, value := range myMap {
		fmt.Println(key, value)
	}

	fmt.Println("=====")
	// mengetahui apakah terdapat nilai di dalam map dengan key tertentu
	value := myMap["key2"]
	fmt.Println(value) // value2
	value = myMap["key3"]
	fmt.Println(value) // kosong
	value, isAvailable := myMap["key1"]
	fmt.Println(value, isAvailable) // value1, true
}