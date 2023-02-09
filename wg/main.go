package main

import (
	"fmt"
	"sync"
)

func customPrint(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup
	words := []string{"alpha", "beta", "delta", "gamma", "pi", "zeta", "eta", "theta", "epsilon"}
	wg.Add(len(words))
	for _, word := range words {
		customPrint(word, &wg)
	}

	wg.Wait()

	wg.Add(1)
	customPrint("this is the last one", &wg)
}
