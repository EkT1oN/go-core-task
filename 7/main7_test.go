package main_7

import (
	"slices"
	"sort"
	"sync"
	"testing"
)

func TestMergeChannels(t *testing.T) {
	const N = 3
	channels := []chan int{}
	for range N {
		in := make(chan int)
		channels = append(channels, in)
	}

	out := mergeChannels(channels...)
	done := make(chan struct{})
	got := []int{}
	var mu sync.Mutex
	go func() {
		for value := range out {
			mu.Lock()
			got = append(got, value)
			mu.Unlock()
		}
		done <- struct{}{}
	}()

	channels[0] <- 0
	channels[0] <- 10
	channels[0] <- 100
	channels[1] <- 1
	channels[1] <- 11
	channels[1] <- 111
	channels[2] <- 2
	channels[2] <- 22
	channels[2] <- 222
	for _, ch := range channels {
		close(ch)
	}

	<-done
	expected := []int{0, 10, 100, 1, 11, 111, 2, 22, 222}

	sort.Ints(got)
	sort.Ints(expected)
	if !slices.Equal(got, expected) {
		t.Errorf("Mismatch channels send/reveice values: expected %v, got %v", expected, got)
	}
	if len(expected) != len(got) {
		t.Errorf("Mismatch channels send/receive values count - expected %v, got %v", len(expected), len(got))
	}
}
