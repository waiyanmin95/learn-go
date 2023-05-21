package main

import (
	"testing"
)

func TestMySum(t *testing.T) {
	x := mySum(2, 3)
	if x != 5 {
		t.Errorf("Excepted 5 and got %v", x)
	}
	x = mySum(5, 9)
	if x != 14 {
		t.Errorf("Excepted 14 and got %v", x)
	}
	x = mySum(10, 1)
	if x != 11 {
		t.Errorf("Excepted 11 and got %v", x)
	}
}
