package main

import (
	"math/rand"
)

func generator(seed rand.Source) chan int {
	r := rand.New(seed)
	ch := make(chan int)
	go func() {
		for {
			ch <- r.Intn(100)
		}
	}()
	return ch
}

func main() {
}
