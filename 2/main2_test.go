package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestGetSlice(t *testing.T) {
	fmt.Println("aaa")
	slice := getSlice()
	if len(slice) != 10 {
		t.Errorf("Expected len %d, got %d", 10, len(slice))
	}
}

func TestSliceExample(t *testing.T) {
	inSlice := []int{0, 1, 2, 3, 4, 5}
	got := sliceExample(inSlice)

	expected := []int{0, 2, 4}
	if !slices.Equal(got, expected) {
		t.Errorf("Expected slice %d, got %d", expected, got)
	}
}

func TestAddElements(t *testing.T) {
	inSlice := []int{0, 1}
	got := addElements(inSlice, 100)

	expected := []int{0, 1, 100}
	if !slices.Equal(got, expected) {
		t.Errorf("Expected slice %d, got %d", expected, got)
	}
}

func TestCopySlice(t *testing.T) {
	expected := []int{0, 1}
	got := copySlice(expected)

	gotAddress := fmt.Sprintf("%p", got)
	expectedAddress := fmt.Sprintf("%p", expected)

	if expectedAddress == gotAddress {
		t.Errorf("Expected different address %q, got %q", expectedAddress, gotAddress)
	}

	expected[0] = 100
	if slices.Equal(expected, got) {
		t.Errorf("Expected different slices %d, got %d", expected, got)
	}
}

func TestRemoveElement(t *testing.T) {
	inSlice := []int{0, 1, 2, 3}
	got := removeElement(inSlice, 1)

	expected := []int{0, 2, 3}
	if !slices.Equal(got, expected) {
		t.Errorf("Expected slice %d, got %d", expected, got)
	}
	if slices.Equal(inSlice, got) {
		t.Errorf("Expected different slices %d, got %d", inSlice, got)
	}
}
