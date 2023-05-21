package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		value []int
		ans   float64
	}

	tests := []test{
		test{[]int{2, 4, 6, 8, 100}, 6},
		test{[]int{2, 8, 8, 8, 100, 800}, 31},
		test{[]int{2, 99, 100, 9, 100}, 69.33333333333333},
	}

	for _, v := range tests {
		f := CenteredAvg(v.value)
		if f != v.ans {
			t.Error("Expected", f, "Got", v.ans)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{9000, 4, 10, 8, 6, 12}))
	// Outputs:
	// 9
}
func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{9000, 4, 10, 8, 6, 12})
	}
}
