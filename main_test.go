package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	// create pipe
	stderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	// wait output
	ch := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		ch <- buf.String()
	}()

	// call main function
	run()

	// reset
	w.Close()
	os.Stderr = stderr
	output := <-ch

	if !strings.Contains(output, "version") {
		t.Fatal("should output 'version'")
	}
	if strings.Contains(output, "hello") {
		t.Fatal("should not output 'hello'")
	}
}
