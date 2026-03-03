package main

import (
	"slices"
	"testing"
)

func TestGetSubSlice(t *testing.T) {
	tests := []struct {
		name     string
		slice1   []int
		slice2   []int
		expected []int
		ok       bool
	}{
		{
			"1",
			[]int{65, 3, 58, 678, 64},
			[]int{64, 2, 3, 43},
			[]int{3, 64},
			true,
		},
		{
			"2",
			[]int{1, 2, 3, 4, 5},
			[]int{6, 7, 8, 9, 10},
			[]int{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, got := getSubSlice(tt.slice1, tt.slice2)

			if ok != tt.ok {
				t.Errorf("Mismatch exists - expected %v, got %v", tt.ok, ok)
			}
			if !slices.Equal(got, tt.expected) {
				t.Errorf("Mismatch slices - expected %v, got %v", tt.expected, got)
			}
		})
	}
}
