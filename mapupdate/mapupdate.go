package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var kv = make(map[string]string)

var w sync.WaitGroup

type command int

const (
	update command = iota
	insert
	read
	exit
)

type data struct {
	index string
	data  string
	cm    command
}

func mapOperations(req, res chan data) {

	// time.Sleep(100 * time.Millisecond)
	// fmt.Println("Here I am")

	for d := range req {
		switch d.cm {
		case update:
		case insert:
			kv[d.index] = d.data
			break
		case read:
			r := data{index: d.index, data: kv[d.index]}
			res <- r
			break
		case exit:
			fmt.Println("Got a message exit")
			close(req)
			break
		}
	}

	w.Done()
}

func extract(req, res chan data) {

	var ctr int
	t1 := time.Now()
	for j := 0; j != 1000; j++ {
		for i := 100; i != 170; i++ {

			ctr++

			t := strconv.Itoa(i)

			d := data{index: t, data: t, cm: read}

			req <- d
			d = <-res

			if i%5 == 0 {
				fmt.Println("Value read is ", d.data)
			}
		}
	}

	t2 := time.Now()

	fmt.Println("Total Time is ", t2.Sub(t1), ctr)

	w.Done()

}
func add(c chan data) {

	for i := 200; i != 276; i++ {

		t := strconv.Itoa(i)

		d := data{index: t, data: t, cm: insert}

		c <- d
	}

	// fmt.Println("Going to call done")
	w.Done()

}
func main() {

	req := make(chan data)
	res := make(chan data)

	for i := 100; i != 176; i++ {
		t := strconv.Itoa(i)

		kv[t] = t
	}

	w.Add(1)

	// Go routine for map operations
	go mapOperations(req, res)

	w.Add(1)
	go add(req)

	w.Add(1)
	go extract(req, res)

	time.Sleep(30 * time.Second)

	d := data{index: "", data: "", cm: exit}
	req <- d
	w.Wait()
	fmt.Println("Total length is ", len(kv))
}
