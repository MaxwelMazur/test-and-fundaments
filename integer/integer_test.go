package integer

import "testing"

func TestSum(t *testing.T) {
	got := sum(2, 2)
	want := 4

	if got != want {
		t.Errorf("want '%d', got '%d'", want, got)
	}
}
