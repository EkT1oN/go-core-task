package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getUint8ValuesChannel() chan uint8 {
	ch := make(chan uint8)
	go func() {
		for {
			value := uint8(rand.Intn(256))
			fmt.Println("Send:", value)
			ch <- value
		}
	}()
	return ch
}

func process(in chan uint8, out chan float64) {
	go func() {
		for {
			value := <-in
			convertValue := float64(value)
			cubeValue := convertValue * convertValue * convertValue
			out <- cubeValue
		}
	}()
}

func main() {
	in := getUint8ValuesChannel()
	out := make(chan float64)
	process(in, out)
	for range 10 {
		time.Sleep(500 * time.Millisecond)
		got := <-out
		fmt.Println("Get:", got)
	}
}
