package main

import "fmt"
import "bufio"
import "os"
import "math"
import "io"

func main() {

	scanner := bufio.NewReader(os.Stdin)

	str, err := scanner.ReadString('\n')

	if err != nil {
		panic("Error scanning")
	}

	var mealCost float64
	fmt.Sscan(str, &mealCost)

	str, err = scanner.ReadString('\n')
	if err != nil {
		panic("Error scanning")
	}

	var tip int64
	fmt.Sscan(str, &tip)

	str, err = scanner.ReadString('\n')

	if err != nil && err != io.EOF {
		panic("Error scanning")
	}

	var tax int64
	fmt.Sscan(str, &tax)

	cost := mealCost + (mealCost*float64(tip))/100.0 + (mealCost*float64(tax))/100

	fmt.Println(cost)

	fmt.Printf("The total meal cost is %.0f ", math.Floor(cost+0.5))
	fmt.Printf(" dollars \n")
}
