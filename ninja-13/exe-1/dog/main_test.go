package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	a := Years(3)
	if a != 21 {
		t.Error("Expected", a, "Got", 21)
	}
}

func TestYearsTwo(t *testing.T) {
	a := YearsTwo(8)
	if a != 56 {
		t.Error("Expected", a, "Got", 56)
	}
}

func ExampleYears() {
	fmt.Println(Years(3))
	// Outputs:
	// 21
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(8))
	// Outputs:
	// 56
}

func BenchmarkYears(b *testing.B) {
	Years(10)
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(20)
	}
}
