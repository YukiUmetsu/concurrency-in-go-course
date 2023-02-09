package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber > NumberOfPizzas {
		return &PizzaOrder{pizzaNumber: pizzaNumber}
	}

	delay := rand.Intn(5) + 1 // 1-5
	fmt.Printf("received order #%d!\n", pizzaNumber)

	rnd := rand.Intn(12) + 1 // 1-12
	msg := ""
	success := false

	if rnd < 5 {
		pizzasFailed++
		msg = fmt.Sprintf("pizza failed for pizza #%d!\n", pizzaNumber)
	} else {
		pizzasMade++
		success = true
		msg = fmt.Sprintf("pizza order #%d is ready!\n", pizzaNumber)
	}
	total++

	fmt.Printf("making pizza #%d. it will take #%d seconds....\n", pizzaNumber, delay)
	// delay for a bit
	time.Sleep(time.Duration(delay) * time.Second)

	p := PizzaOrder{
		pizzaNumber: pizzaNumber,
		message:     msg,
		success:     success,
	}
	return &p

}

func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	var i = 0
	// run forever or until we receive a quit notification
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			case pizzaMaker.data <- *currentPizza:
			case quitChan := <-pizzaMaker.quit:
				// close channels
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// print out messages to
	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("----------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)

	// create and run consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber > NumberOfPizzas {
			color.Cyan("Done making pizzas...")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** Error closing channel!", err)
			}
			break
		}

		if i.success {
			color.Green(i.message)
			color.Green("Order #%d is out for delivery!", i.pizzaNumber)
		} else {
			color.Red(i.message)
			color.Red("The customer is really mad!")
		}
	}

	// print out the ending message
	// print out the ending message
	color.Cyan("-----------------")
	color.Cyan("Done for the day.")

	color.Cyan("We made %d pizzas, but failed to make %d, with %d attempts in total.", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 9:
		color.Red("It was an awful day...")
	case pizzasFailed >= 6:
		color.Red("It was not a very good day...")
	case pizzasFailed >= 4:
		color.Yellow("It was an okay day....")
	case pizzasFailed >= 2:
		color.Yellow("It was a pretty good day!")
	default:
		color.Green("It was a great day!")
	}
}
