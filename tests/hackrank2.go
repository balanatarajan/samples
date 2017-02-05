package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var i uint32 = 4
	var d float32 = 4.0
	var s string = "HackerRank"

	scanner := bufio.NewReader(os.Stdin)

	// Declare second integer, double, and String variables.
	var si uint32
	var sd float32
	var st string

	// Read and save an integer, double, and String to your variables.
	str, err := scanner.ReadString('\n')

	if err != nil {
		panic("Error scanning")
	}

	fmt.Sscan(str, &si)

	str, err = scanner.ReadString('\n')
	if err != nil {
		panic("Error scanning")
	}

	fmt.Sscan(str, &sd)

	st, err = scanner.ReadString('\n')
	if err != nil {
		panic("Error scanning")
	}

	// Print the sum of both integer variables on a new line.
	fmt.Println(i + si)

	// Print the sum of the double variables on a new line.
	fmt.Println(d + sd)

	// Concatenate and print the String variables on a new line
	fmt.Println(s + st)
}
