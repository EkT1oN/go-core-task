package main

import (
	"fmt"
	"math/rand"
)

func getSlice() []int {
	randomSlice := make([]int, 10)
	for i := range randomSlice {
		randomSlice[i] = rand.Int()
	}
	return randomSlice
}

func sliceExample(in []int) []int {
	evenSlice := []int{}
	for i := range in {
		if (in[i] % 2) == 0 {
			evenSlice = append(evenSlice, in[i])
		}
	}
	return evenSlice
}

func addElements(in []int, element int) []int {
	out := make([]int, len(in))
	copy(out, in)
	out = append(out, element)
	return out
}

func copySlice(in []int) []int {
	out := make([]int, len(in))
	copy(out, in)
	return out
}

func removeElement(in []int, index int) []int {
	if index < 0 || index >= len(in) {
		return copySlice(in)
	}

	result := make([]int, 0, len(in)-1)
	result = append(result, in[:index]...)
	result = append(result, in[index+1:]...)
	return result
}

func main() {
	originalSlice := getSlice()
	fmt.Println("originalSlice:", originalSlice)
	evenSlice := sliceExample(originalSlice)
	fmt.Println("evenSlice:", evenSlice)
	newSlice := addElements(evenSlice, rand.Int())
	fmt.Println("newSlice:", newSlice)
	copyNewSlice := copySlice(newSlice)
	fmt.Println("copyNewSlice:", copyNewSlice)

	fmt.Println("adress compare")
	fmt.Printf("newSlice: %p\n", newSlice)
	fmt.Printf("copyNewSlice: %p\n", copyNewSlice)

	removeSlice := removeElement(copyNewSlice, 5)
	fmt.Println("removeSlice:", removeSlice)

}
