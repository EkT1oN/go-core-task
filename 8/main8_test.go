package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCustomWaitGroup(t *testing.T) {
	wg := NewCustomWaitGroup()

	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			fmt.Println("goroutine", id, "start")
			time.Sleep(time.Duration(id) * 500 * time.Millisecond)
			fmt.Println("goroutine", id, "finished")
		}(i)
	}

	fmt.Println("wait...")
	wg.Wait()
	fmt.Println("all goroutines finished, program end")
}
