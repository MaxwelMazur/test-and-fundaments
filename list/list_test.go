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

func TestSumAllRest(t *testing.T) {
	t.Run("get the rest collection", func(t *testing.T) {
		result := sumAllRest([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("result %v expected %v", result, expected)
		}
	})

	t.Run("get the rest collection and set zero in list empty", func(t *testing.T) {
		result := sumAllRest([]int{}, []int{2, 12})
		expected := []int{0, 12}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("result %v expected %v", result, expected)
		}
	})
}
