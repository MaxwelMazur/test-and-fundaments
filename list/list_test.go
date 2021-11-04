package list

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	result := sum(numbers)
	expected := 15

	if expected != result {
		t.Errorf("result %d, expected %d, dado %v", result, expected, numbers)
	}
}
