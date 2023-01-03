package main

import (
	"fmt"
	"runtime"
)

func printMessage(what chan string) {
	fmt.Println(<-what)
}

func main() {
	// channel

	// part 1
	// runtime.GOMAXPROCS(2)

	// var messages = make(chan string)

	// var sayHelloTo = func(who string) {
	// 	var data = fmt.Sprintf("hello %s", who)
	// 	messages <- data
	// }

	// go sayHelloTo("john wick")
	// go sayHelloTo("ethan hunt")
	// go sayHelloTo("jason bourne")

	// var message1 = <-messages
	// fmt.Println(message1)

	// var message2 = <-messages
	// fmt.Println(message2)

	// var message3 = <-messages
	// fmt.Println(message3)

	// part 2

	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	for _, each := range []string{"wick", "hunt", "bourne"} {
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			messages <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages)
	}
}
