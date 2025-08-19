package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	// Capture stdout
	var buf bytes.Buffer
	log.SetOutput(&buf)
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout

	var output strings.Builder
	_, err := output.ReadFrom(r)
	if err != nil {
		t.Fatalf("Failed to read output: %v", err)
	}

	expected := "Hello, Go!\n"
	if output.String() != expected {
		t.Errorf("Expected output %q, got %q", expected, output.String())
	}
}
