package main

import (
	"testing"
)

// benchmark the mergesort function
func BenchmarkMergeSort(b *testing.B) {
	slc := []int{29, 10, 14, 37, 14, 43, 6, 15, 52, 53, 23, 21, 66, 32, 11, 52, 87, 93, 4, 225, 212, 5, 44, 3}
	for i := 0; i < b.N; i++ {
		mergeSort(slc)
	}
}
