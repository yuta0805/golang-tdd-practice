package arrays

import "testing"

func TestSum(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	sum := Sum(arr)
	expected := 15
	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
} 
