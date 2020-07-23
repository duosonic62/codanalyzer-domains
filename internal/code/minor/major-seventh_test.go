package minor

import (
	"github.com/duosonic62/codanalyzer-domains/internal/code"
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
	"testing"
)

var minorMajorSeventh = NewMinorMajorSeventh()

func TestMajorSeventh_Name_Positive(t *testing.T) {
	if minorMajorSeventh.Name() != "Major Seventh" {
		t.Error("Expected: Major Seventh, but actual: " + minorMajorSeventh.Name())
	}
}

func TestMajorSeventh_GetIntervals_Positive(t *testing.T) {
	expected := *minorMajorSeventhIntervals()

	for i, actual := range *minorMajorSeventh.intervals {
		if actual != expected[i] {
			t.Error("Expected: ", expected[i], ", but actual: ", actual)
		}
	}
}

func TestMajorSeventh_GetCode_Positive(t *testing.T) {
	actual := minorMajorSeventh.GetCode(&tone.ScaleTones.C)
	expected, err := code.NewCode("Major Seventh", minorMajorSeventhIntervals(), &tone.ScaleTones.C)
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
func minorMajorSeventhIntervals() *[]tone.Interval {
	return &[]tone.Interval{
		*tone.NewInterval(0),
		*tone.NewInterval(3),
		*tone.NewInterval(7),
		*tone.NewInterval(11),
	}
}