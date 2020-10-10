package code

import (
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
	"testing"
)

func TestNewTriadPattern_Positive(t *testing.T) {
	actual, err := NewTriadInversions(cTriad())
	if err != nil {
		t.Error("Expected: error is nil, but ", err)
	}

	expected := []TriadInversion{
		{
			Root: &tone.ScaleTones.C,
			Intervals: &[]tone.Interval{
				*calcInterval(tone.ScaleTones.C, tone.ScaleTones.C),
				*calcInterval(tone.ScaleTones.C, tone.ScaleTones.E),
				*calcInterval(tone.ScaleTones.C, tone.ScaleTones.G),
			},
		},
		{
			Root: &tone.ScaleTones.E,
			Intervals: &[]tone.Interval{
				*calcInterval(tone.ScaleTones.E, tone.ScaleTones.E),
				*calcInterval(tone.ScaleTones.E, tone.ScaleTones.G),
				*calcInterval(tone.ScaleTones.E, tone.ScaleTones.C),
			},
		},
		{
			Root: &tone.ScaleTones.G,
			Intervals: &[]tone.Interval{
				*calcInterval(tone.ScaleTones.G, tone.ScaleTones.G),
				*calcInterval(tone.ScaleTones.G, tone.ScaleTones.C),
				*calcInterval(tone.ScaleTones.G, tone.ScaleTones.E),
			},
		},
	}

	for i, v := range *actual {
		expectedIntervals := *expected[i].Intervals
		for j, interval := range *v.Intervals {
			if !interval.IsEquivalent(&expectedIntervals[j]) {
				t.Error("Expected: ", interval, ", but ", expectedIntervals[j])
			}
		}
	}
}

func TestNewTriadPattern_Negative(t *testing.T) {
	// 2音以下でエラー
	_, err := NewTriadInversions(&[]tone.ScaleTone{tone.ScaleTones.C, tone.ScaleTones.E})
	if err == nil {
		t.Error("Expected: error is not nil")
	}

	// 4音以上でエラー
	_, err = NewTriadInversions(&[]tone.ScaleTone{tone.ScaleTones.C, tone.ScaleTones.E, tone.ScaleTones.F, tone.ScaleTones.G})
	if err == nil {
		t.Error("Expected: error is not nil")
	}
}

func cTriad() *[]tone.ScaleTone {
	triad := []tone.ScaleTone{tone.ScaleTones.C, tone.ScaleTones.E, tone.ScaleTones.G}
	return &triad
}
