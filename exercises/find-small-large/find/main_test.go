package find

import (
	"fmt"
	"testing"
)

func TestFindsmall(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	blah := []test{
		{[]int{6, 7, 44, 7, 88, 5, 66}, 5},
		{[]int{88, 99, 44, 99, 88, 7, 66}, 7},
		{[]int{100, 99, 100, 99, 88, 34, 66}, 34},
	}

	for _, v := range blah {
		a := Findsmall(v.data)
		if a != v.answer {
			t.Error("Expected", v.answer, "Got", a)
		}
	}
}

func TestFindlarge(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	blah := []test{
		{[]int{6, 7, 44, 7, 88, 5, 66}, 88},
		{[]int{88, 99, 44, 99, 88, 7, 66}, 99},
		{[]int{100, 99, 200, 99, 88, 34, 66}, 200},
	}

	for _, v := range blah {
		a := Findlarge(v.data)
		if a != v.answer {
			t.Error("Expected", v.answer, "Got", a)
		}
	}
}

func ExampleFindsmall() {
	fmt.Println(Findsmall([]int{6, 7, 44, 7, 88, 5, 66}))
	// Output:
	// 5
}

func ExampleFindlarge() {
	fmt.Println(Findlarge([]int{6, 7, 44, 7, 88, 5, 66}))
	// Output:
	// 88
}

func BenchmarkFindsmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Findsmall([]int{6, 7, 44, 7, 88, 5, 66})
	}
}

func BenchmarkFindlarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Findlarge([]int{6, 7, 44, 7, 88, 5, 66})
	}
}
