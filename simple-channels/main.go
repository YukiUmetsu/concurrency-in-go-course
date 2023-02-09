package main

import (
	"fmt"
	"strings"
)

// you can only receive from ping channel
// you can only send to poing channel
func shout(ping <-chan string, pong chan<- string) {
	for {
		s := <-ping
		pong <- fmt.Sprintf("%s!!!!", strings.ToUpper(s))
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)
	defer close(ping)
	defer close(pong)

	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (type Q to quit")

	for {
		fmt.Print("->")

		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if strings.ToLower(userInput) == "q" {
			break
		}
		ping <- userInput

		res := <-pong

		fmt.Printf("%v\n", res)
	}

	fmt.Println("DONE!")
}
