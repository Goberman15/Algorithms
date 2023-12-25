package main

import "fmt"

func main() {
	slc := []int{23, 12, 13, 53, 6, 4, 2, 16, 109, 84, 25, 35, 33, 48, 82, 67}
	insertionSort(slc)

	fmt.Println(slc)
}

func insertionSort(slc []int) {
	for x := 1; x < len(slc); x++ {
		sVal := slc[x]
		sIdx := x

		for y := x - 1; y >= 0; y-- {
			if sVal < slc[y] {
				slc[y+1] = slc[y]
				sIdx--
			} else {
				break
			}
		}
		slc[sIdx] = sVal
	}
}
