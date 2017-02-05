package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
	"time"
)

const MAX = 100
const GR = 100000

func sum(values []int64) int64 {
	var sum int64
	for _, v := range values {
		sum += v
	}
	return sum
}

func mean(values []int64) float64 {
	if 0 == len(values) {
		return 0.0
	}
	return float64(sum(values)) / float64(len(values))
}

// SampleMin returns the minimum value of the slice of int64.
func min(values []int64) int64 {
	if 0 == len(values) {
		return 0
	}
	var min int64 = math.MaxInt64
	for _, v := range values {
		if min > v {
			min = v
		}
	}
	return min
}

// SampleMin returns the minimum value of the slice of int64.
func max(values []int64) int64 {
	if 0 == len(values) {
		return 0
	}
	var max int64 = 0
	for _, v := range values {
		if max < v {
			max = v
		}
	}
	return max
}

// int64Slice attaches the methods of Interface to []int, sorting in increasing order.
type int64Slice []int64

func (p int64Slice) Len() int           { return len(p) }
func (p int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func percentile(values int64Slice, p float64) float64 {

	ps := []float64{p}
	scores := make([]float64, 1)
	size := len(values)
	if size > 0 {
		sort.Sort(values)
		for i, p := range ps {
			pos := p * float64(size+1)
			if pos < 1.0 {
				scores[i] = float64(values[0])
			} else if pos >= float64(size) {
				scores[i] = float64(values[size-1])
			} else {
				lower := float64(values[int(pos)-1])
				upper := float64(values[int(pos)])
				scores[i] = lower + (pos-math.Floor(pos))*(upper-lower)
			}
		}
	}
	return scores[0]
}

func consume(c chan int64, w *sync.WaitGroup, r int) {

	defer w.Done()

	a := make(int64Slice, MAX*r)

	ctr := 0

	for {
		t1 := <-c
		if t1 == 0 {
			ctr++
			if ctr == r {
				break
			}
		}

		t2 := time.Now().UnixNano()
		if t2-t1 < 0 {
			fmt.Println(t2, t1)
		}
		a = append(a, t2-t1)
	}

	// fmt.Println(a)
	fmt.Println("95th Percentile is ", percentile(a, 95.0))
	fmt.Println(max(a))
	fmt.Println("Done")
}

func produce(c chan int64, w *sync.WaitGroup) {

	defer w.Done()
	for i := 0; i != MAX; i++ {

		c <- time.Now().UnixNano()
	}

	c <- 0
	// fmt.Println("Am I done")
}

var cw sync.WaitGroup

func work() {
	c := make(chan int64)

	cw.Add(1)
	go consume(c, &cw, GR)

	for i := 0; i != GR; i++ {
		cw.Add(1)
		go produce(c, &cw)
	}

	cw.Wait()
}

func main() {
	work()
}
