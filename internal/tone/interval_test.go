package tone

import "testing"

func TestString_PositiveWholeTone(t *testing.T) {
	actual, _ := NewInterval(10)
	if actual.String() != "5" {
		t.Error("Expected: 5, but actual: " + actual.String())
	}
}

func TestString_PositiveHalfTone(t *testing.T) {
	actual, _ := NewInterval(11)
	if actual.String() != "5.5" {
		t.Error("Expected: 5.5, but actual: " + actual.String())
	}
}