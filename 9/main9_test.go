package main

import (
	"testing"
	"time"
)

func TestGetUint8ValuesChannel(t *testing.T) {
	ch := getUint8ValuesChannel()

	if ch == nil {
		t.Fatal("The channel must not be nil")
	}

	select {
	case val := <-ch:
		if val < 0 || val > 255 {
			t.Errorf("The value %d is out of range of uint8", val)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Timeout while reading from channel")
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		name     string
		input    uint8
		expected float64
	}{
		{
			name:     "zero",
			input:    0,
			expected: 0,
		},
		{
			name:     "one",
			input:    1,
			expected: 1,
		},
		{
			name:     "two",
			input:    2,
			expected: 8,
		},
		{
			name:     "max value",
			input:    255,
			expected: 255 * 255 * 255, // 16 581 375
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := make(chan uint8)
			out := make(chan float64)

			process(in, out)
			in <- tt.input
			select {
			case got := <-out:
				if got != tt.expected {
					t.Errorf("Mismatch values expected %v, got %v", tt.expected, got)
				}
			case <-time.After(100 * time.Millisecond):
				t.Fatal("Timeout while reading from channel")
			}
		})
	}
}

func TestProcessMultipleValues(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	process(in, out)

	inputs := []uint8{1, 2, 3, 4, 5, 255}
	expected := []float64{1, 8, 27, 64, 125, 16581375}

	go func() {
		for _, v := range inputs {
			in <- v
		}
	}()

	for i, exp := range expected {
		select {
		case got := <-out:
			if got != exp {
				t.Errorf("Mismatch value - for index %d expected %v, got %v", i, exp, got)
			}
		case <-time.After(100 * time.Millisecond):
			t.Fatalf("Timeout while reading from channel of index %d", i)
		}
	}
}
