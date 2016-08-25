package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func output() string {
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
	res := <-ch
	return res
}

func TestMain(t *testing.T) {
	s := output()

	if !strings.Contains(s, "version") {
		t.Fatal("should output 'version'")
	}

	if strings.Contains(s, "hello") {
		t.Fatal("should not output 'hello'")
	}
}
