package main

import (
	_ "bufio"
	"fmt"
	"net"
	_ "os"
	"runtime"
)

func doWork(c net.Conn, ch chan int) {

	fmt.Println("Local Address is ", c.LocalAddr())
	fmt.Println("Doing work")
	// rd := bufio.NewReader(os.Stdin)
	// b := make([]byte, 128)
	for {
		runtime.Gosched()
		// <-ch
		// rd.Read(b)
	}
}

func main() {

	ln, err := net.Listen("tcp", ":10007")

	if err != nil {
		fmt.Println("Exiting")
		return
	}

	c := make(chan int)

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go doWork(conn, c)
	}
}
