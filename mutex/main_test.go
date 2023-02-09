package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_Main(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w
	main()
	_ = w.Close()
	result, _ := io.ReadAll(r)
	resStr := string(result)

	os.Stdout = stdOut
	if !strings.Contains(resStr, "$34320.00") {
		t.Errorf("expected $ is not found")
	}
}
