package major

import (
	"github.com/duosonic62/codanalyzer-domains/internal/code"
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
	"testing"
)

var triad = NewMajorTriad()

func TestMajorTriad_Name_Positive(t *testing.T) {
	if triad.Name() != "Major Triad" {
		t.Error("Expected: Major Triad, but actual: " + triad.Name())
	}
}

func TestMajorTriad_GetIntervals_Positive(t *testing.T) {
	expected := []tone.Interval{
		*tone.NewInterval(0),
		*tone.NewInterval(4),
		*tone.NewInterval(7),
	}

	for i, actual := range *triad.intevals {
		if actual != expected[i] {
			t.Error("Expected: ", expected[i], ", but actual: ", actual)
		}
	}
}

func TestMajorTriad_GetCode_Positive(t *testing.T) {
	actual := triad.GetCode(&tone.ScaleTones.C)
	expected, err := code.NewCode("Major Triad", majorTriadIntervals(), &tone.ScaleTones.C)
	if err != nil {
		t.Error(err)
	}

	if *actual.Name != *expected.Name {
		t.Error("Expected: " + expected.Name.Value + ", but actual: " + actual.Name.Value, *actual.Tones, *actual.Root)
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
func majorTriadIntervals() *[]tone.Interval {
	return &[]tone.Interval {
		*tone.NewInterval(0),
		*tone.NewInterval(4),
		*tone.NewInterval(7),
	}
}

// テストヘルパーメソッド
func tonesOnCMajorCode() *[]tone.ScaleTone {
	return &[]tone.ScaleTone{tone.ScaleTones.C, tone.ScaleTones.E, tone.ScaleTones.G}
}