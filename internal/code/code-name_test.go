package code

import (
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
	"testing"
)

func TestNewCodeName_Positive(t *testing.T) {
	actual := NewCodeName(&tone.ScaleTones.C, "Major")
	if actual.Value != "C Major" {
		t.Error("Expected: C Major, but actual: " + actual.Value)
	}
}