package main

import "net"
import "fmt"
import "time"

func main() {

	conn, err := net.Dial("tcp", ":10007")
	if err != nil {
		fmt.Println("Exiting")
		return
	}

	for {
		time.Sleep(100 * time.Millisecond)
		fmt.Fprintf(conn, "1")
	}

}
