package utils

import (
	"slices"
	"testing"
)

func TestReplace(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 11, 11}
	expected := []int{1, 2, 3, 4, 100, 100, 100, 100, 9, 0, 11, 11, 11}
	Replace(&input, 4, 4, 100)
	if !slices.Equal(input, expected) {
		t.Errorf("Slices not equal\ninp: %v\nexp: %v", input, expected)
	}
	expected = []int{1, 2, 3, 55, 55, 55, 55, 55, 55, 55, 11, 11, 11}
	Replace(&input, 3, 7, 55)
	if !slices.Equal(input, expected) {
		t.Errorf("Slices not equal\ninp: %v\nexp: %v", input, expected)
	}
}
