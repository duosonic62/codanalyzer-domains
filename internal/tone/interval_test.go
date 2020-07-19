package tone

import "testing"

func TestString_PositiveWholeTone(t *testing.T) {
	actual := NewInterval(10).String()
	if actual != "5" {
		t.Error("Expected: 5, but actual: " + actual)
	}
}

func TestString_PositiveHalfTone(t *testing.T) {
	actual := NewInterval(11).String()
	if actual != "5.5" {
		t.Error("Expected: 5.5, but actual: " + actual)
	}
}