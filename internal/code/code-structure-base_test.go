package code

import (
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
	"testing"
)

var majorSeventh, _ = NewCodeStructureBase("M7", majorSeventhIntervals())

func TestNewCodeStructureBase_Positive(t *testing.T) {
	_, err := NewCodeStructureBase("M7", majorSeventhIntervals())
	if err != nil {
		t.Error("Expected error is nil but, actual: ", err)
	}
}

func TestNewCodeStructureBase_Negative(t *testing.T) {
	intervals := []tone.Interval{
		*wrapInterval(0),
		*wrapInterval(1),
	}

	//コード構成音が2音以下
	_, err := NewCodeStructureBase("M7", &intervals)
	if err == nil {
		t.Error("Expected error not nil but, actual: ", err)
	}
}

func TestStructureBase_Name_Positive(t *testing.T) {
	if majorSeventh.Name() != "M7" {
		t.Error("Expected M7 but, actual: ", majorSeventh.Name())
	}
}


func TestStructureBase_GetIntervals_Positive(t *testing.T) {
	expected := *majorSeventhIntervals()

	for i, actual := range *majorSeventh.intervals {
		if actual != expected[i] {
			t.Error("Expected: ", expected[i], ", but actual: ", actual)
		}
	}
}

func TestStructureBase_GetCode_Positive(t *testing.T) {
	actual := majorSeventh.GetCode(&tone.ScaleTones.C)
	expected, err := NewCode("M7", majorSeventhIntervals(), &tone.ScaleTones.C)
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

func wrapInterval(interval int) *tone.Interval {
	i, _ := tone.NewInterval(interval)
	return i
}