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