package main

import "fmt"

func main() {
	slc := []int{29, 10, 14, 37, 14, 43, 6, 15, 52, 53, 23, 21, 66, 32, 11, 52, 87, 93, 4, 225, 212, 5, 44, 3}
	mergeSort(slc)

	fmt.Println(slc)
}

func merge(slc []int, mid int) {
	left := make([]int, 0, len(slc))
	right := make([]int, 0, len(slc))

	left = append(left, slc[:mid]...)
	right = append(right, slc[mid:]...)

	i, j, k := 0, 0, 0

	for i < len(left) || j < len(right) {
		if i == len(left) {
			slc[k] = right[j]
			j++
			k++
			continue
		}

		if j == len(right) {
			slc[k] = left[i]
			i++
			k++
			continue
		}

		if left[i] < right[j] {
			slc[k] = left[i]
			i++
		} else {
			slc[k] = right[j]
			j++
		}
		k++
	}
}

func mergeSort(slc []int) {
	if len(slc) < 2 {
		return
	}

	mid := len(slc) / 2

	mergeSort(slc[:mid])
	mergeSort(slc[mid:])
	merge(slc, mid)
}
