package main

import (
	"slices"
	"testing"
)

func TestGetSubSlice(t *testing.T) {
	tests := []struct {
		name     string
		slice1   []string
		slice2   []string
		expected []string
	}{
		{
			"equal",
			[]string{"a", "b", "c", "d", "e", "f"},
			[]string{"a", "b", "c", "d", "e", "f"},
			[]string{},
		},
		{
			"1",
			[]string{"a", "b", "c", "d", "e", "f"},
			[]string{"a", "c", "e"},
			[]string{"b", "d", "f"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getSubSlice(tt.slice1, tt.slice2)

			if !slices.Equal(got, tt.expected) {
				t.Errorf("Mismatch slices - expected %q, got %q", tt.expected, got)
			}
		})
	}

}
