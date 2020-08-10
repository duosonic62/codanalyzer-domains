package code

import (
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
	"reflect"
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
				*tone.ScaleTones.C.CalculateInterval(tone.ScaleTones.C),
				*tone.ScaleTones.C.CalculateInterval(tone.ScaleTones.E),
				*tone.ScaleTones.C.CalculateInterval(tone.ScaleTones.G),
			},
		},
		{
			Root: &tone.ScaleTones.E,
			Intervals: &[]tone.Interval{
				*tone.ScaleTones.E.CalculateInterval(tone.ScaleTones.E),
				*tone.ScaleTones.E.CalculateInterval(tone.ScaleTones.G),
				*tone.ScaleTones.E.CalculateInterval(tone.ScaleTones.C),
			},
		},
		{
			Root: &tone.ScaleTones.G,
			Intervals: &[]tone.Interval{
				*tone.ScaleTones.G.CalculateInterval(tone.ScaleTones.G),
				*tone.ScaleTones.G.CalculateInterval(tone.ScaleTones.C),
				*tone.ScaleTones.G.CalculateInterval(tone.ScaleTones.E),
			},
		},
	}


	for i, v := range *actual {
		if !reflect.DeepEqual(v, expected[i]) {
			t.Error("Expected: ", *v.Intervals, ", but ", *expected[i].Intervals)
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
