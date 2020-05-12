package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 3)

	c <- "hallo1"
	c <- "hallo2"
	c <- "hallo3"

	// go func(input chan string) {
	// 	fmt.Println("sending 1 to channel")
	// 	input <- "Hello1"

	// 	fmt.Println("sending 2 to channel")
	// 	input <- "Hello2"

	// 	fmt.Println("sending 3 to channel")
	// 	input <- "Hello3"

	// 	fmt.Println("sending 4 to channel")
	// 	input <- "Hello4"

	// 	fmt.Println("sending 5 to channel")
	// 	input <- "Hello5"
	// }(c)

	// fmt.Println("Received from channel")
	// for greeting := range c {
	// 	fmt.Println("Greeting Received")
	// 	fmt.Println(greeting)
	// }

}

func HelloWorld() {
	fmt.Println("Hello World!")
}
