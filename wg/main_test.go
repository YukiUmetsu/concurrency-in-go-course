package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func test_CustomPrint(t *testing.T) {
	// store the original stdout
	stdOut := os.Stdout

	// pipe connects read & write process
	r, w, _ := os.Pipe()
	os.Stdout = w

	// create wait group to wait for the customPrint
	var wg sync.WaitGroup
	wg.Add(1)
	go customPrint("epsilon", &wg)

	// wait for customPrint to end
	wg.Wait()

	_ = w.Close()

	// read from the reader
	result, _ := io.ReadAll(r)
	output := string(result)

	// bring the os.Stdout back to the original
	os.Stdout = stdOut

	if !strings.Contains(output, "epsilon") {
		t.Errorf("Espected to find epsilon but it is not there.")
	}
}
