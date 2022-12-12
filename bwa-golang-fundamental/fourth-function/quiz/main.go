package main

import (
	"errors"
	"fmt"
)

func main() {
	// 1. membuat sebuah function dengan nama 'sum' -> parameter array int
	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)
	fmt.Println(total)

	// 2. membuat sebuah function dengan nama 'calculate', parameter(number, number2, operation) int
	result, err := calculate(10,2,"+")
	fmt.Println(result, err)
	result, err = calculate(10,2,"-")
	fmt.Println(result, err)
	result, err = calculate(10,2,"*")
	fmt.Println(result, err)
	result, err = calculate(10,2,"a")
	fmt.Println(result, err)

	// untuk mengecek error, ada error
	if err != nil {
		fmt.Println("ada error:", err.Error())
	}
	// err -> error
	// simple string-based error
	// err1 := errors.New("math: square root of negative number")



}

func sum(numbers []int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func calculate(a, b int, opr string) (result int, err error) {
	switch opr {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		err = errors.New("invalid operation")
	}
	return
}