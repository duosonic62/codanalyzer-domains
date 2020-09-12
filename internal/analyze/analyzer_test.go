package analyze

import (
	"github.com/duosonic62/codanalyzer-domains/internal/code"
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
	"testing"
)

func TestAnalyzer_GetTriadWithContainedSameTonesInCode(t *testing.T) {
	//メジャートライアドのアナライザを作成
	analyzer := Analyzer{triadCodeStructures: &[]code.TriadCodeStructure{
		major(),
		minor(),
	}}

	//CM7を解析する
	actual, err := analyzer.GetTriadWithContainedSameTonesInCode(cM7())
	if err != nil {
		t.Error("Expected: error is nil, but ", err)
	}

	expected := []code.TriadCode{
		*major().GetTriadCode(&tone.ScaleTones.C),
		*minor().GetTriadCode(&tone.ScaleTones.E),
	}

	for i, v := range *actual {
		if v.Root == expected[i].Root {
			t.Error("Expected:", expected[i].Root, ", but ", v.Root)
		}
		if v.Tones == expected[i].Tones {
			t.Error("Expected:", expected[i].Tones, ", but ", v.Tones)
		}
		if v.Name == expected[i].Name {
			t.Error("Expected:", expected[i].Name, ", but ", v.Name)
		}
	}
}

func major() *code.TriadStructureBase {
	structure, _ := code.NewTriadStructureBase(
		"major",
		&[]tone.Interval{
			*wrapInterval(0),
			*wrapInterval(4),
			*wrapInterval(7),
		},
	)

	return structure
}

func minor() *code.TriadStructureBase {
	structure, _ := code.NewTriadStructureBase(
		"minor",
		&[]tone.Interval{
			*wrapInterval(0),
			*wrapInterval(3),
			*wrapInterval(7),
		},
	)

	return structure
}

func cM7() *code.Code {
	code, _ := code.NewCode(
		"CM7",
		&[]tone.Interval{
			*wrapInterval(0),
			*wrapInterval(4),
			*wrapInterval(7),
			*wrapInterval(11),
		},
		&tone.ScaleTones.C,
	)
	return code
}

func wrapInterval(interval int) *tone.Interval {
	i, _ := tone.NewInterval(interval)
	return i
}