package tone

import (
	"strconv"
	"testing"
)

// 比較する音が元の音より音階では上の場合
// C -> D
func TestCalculateInterval_Positive_SameOctave(t *testing.T) {
	actual, _ := ScaleTones.C.CalculateInterval(ScaleTones.D)
	if actual.value != 2 {
		t.Error("Expected: 2, but actual: " + strconv.Itoa(actual.value))
	}
}

// 比較する音が元の音より音階では上の場合
// G -> C
func TestCalculateInterval_Positive_UpperOctave(t *testing.T) {
	actual, _ := ScaleTones.G.CalculateInterval(ScaleTones.C)
	if actual.value != 5 {
		t.Error("Expected: 5, but actual: " + strconv.Itoa(actual.value))
	}
}

// 比較する音が元の音と同じ
// C -> C
func TestScaleTone_CalculateInterval_SameTone(t *testing.T) {
	actual, _ := ScaleTones.C.CalculateInterval(ScaleTones.C)
	if actual.value != 0 {
		t.Error("Expected: 0, but actual: " + strconv.Itoa(actual.value))
	}
}

// 取得する音が元の音と同じオクターブ
func TestScaleTone_GetToneWithApartInterval_Positive_SameOctave(t *testing.T) {
	// Cから二音離れた音(D)を取得する
	actual := ScaleTones.C.GetToneWithApartInterval(wrapInterval(2))
	if *actual != ScaleTones.D {
		t.Error("Expected: D, but actual: " + actual.String())
	}
}

// 取得する音が基音のオクターブより上
func TestScaleTone_GetToneWithApartInterval_Positive_UpperOctave(t *testing.T) {
	// Gから5音離れた音(C)を取得する
	actual := ScaleTones.G.GetToneWithApartInterval(wrapInterval(5))
	if *actual != ScaleTones.C {
		t.Error("Expected: C, but actual: " + actual.String())
	}

	// Gから7音離れた音(D)を取得する
	actual = ScaleTones.G.GetToneWithApartInterval(wrapInterval(7))
	if *actual != ScaleTones.D {
		t.Error("Expected: D, but actual: " + actual.String())
	}
}

func wrapInterval(interval int) *Interval {
	i, _ := NewInterval(interval)
	return i
}