package main

import (
	"fmt"
	"os"
	"strconv"

	"../../rb"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Fprintf(os.Stderr, "error : args too short\n")
		return
	}
	b, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error port: %s\n", args[2])
	}
	rb.StartClient(args[1], b)
}
