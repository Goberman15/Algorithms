package main

import "fmt"

func main() {
	slc := []int{29, 10, 14, 37, 14, 43, 6, 15, 52, 53, 23, 21, 66, 32, 11, 52, 87, 93, 4, 225, 212, 5, 44, 3}
	quickSort(slc)

	fmt.Println(slc)

}

func quickSort(slc []int) {
	if len(slc) < 2 {
		return
	}

	swap := len(slc)
	pvt := 0

	for x := len(slc) - 1; x >= pvt; x-- {
		if slc[x] >= slc[pvt] {
			swap--
			if x < swap {
				slc[x], slc[swap] = slc[swap], slc[x]
			}
		}
	}

	high := swap + 1

	if high >= len(slc) {
		high = len(slc) - 1
	}

	quickSort(slc[:swap])
	quickSort(slc[high:])
}
