package main

func getSubSlice(in1 []int, in2 []int) (bool, []int) {
	set := make(map[int]struct{})
	for _, v := range in2 {
		set[v] = struct{}{}
	}

	res := []int{}
	for _, element := range in1 {
		if _, exists := set[element]; exists {
			res = append(res, element)
		}
	}

	return len(res) > 0, res
}

func main() {
}
