package main

import "fmt"

func main() {
	bubbleSort()
}
func bubbleSort() {
	var scores = []int{90, 70, 50, 80, 60, 85}
	fmt.Println(scores)
	var length = len(scores)
	for i := 0; i < length; i++ {
		sorted := false
		for j := i + 1; j < length; j++ {
			if scores[i] > scores[j] {
				minor := scores[j]
				scores[j] = scores[i]
				scores[i] = minor
				sorted = true
			}
		}
		if !sorted {
			break
		}

	}
	fmt.Println(scores)
}
