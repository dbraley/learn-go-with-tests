package main

import (
	"fmt"
	"io"
	"os"
)

// Countdown counts down from 3
func Countdown(out io.Writer) {
	fmt.Fprint(out, "3")
}

func main() {
	Countdown(os.Stdout)
}
