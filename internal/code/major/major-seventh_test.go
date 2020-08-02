package major

import (
	"github.com/duosonic62/codanalyzer-domains/internal/code"
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
	"testing"
)

var majorSeventh = NewMajorSeventh()

func TestMajorSeventh_Name_Positive(t *testing.T) {
	if majorSeventh.Name() != "Major MajorSeventh" {
		t.Error("Expected: Major MajorSeventh, but actual: " + majorSeventh.Name())
	}
}

func TestMajorSeventh_GetIntervals_Positive(t *testing.T) {
	expected := *majorSeventhIntervals()

	for i, actual := range *majorSeventh.intervals {
		if actual != expected[i] {
			t.Error("Expected: ", expected[i], ", but actual: ", actual)
		}
	}
}

func TestMajorSeventh_GetCode_Positive(t *testing.T) {
	actual := majorSeventh.GetCode(&tone.ScaleTones.C)
	expected, err := code.NewCode("Major MajorSeventh", majorSeventhIntervals(), &tone.ScaleTones.C)
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
func majorSeventhIntervals() *[]tone.Interval {
	return &[]tone.Interval{
		*tone.NewInterval(0),
		*tone.NewInterval(4),
		*tone.NewInterval(7),
		*tone.NewInterval(11),
	}
}