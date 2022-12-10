package basic

import "fmt"

func ConditionSwitch() {
	number := 3

	switch number {
	case 1:
		fmt.Println("satu")
	case 2:
		fmt.Println("dua")
	case 3:
		fmt.Println("tiga")
	case 4:
		fmt.Println("empat")
	default:
		fmt.Println("gatau")
	}
}
