package main

import "fmt"

func main() {
	slc := []int{29, 10, 14, 37, 14, 43, 6, 15, 52, 21, 66}

	bubbleSort(slc)
	fmt.Println(slc)
}

func bubbleSort(slc []int) {
	for i := len(slc) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if slc[j] > slc[j+1] {
				slc[j], slc[j+1] = slc[j+1], slc[j]
			}
		}
	}
}
