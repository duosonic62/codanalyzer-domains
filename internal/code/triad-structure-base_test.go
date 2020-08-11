package code

import (
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
	"testing"
)

var triad, _ = NewTriadStructureBase("Major Triad", majorTriadIntervals())

func TestNewTriadStructureBase_Positive(t *testing.T) {
	_, err := NewTriadStructureBase("Major Triad", majorTriadIntervals())
	if err != nil {
		t.Error("Expected error is nil but, actual: ", err)
	}
}

func TestNewTriadStructureBase_Negative(t *testing.T) {
	_, err := NewTriadStructureBase("Major Triad", majorSeventhIntervals())
	if err == nil {
		t.Error("Expected error not nil but, actual: ", err)
	}
}

func TestTriadStructureBase_Name_Positive(t *testing.T) {
	if triad.Name() != "Major Triad" {
		t.Error("Expected: Major Triad, but actual: " + triad.Name())
	}
}

func TestTriadStructureBase_GetCode_Positive(t *testing.T) {
	actual := triad.GetCode(&tone.ScaleTones.C)
	expected, err := NewCode("Major Triad", majorTriadIntervals(), &tone.ScaleTones.C)
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

func TestMajorTriad_GetTriadCode_Positive(t *testing.T) {
	actual := triad.GetTriadCode(&tone.ScaleTones.C)
	expected, err := NewTriadCode("Major Triad", majorTriadIntervals(), &tone.ScaleTones.C)
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