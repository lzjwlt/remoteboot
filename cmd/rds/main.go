package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"../../rb"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Println("Error: args 'port' cant be detected")
		return
	}
	portStr := args[1]
	b, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error port: %s\n", portStr)
		return
	}
	rb.StartServer(b)
}
