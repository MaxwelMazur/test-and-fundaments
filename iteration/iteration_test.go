package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeats := repeat("a")
	expected := "aaaa"
	if repeats != expected {
		t.Errorf("expected '%s', repeats '%s'", expected, repeats)
	}
}
