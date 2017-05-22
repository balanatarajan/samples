package main

import "fmt"

func main() {

	mp := map[int]int{4: 1, 5: 4}

	fmt.Println(mp)

	mp[4]++
	mp[1]++

	fmt.Println(mp)
}
