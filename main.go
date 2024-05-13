package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("Usage: mygocli <operation> <num1> <num2>")
		os.Exit(1)
	}

	op := args[0]
	num1, _ := strconv.Atoi(args[1])
	num2, _ := strconv.Atoi(args[2])

	switch op {
	case "add":
		fmt.Println("Result:", num1+num2)
	case "sub":
		fmt.Println("Result:", num1-num2)
	default:
		fmt.Println("Unsupported operation")
		os.Exit(1)
	}
}
