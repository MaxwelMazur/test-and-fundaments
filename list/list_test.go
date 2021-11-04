package list

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection in five numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := sum(numbers)
		expected := 15

		if expected != result {
			t.Errorf("result %d, expected %d, dado %v", result, expected, numbers)
		}
	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := sum(numbers)
		expected := 6

		if expected != result {
			t.Errorf("result %d, expected %d, dado %v", result, expected, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	result := sumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result %v expected %v", result, expected)
	}
}
