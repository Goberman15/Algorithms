package main

import "fmt"

func main() {
	slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	idx := binarySearch(slc, 1)
	fmt.Println(idx)
}

func binarySearch(slc []int, target int) int {
	top, bot := len(slc)-1, 0

	for top >= bot {
		mid := (top + bot) / 2

		if target == slc[mid] {
			return mid
		} else if target < slc[mid] {
			top = mid - 1
		} else {
			bot = mid + 1
		}
	}

	return -1
}
