package main

import (
	"testing"
	"time"
)

func Test_Dine(t *testing.T) {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		completeOrder = []string{}
		dine()
		if len(completeOrder) != 5 {
			t.Errorf("incorrect length of slice; expected 5 but got %d", len(completeOrder))
		}
	}
}
