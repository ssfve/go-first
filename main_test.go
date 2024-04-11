package main

import (
	_ "bytes"
	_ "io/ioutil"
	"os"
	"testing"
)

func TestPrintHelloWorld(t *testing.T) {
	// Redirect stdout to a buffer
	//var outBuf bytes.Buffer
	tmpfile, err := os.CreateTemp("", "test_stdout_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {

		}
	}(tmpfile.Name())

	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }() // Restore stdout after the test
	os.Stdout = tmpfile

	// Call the function
	main()

	// Close the temporary file to ensure the data is written
	err = tmpfile.Close()
	if err != nil {
		t.Fatalf("Failed to close temporary file: %v", err)
	}

	// Read the contents of the temporary file
	content, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to read temporary file: %v", err)
	}
	// Check the output
	expected := "Hello, World!\n"
	if string(content) != expected {
		t.Errorf("Expected output: %q, but got: %q", expected, string(content))
	}
}
