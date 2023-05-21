package main

import "fmt"

func main() {
	xi := []int{6, 8, 4, 7}

	fmt.Println(quick(xi, 0, len(xi)-1))

}

func divide(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low

	fmt.Println("\tBefore Pivot: ", pivot, " | ", "Before I: ", i, " | ", "Before ARRAY: ", arr, "\n")
	for j := i; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	fmt.Println("After Pivot: ", pivot, " | ", "After I: ", i, " | ", "After ARRAY: ", arr, "\n")
	return arr, i
}

func quick(arr []int, low, high int) []int {
	if low > high {
		return arr
	}

	var p int
	arr, p = divide(arr, low, high)
	arr = quick(arr, low, p-1)
	arr = quick(arr, p+1, high)

	return arr
}
