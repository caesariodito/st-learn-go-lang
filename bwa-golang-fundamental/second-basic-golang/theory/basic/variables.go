package basic

import "fmt"

func Variables() {
	var name string = "Golang"
	// var name string
	// name = "Golang2"
	var yakin bool = true

	fmt.Println(name)
	fmt.Println(yakin)

	// jika ingin membuat variabel dan langsung diisi, gunakan line yang dibawah ini saja
	jago := false
	umur := 20

	fmt.Println(jago)
	fmt.Println(umur)

	umur = 22

	fmt.Println(umur)
}