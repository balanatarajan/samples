package main

import (
	"fmt"
	"math"
)

func bubbleup(arr []int, n int) {

	pofn := int(n / 2)

	if arr[n] < arr[pofn] {
		arr[n], arr[pofn] = arr[pofn], arr[n]
		bubbleup(arr, pofn)
	}

}

func bubbledown(arr []int, start int) {
	ch1, ch2 := int(math.Pow(2, float64(start))), int(math.Pow(2, float64(start))+1)

	if ch2 > len(arr) {
		return
	}

	diff1, diff2 := arr[start]-arr[ch1], arr[start]-arr[ch2]

	if diff1 > diff2 {
		arr[start], arr[ch1] = arr[ch1], arr[start]
		bubbledown(arr, ch1)
	} else if diff2 > diff1 {
		arr[start], arr[ch2] = arr[ch2], arr[start]
		bubbledown(arr, ch2)
	}
}

func main() {
	x := []int{0, 18, 20, 0, 34, 56, 78}

	x[3] = 6

	fmt.Println(x)
	bubbleup(x, 3)
	fmt.Println(x)

	x[1] = 99
	fmt.Println(x)
	bubbledown(x, 1)
	fmt.Println(x)
}
