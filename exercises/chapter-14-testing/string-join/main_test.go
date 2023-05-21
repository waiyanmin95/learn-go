package main

import "testing"

type testData struct {
	data []string
	want string
}

func TestJoinWithComma(t *testing.T) {
	tests := []testData{
		{
			data: []string{},
			want: "Please make an input",
		},
		{
			data: []string{"apple"},
			want: "apple",
		},
		{
			data: []string{"apple", "orange"},
			want: "apple and orange",
		},
		{
			data: []string{"apple", "orange", "banana", "grape"},
			want: "apple, orange, banana, and grape",
		},
	}

	for _, v := range tests {
		got := JoinWithComma(v.data)
		if got != v.want {
			t.Errorf("Expected %v but got %v", v.want, got)
		}
	}
}
