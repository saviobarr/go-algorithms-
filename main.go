package main

func main() {
	bubbleSort()
}
func bubbleSort() {
	var scores = []int{90, 70, 50, 80, 60, 85}
	var length = len(scores)
	for i := 0; i < length-1; i++ {
		var isSwap = false
		for j := 0; j < length-i-1; j++ {
			if scores[j] > scores[j+1] {
				var flag = scores[j]
				scores[j] = scores[j+1]
				scores[j+1] = flag
				isSwap = true
			}
		}
		if isSwap == false {
			break
		}
	}
}
