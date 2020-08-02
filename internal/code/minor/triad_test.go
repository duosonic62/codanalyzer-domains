package minor

import (
	"github.com/duosonic62/codanalyzer-domains/internal/code"
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
	"testing"
)

var triad = NewTriad()

func TestMinorTriad_Name_Positive(t *testing.T) {
	if triad.Name() != "Minor Triad" {
		t.Error("Expected: Minor Triad, but actual: " + triad.Name())
	}
}

func TestMinorTriad_GetIntervals_Positive(t *testing.T) {
	expected := *minorTriadIntervals()

	for i, actual := range *triad.intervals {
		if actual != expected[i] {
			t.Error("Expected: ", expected[i], ", but actual: ", actual)
		}
	}
}

func TestMinorTriad_GetCode(t *testing.T) {
	actual := triad.GetCode(&tone.ScaleTones.C)
	expected, err := code.NewCode("Minor Triad", minorTriadIntervals(), &tone.ScaleTones.C)
	if err != nil {
		t.Error(err)
	}

	if *actual.Name != *expected.Name {
		t.Error("Expected: " + expected.Name.Value + ", but actual: " + actual.Name.Value)
	}

	if *actual.Root != *expected.Root {
		t.Error("Expected: " + expected.Root.String() + ", but actual: " + actual.Root.String())
	}

	for i, actualTone := range *actual.Tones {
		expectedTones := *expected.Tones
		if actualTone != expectedTones[i] {
			t.Error("Expected: " + expectedTones[i].String() + ", but actual: " + actualTone.String())
		}
	}
}

// テストヘルパーメソッド
func minorTriadIntervals() *[]tone.Interval {
	return &[]tone.Interval {
		*tone.NewInterval(0),
		*tone.NewInterval(3),
		*tone.NewInterval(7),
	}
}