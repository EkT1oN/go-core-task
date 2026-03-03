package main_7

import (
	"fmt"
	"sync"
)

func mergeChannels(in ...chan int) chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	for _, ch := range in {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for value := range ch {
				fmt.Println("Send:", value)
				out <- value
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
}
