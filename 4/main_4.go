package main

import (
	"slices"
)

func getSubSlice(src []string, target []string) []string {
	res := []string{}
	for _, value := range src {
		if !slices.Contains(target, value) {
			res = append(res, value)
		}
	}
	return res
}

func main() {
}
