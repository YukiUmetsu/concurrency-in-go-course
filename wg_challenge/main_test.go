package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestUpdateMessage(t *testing.T) {
	wg.Add(1)
	go updateMessage("epsilon")
	wg.Wait()
	if msg != "epsilon" {
		t.Errorf("expected to find epsilon, but it is not there")
	}
}

func TestPrintMessage(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	msg = "epsilon"
	printMessage()
	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut
	if !strings.Contains(output, "epsilon") {
		t.Errorf("expected to find epsilon, but it is not there")
	}
}

func TestMain(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()

	msg = "Hello, world!"

	wg.Add(1)
	go updateMessage("Hello, universe!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, cosmos!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, world!")
	wg.Wait()
	printMessage()

	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut

	universeIndex := strings.Index(output, "universe")
	cosmosIndex := strings.Index(output, "cosmos")
	worldIndex := strings.Index(output, "world!")
	if universeIndex > cosmosIndex || worldIndex > cosmosIndex || universeIndex > worldIndex {
		t.Errorf("expected to print universe - cosmos - world in order, but did not get the order")
	}
}
