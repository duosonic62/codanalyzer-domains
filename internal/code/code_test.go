package code

import (
	"fmt"
	"github.com/duosonic62/codanalyzer-domains/internal/analyze"
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
	"reflect"
	"strconv"
	"testing"
)

// コードの生成(正常系)
func TestNewCode_Positive(t *testing.T) {
	actual, err := NewCode("Major", majorTriadIntervals(), &tone.ScaleTones.C)
	if err != nil {
		t.Error(err)
	}

	expected := Code{
		Name:  NewCodeName(&tone.ScaleTones.C, "Major"),
		Root:  &tone.ScaleTones.C,
		Tones: tonesOnCMajorCode(),
	}
	if *actual.Name != *expected.Name {
		t.Error("Expected: "+expected.Name.Value+", but actual: "+actual.Name.Value, *actual.Tones, *actual.Root)
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

// コードの生成(異常系 構成音が3音ない場合)
func TestNewCode_Negative(t *testing.T) {
	// 構成音が3音ない場合
	tones := &[]tone.Interval{
		*tone.NewInterval(0),
		*tone.NewInterval(4),
	}
	_, err := NewCode("Major", tones, &tone.ScaleTones.C)
	if err == nil {
		t.Error("Error must not be nil")
	}
}

// トライアドコードの生成(正常系)
func TestNewTriadCode_Positive(t *testing.T) {
	actual, err := NewTriadCode("Major", majorTriadIntervals(), &tone.ScaleTones.C)
	if err != nil {
		t.Error(err)
	}

	expected := TriadCode{
		Name:  NewCodeName(&tone.ScaleTones.C, "Major"),
		Root:  &tone.ScaleTones.C,
		Tones: tonesOnCMajorCode(),
	}
	if *actual.Name != *expected.Name {
		t.Error("Expected: "+expected.Name.Value+", but actual: "+actual.Name.Value, *actual.Tones, *actual.Root)
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

// トライアドコードの生成(異常系 構成音が3音ない場合)
func TestNewTriadCode_Negative_UnderThreeTones(t *testing.T) {
	// 構成音が3音ない場合
	tones := &[]tone.Interval{
		*tone.NewInterval(0),
		*tone.NewInterval(4),
	}
	_, err := NewTriadCode("Major", tones, &tone.ScaleTones.C)
	if err == nil {
		t.Error("Error must not be nil")
	}
}

// トライアドコードの生成(異常系 構成音が4音以上の場合)
func TestNewTriadCode_Negative_OverThreeTones(t *testing.T) {
	// 構成音が3音ない場合
	tones := &[]tone.Interval{
		*tone.NewInterval(0),
		*tone.NewInterval(4),
		*tone.NewInterval(7),
		*tone.NewInterval(10),
	}
	_, err := NewTriadCode("7", tones, &tone.ScaleTones.C)
	if err == nil {
		t.Error("Error must not be nil")
	}
}

// TriadCode -> Codeへの変換を行うテスト
func TestTriadCode_Code_Positive(t *testing.T) {
	triad, err := NewTriadCode("Major", majorTriadIntervals(), &tone.ScaleTones.C)
	if err != nil {
		t.Error(err)
	}

	actual := triad.Code()
	expected := Code{
		Name:  NewCodeName(&tone.ScaleTones.C, "Major"),
		Root:  &tone.ScaleTones.C,
		Tones: tonesOnCMajorCode(),
	}
	if *actual.Name != *expected.Name {
		t.Error("Expected: "+expected.Name.Value+", but actual: "+actual.Name.Value, *actual.Tones, *actual.Root)
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

func TestCode_ExtractTriadPatterns(t *testing.T) {
	// CM7
	majorSeventh, _ := NewCode("Major Seventh", majorSeventhIntervals(), &tone.ScaleTones.C)
	actual, err := majorSeventh.ExtractTriadPatterns()
	if err != nil {
		t.Error("Expected: error is nil, but actual ", err)
	}
	if len(*actual) != 12 {
		t.Error("Expected: 4, but actual: " + strconv.Itoa(len(*actual)))
	}

	expected := []analyze.TriadInversion{
		// C, E, G
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
		// C, E, B
		{
			Root: &tone.ScaleTones.C,
			Intervals: &[]tone.Interval{
				*tone.ScaleTones.C.CalculateInterval(tone.ScaleTones.C),
				*tone.ScaleTones.C.CalculateInterval(tone.ScaleTones.B),
				*tone.ScaleTones.C.CalculateInterval(tone.ScaleTones.G),
			},
		},
		{
			Root: &tone.ScaleTones.E,
			Intervals: &[]tone.Interval{
				*tone.ScaleTones.E.CalculateInterval(tone.ScaleTones.E),
				*tone.ScaleTones.E.CalculateInterval(tone.ScaleTones.B),
				*tone.ScaleTones.E.CalculateInterval(tone.ScaleTones.C),
			},
		},
		{
			Root: &tone.ScaleTones.B,
			Intervals: &[]tone.Interval{
				*tone.ScaleTones.B.CalculateInterval(tone.ScaleTones.B),
				*tone.ScaleTones.B.CalculateInterval(tone.ScaleTones.C),
				*tone.ScaleTones.B.CalculateInterval(tone.ScaleTones.E),
			},
		},
		// C, G, B
		{
			Root: &tone.ScaleTones.C,
			Intervals: &[]tone.Interval{
				*tone.ScaleTones.C.CalculateInterval(tone.ScaleTones.C),
				*tone.ScaleTones.C.CalculateInterval(tone.ScaleTones.G),
				*tone.ScaleTones.C.CalculateInterval(tone.ScaleTones.B),
			},
		},
		{
			Root: &tone.ScaleTones.G,
			Intervals: &[]tone.Interval{
				*tone.ScaleTones.G.CalculateInterval(tone.ScaleTones.G),
				*tone.ScaleTones.G.CalculateInterval(tone.ScaleTones.B),
				*tone.ScaleTones.G.CalculateInterval(tone.ScaleTones.C),
			},
		},
		{
			Root: &tone.ScaleTones.B,
			Intervals: &[]tone.Interval{
				*tone.ScaleTones.B.CalculateInterval(tone.ScaleTones.B),
				*tone.ScaleTones.B.CalculateInterval(tone.ScaleTones.C),
				*tone.ScaleTones.B.CalculateInterval(tone.ScaleTones.G),
			},
		},
		// E, G, B
		{
			Root: &tone.ScaleTones.E,
			Intervals: &[]tone.Interval{
				*tone.ScaleTones.E.CalculateInterval(tone.ScaleTones.E),
				*tone.ScaleTones.E.CalculateInterval(tone.ScaleTones.G),
				*tone.ScaleTones.E.CalculateInterval(tone.ScaleTones.B),
			},
		},
		{
			Root: &tone.ScaleTones.G,
			Intervals: &[]tone.Interval{
				*tone.ScaleTones.G.CalculateInterval(tone.ScaleTones.G),
				*tone.ScaleTones.G.CalculateInterval(tone.ScaleTones.B),
				*tone.ScaleTones.G.CalculateInterval(tone.ScaleTones.E),
			},
		},
		{
			Root: &tone.ScaleTones.B,
			Intervals: &[]tone.Interval{
				*tone.ScaleTones.B.CalculateInterval(tone.ScaleTones.B),
				*tone.ScaleTones.B.CalculateInterval(tone.ScaleTones.E),
				*tone.ScaleTones.B.CalculateInterval(tone.ScaleTones.G),
			},
		},
	}

	for _, v := range *actual  {
		fmt.Println(*v.Intervals)
	}

	if reflect.DeepEqual(actual, expected) {
		t.Error("Expected: ", expected, ", but ", actual)
	}
}

// テストヘルパーメソッド
func majorTriadIntervals() *[]tone.Interval {
	return &[]tone.Interval{
		*tone.NewInterval(0),
		*tone.NewInterval(4),
		*tone.NewInterval(7),
	}
}
func majorSeventhIntervals() *[]tone.Interval {
	return &[]tone.Interval{
		*tone.NewInterval(0),
		*tone.NewInterval(4),
		*tone.NewInterval(7),
		*tone.NewInterval(11),
	}
}

// テストヘルパーメソッド
func tonesOnCMajorCode() *[]tone.ScaleTone {
	return &[]tone.ScaleTone{tone.ScaleTones.C, tone.ScaleTones.E, tone.ScaleTones.G}
}
