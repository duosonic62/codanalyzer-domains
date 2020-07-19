package tone

import (
	"strconv"
	"testing"
)

// 比較する音が元の音より音階では上の場合
// C -> D
func TestCalculateInterval_Positive_UnderTone(t *testing.T) {
	actual := ScaleTones.C.CalculateInterval(ScaleTones.D)
	if actual.value != 2 {
		t.Error("Expected: 2, but actual: " + strconv.Itoa(actual.value))
	}
}

// 比較する音が元の音より音階では上の場合
// G -> C
func TestCalculateInterval_Positive_OverTone(t *testing.T) {
	actual := ScaleTones.G.CalculateInterval(ScaleTones.C)
	if actual.value != 5 {
		t.Error("Expected: 5, but actual: " + strconv.Itoa(actual.value))
	}
}