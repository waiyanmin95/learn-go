package arithmetic

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	want := 4.0
	plus := Add(1, 3)
	if plus != want {
		t.Errorf(helper(want, plus))
	}
}

func TestSubtract(t *testing.T) {
	want := 2.0
	minus := Subtract(4, 2)
	if minus != want {
		t.Errorf(helper(want, minus))
	}
}

func helper(want float64, got float64) string {
	return fmt.Sprintf("Expected %v, but got %v", want, got)
}
