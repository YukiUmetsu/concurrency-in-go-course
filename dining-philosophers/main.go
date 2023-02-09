package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	rightFork int // number of folk
	leftFork  int // number of folk
}

var hunger = 3 // how many times does a person eat?
var eatTime = 1 * time.Second
var thinkTime = 1 * time.Second
var sleepTime = 1 * time.Second

var philosophers = []Philosopher{
	{name: "a", leftFork: 4, rightFork: 0},
	{name: "b", leftFork: 0, rightFork: 1},
	{name: "c", leftFork: 1, rightFork: 2},
	{name: "d", leftFork: 2, rightFork: 3},
	{name: "e", leftFork: 3, rightFork: 4},
}

// make sure to lock the completion order array to avoid changing the array at the same time
var completeOrderMutex sync.Mutex
var completeOrder []string

func main() {
	fmt.Println("Dining Philosophers Problem")
	fmt.Println("----------------------------")

	dine()
	fmt.Println("Done")
}

func (p Philosopher) eat(wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	// seat the philosopher at the table
	fmt.Printf("Philosopher %s is seated at the table\n", p.name)
	seated.Done()

	// wait until all philosophers seated
	seated.Wait()

	for i := hunger; i > 0; i-- {
		// lock both forks

		// this is to avoid everyone taking left fork and freeze, no one can eat
		// make sure to take the lower number fork first
		if p.leftFork > p.rightFork {
			// only philosopher a
			forks[p.rightFork].Lock()
			fmt.Printf("\tPhilosopher %s takes the right fork\n", p.name)
			forks[p.leftFork].Lock()
			fmt.Printf("\tPhilosopher %s takes the left fork\n", p.name)
		} else {
			forks[p.leftFork].Lock()
			fmt.Printf("\tPhilosopher %s takes the left fork\n", p.name)
			forks[p.rightFork].Lock()
			fmt.Printf("\tPhilosopher %s takes the right fork\n", p.name)
		}

		fmt.Printf("\tPhilosopher %s has both forks and is eating\n", p.name)
		time.Sleep(eatTime)

		fmt.Printf("\tPhilosopher %s is thinking\n", p.name)
		time.Sleep(thinkTime)

		forks[p.leftFork].Unlock()
		forks[p.rightFork].Unlock()

		fmt.Printf("\tPhilosopher %s put down the forks\n", p.name)
	}
	fmt.Printf("Philosopher %s is satisfied and left the table \n", p.name)

	// add finished philosopher to the complete order list
	completeOrderMutex.Lock()
	completeOrder = append(completeOrder, p.name)
	completeOrderMutex.Unlock()
}

func dine() {
	wg := &sync.WaitGroup{}
	// wait until everyone is done eating
	wg.Add(len(philosophers))

	// wait until everyone is seated
	seated := &sync.WaitGroup{}
	seated.Add((len(philosophers)))

	// forks is a map of all 5 folks.
	var forks = prepareForks()

	// start the meal
	for _, philosopher := range philosophers {
		go philosopher.eat(wg, forks, seated)
	}

	wg.Wait()

	fmt.Println("finished eating order: ", strings.Join(completeOrder, ", "))
}

// forks are like locks that stops people eating
func prepareForks() map[int]*sync.Mutex {
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}
	return forks
}
