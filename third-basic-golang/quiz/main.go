package main

import "fmt"

func main() {

	// hitung rata-rata (jika nilai rata-rata dalam bentuk pecahan, maka harus pecahan)
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	total := 0

	for _, score := range scores {
		total += score
	}

	fmt.Println(float32(total)/float32(len(scores)))


	fmt.Println("====")

	// pilih yang nilainya >= 90 untuk dimasukkan ke dalam slice goodScores
	// scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	var goodScores []int

	for _, score := range scores {
        if score >= 90 {
			goodScores = append(goodScores, score)
		}
	}
	
	fmt.Println(goodScores)
}