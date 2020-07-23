package code

import (
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
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
	actual := majorSeventh.ExtractTriadPatterns()
	if len(*actual) != 4 {
		t.Error("Expected: 4, but actual: " + strconv.Itoa(len(*actual)))
	}

	expected := [][]tone.ScaleTone {
		{tone.ScaleTones.C, tone.ScaleTones.E, tone.ScaleTones.G},
		{tone.ScaleTones.C, tone.ScaleTones.E, tone.ScaleTones.B},
		{tone.ScaleTones.C, tone.ScaleTones.G, tone.ScaleTones.B},
		{tone.ScaleTones.E, tone.ScaleTones.G, tone.ScaleTones.B},
	}

	for i, tones := range *actual{
		for j, tone := range tones{
			if tone != expected[i][j] {
				t.Error("Expected: " + expected[i][j].String() + ", but actual: " + tone.String())
			}
		}
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
