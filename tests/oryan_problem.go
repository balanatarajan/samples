package main

import "fmt"
import "os"
import "bufio"

func buildArray(str []string) []int {

}

func main() {

	//Enter your code here. Read input from STDIN. Print output to STDOUT
	scanner := bufio.NewReader(os.Stdin)
	ls, err := scanner.ReadString('\n')

	if err != nil {
		panic("Some Error while reading the string")
	}

	sts := strings.SplitN(ls, " ", -1)

}
