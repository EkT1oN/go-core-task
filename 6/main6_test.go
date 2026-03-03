package main

import (
	"math/rand"
	"slices"
	"testing"
)

func TestGenerator(t *testing.T) {
	ch := generator(rand.NewSource(0))
	length := 10
	slice := []int{}
	for range length {
		v := <-ch
		slice = append(slice, v)
	}

	if len(slice) != length {
		t.Errorf("Mismatch channel read count - expected %v, got %v", length, len(slice))
	}
}

func TestGeneratorWithSeed(t *testing.T) {
	seed := rand.New(rand.NewSource(42))
	ch := generator(seed)
	length := 10
	got := []int{}
	for range length {
		v := <-ch
		got = append(got, v)
	}

	expected := []int{5, 87, 63, 97, 38, 11, 79, 74, 7, 0}

	if slices.Equal(expected, got) {
		t.Errorf("Mismatch values - expected %v, got %v", expected, got)
	}
}
