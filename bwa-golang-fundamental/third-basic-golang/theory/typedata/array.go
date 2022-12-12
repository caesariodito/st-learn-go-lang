package typedata

import "fmt"

func Array() {
	// tipe data di dalam array harus sama

	var languages [5]string
	languages[0] = "Go"
	languages[1] = "Java"
	languages[2] = "Ruby"
	languages[3] = "Javascript"
	languages[4] = "Python"

	fmt.Println(len(languages))
	fmt.Println(languages)

	routines := [3]string{"Wakeup", "Sleep", "Gaming"}

	fmt.Println(len(routines))
	fmt.Println(routines)

	// pengisian secara vertikal membutuhkan koma pada element terakhir
	// tanda ... akan menghitung sendiri berapa element yang terdapat di array
	games := [...]string{
		"Valorant",
		"Dota",
		"PUBG",
		"Overwatch",
	}

	fmt.Println(len(games))
	fmt.Println(games)

	for index, elem := range games {
		fmt.Println(index, elem)
	}
}