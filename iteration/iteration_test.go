package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeats := repeat("a")
	expected := "aaaa"
	if repeats != expected {
		t.Errorf("expected '%s', repeats '%s'", expected, repeats)
	}
}

func TestReturnNumber(t *testing.T) {
	result := returnNumber()
	if result != 6 {
		t.Errorf("expected 6, result '%s'", result)
	}
}
