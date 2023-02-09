package main

import (
	"fmt"
	"sync"
)

type Income struct {
	Source string
	Amount int
}

var wg sync.WaitGroup

func main() {
	var bankBalance int
	var mux sync.Mutex

	incomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}

	wg.Add(len(incomes))
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				mux.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				mux.Unlock()
			}
		}(i, income)
	}

	wg.Wait()
	fmt.Printf("final balance: $%d.00\n", bankBalance)
}
