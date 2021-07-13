package scale

import "testing"

func TestInterval_IsEquivalent(t *testing.T) {
	a := Intervals.Natural9
	b := Intervals.Major2

	if !a.IsEquivalent(&b) {
		t.Error("Expected: 2 and 9 is Equivalent")
	}

	c := Intervals.R
	d := Intervals.Major2

	if c.IsEquivalent(&d) {
		t.Error("Expected: R and 2 is not Equivalent")
	}
}