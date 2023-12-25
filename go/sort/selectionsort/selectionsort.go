package main

import "fmt"

func main() {
	slc := []int{23,12,53,6,4,2,16,84,25,35,33,48,82,67}
	selectionSort(slc)

	fmt.Println(slc)
}

func selectionSort(slc []int) {
	for idx := 0; idx < len(slc);idx++ {
		min := slc[idx]
		minIdx := idx

		for i := idx + 1; i < len(slc); i++ {
			if slc[i] < min {
				min = slc[i]
				minIdx = i
			}
		}

		slc[idx], slc[minIdx] = slc[minIdx], slc[idx]
	}
}
