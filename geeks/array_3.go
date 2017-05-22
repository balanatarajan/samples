package main

// Given an array A[] and a number x, check for pair in A[] with sum as x
import (
	"fmt"
	"sort"
)

func sumFind(arr []int, sum int) (int, int) {

	sort.Ints(arr)

	fmt.Println(arr)
	for i, j := 0, len(arr)-1; i != len(arr)-1 || j != 0; {

		fmt.Println(arr[i], arr[j])
		if arr[i]+arr[j] < sum {
			i++
			continue
		}

		if arr[i]+arr[j] > sum {
			j--
			continue
		}

		if arr[i]+arr[j] == sum {
			return arr[i], arr[j]
		}
	}

	return 0, 0
}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(sumFind(arr, 9))

}
