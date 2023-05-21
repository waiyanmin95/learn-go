package main

import (
	"testing"
)

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{[]int{22, 22}, 44},
		{[]int{8, 9}, 17},
		{[]int{2, 8, 0}, 10},
		{[]int{99, 100, 1}, 200},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Errorf("Expected answer %v but got %v", v.answer, x)
		}
	}
}
