package word

import (
	"fmt"
	"github.com/ivan-amity/learn-go/ninja-13/exe-2/quote"
	"testing"
)

func TestUseCount(t *testing.T) {
	m := UseCount("one two two three three three")
	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "two":
			if v != 2 {
				t.Error("got", v, "want", 2)
			}
		case "three":
			if v != 3 {
				t.Error("got", v, "want", 3)
			}
		}
	}
}

func TestCount(t *testing.T) {
	n := Count("one two three five")
	if n != 4 {
		t.Error("got", n, "want", 4)
	}
}

func ExampleCount() {
	fmt.Println(Count("one two three"))
	// Output:
	// 3
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
