package main

import (
	"fmt"
)

func cmpPtr(i, j *int) {
	fmt.Println("Second one", i == j)
	fmt.Println(i, j)
}
func main() {

	var j int
	a, b := &j, &j

	fmt.Println("First one", a == b)
	fmt.Println(a, b)
	cmpPtr(a, &j)

}
