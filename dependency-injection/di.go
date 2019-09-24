package main

import (
	"fmt"
	"io"
	"os"
)

// Greet says hello to someone on the provided input streams
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s\n", name)
}

func main() {
	Greet(os.Stdout, "Elodie")
}
