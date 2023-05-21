package main

import "testing"

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort([]int{277, 10, 6, 9, 2, 5, 8, 3, 4, 7, 88, 388, 30})

	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectionSort([]int{22, 10, 6, 2, 1, 5, 8, 3, 4, 7, 9, 100, 30})
	}
}
