package main

import (
	"fmt"
	"os"
	"strconv"

	"./rb"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Fprintf(os.Stderr, "error : args to short\n")
		return
	}
	role := args[1]

	switch role {
	case "server":
		portStr := args[2]
		b, err := strconv.Atoi(portStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error port :%s\n", portStr)
			return
		}
		rb.StartServer(b)
	case "client":
		if len(args) < 4 {
			fmt.Fprintf(os.Stderr, "error : args too short\n")
			return
		}
		b, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error port :%s\n", args[3])
		}
		rb.StartClient(args[2], b)
	default:
		fmt.Fprintf(os.Stderr, "error cmd %s\n", args[1])
	}

}
