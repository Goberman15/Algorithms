package main

import "fmt"

func main() {
	// generate unsorted integer slice with 30 elements
	slc := []int{29, 10, 14, 37, 14, 43, 6, 15, 52, 53, 23, 21, 66, 32, 11, 52, 87, 93, 4, 225, 212, 5, 44, 3}
	quickSort(slc)

	fmt.Println(slc)

}

func quickSort(slc []int) {
	if len(slc) < 2 {
		return
	}

	swap := -1
	pvt := len(slc) - 1

	for x := 0; x <= pvt; x++ {
		if slc[x] > slc[pvt] {
			continue
		} else {
			swap++

			if x > swap {
				slc[x], slc[swap] = slc[swap], slc[x]
			} else {
				continue
			}
		}
	}

	fmt.Println(slc[:swap])
	fmt.Println(slc[swap:])

	high := swap + 1

	if high > len(slc)-1 {
		high = len(slc) - 1
	}

	quickSort(slc[:swap])
	quickSort(slc[high:])
}
