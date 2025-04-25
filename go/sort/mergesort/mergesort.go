package main

import "fmt"

func main() {
	slc := []int{29, 10, 14, 37, 14, 43, 6, 15, 52, 53, 23, 21, 66, 32, 11, 52, 87, 93, 4, 225, 212, 5, 44, 3}
	mergeSort(slc)

	fmt.Println(slc)
}

func merge(slc []int, mid int) {
	aux := make([]int, len(slc))
	copy(aux, slc)

	i, j, k := 0, mid, 0

	for i < mid && j < len(slc) {
		if aux[i] <= aux[j] {
			slc[k] = aux[i]
			i++
		} else {
			slc[k] = aux[j]
			j++
		}
		k++
	}

	for i < mid {
		slc[k] = aux[i]
		i++
		k++
	}
	for j < len(slc) {
		slc[k] = aux[j]
		j++
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
