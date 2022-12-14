package methods

import (
	"fmt"
	"math"
)

// seperti abstract class di java
// kedua method di dalam interface harus diimplementasikan semua ke dalam struct,
// jika tidak, akan error
type hitungLama interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return 4 * p.sisi
}

type segitiga struct {
	alas   float64
	tinggi float64
}

func (s segitiga) luas() float64 {
	return s.alas * s.tinggi / 2
}

func (s segitiga) keliling() float64 {
	return 0
}

func Intro() {
	var bangunDatar hitungLama

	bangunDatar = persegi{10.0}
	fmt.Println("===== persegi")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())

	bangunDatar = lingkaran{5.0}
	fmt.Println("===== lingkaran")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())
	fmt.Println("jari-jari :", bangunDatar.(lingkaran).jariJari())

	bangunDatar = segitiga{10.0, 4.0}
	fmt.Println("===== segitiga")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())
}
