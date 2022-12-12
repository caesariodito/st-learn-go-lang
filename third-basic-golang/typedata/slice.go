package typedata

import "fmt"

func Slice() {
	// slice -> dynamic
	var gamingConsoles []string
	// gabisa melakukan ini, harus menggunakan append
	// gamingConsoles[0] = "Playstation4"
	gamingConsoles = append(gamingConsoles, "Playstation 4")
	gamingConsoles = append(gamingConsoles, "Nintendo Switch")
	gamingConsoles = append(gamingConsoles, "XBOX One")
	fmt.Println(gamingConsoles)

	// this is slice
	peripherals := []string{"Mouse", "Keyboard", "Headset"}
	peripherals = append(peripherals, "Mic")
	fmt.Println(peripherals)

	for _, elem := range peripherals {
		fmt.Println(elem)
	}
}

func SliceOfMap(){
	// map dan slice bisa digabung
	// di dalam slice terdapat 3 map
	students := []map[string]string{
		{"name": "Agung", "score": "A"},
		{"name": "Link", "score": "B"},
		{"name": "Mario", "score": "E"},
	}

	for _, student := range students {
		// fmt.Println(student)
		// map[name:Agung score:A]
		// map[name:Link score:B]
		// map[name:Mario score:E]

		// fmt.Println(student["name"])
		// Agung
		// Link
		// Mario

		// fmt.Println(student["score"])
		// A
		// B
		// E

		fmt.Println(student["name"], "-", student["score"])
		// Agung - A
		// Link - B
		// Mario - E
	}
}