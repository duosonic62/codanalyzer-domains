package analyze

import (
	"github.com/duosonic62/codanalyzer-domains/internal/code"
	"github.com/duosonic62/codanalyzer-domains/internal/code/major"
	"github.com/duosonic62/codanalyzer-domains/internal/code/minor"
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
	"testing"
)

func TestAnalyzer_GetTriadWithContainedSameTonesInCode(t *testing.T) {
	//メジャートライアドのアナライザを作成
	analyzer := Analyzer{triadCodeStructures: &[]code.TriadCodeStructure{
		major.NewTriad(),
		minor.NewTriad(),
	}}

	//CM7を解析する
	actual, err := analyzer.GetTriadWithContainedSameTonesInCode(cM7())
	if err != nil {
		t.Error("Expected: error is nil, but ", err)
	}

	expected := []code.TriadCode{
		*major.NewTriad().GetTriadCode(&tone.ScaleTones.C),
		*minor.NewTriad().GetTriadCode(&tone.ScaleTones.E),
	}

	for i, v := range *actual {
		if v.Root == expected[i].Root {
			t.Error("Expected:", expected[i].Root,  ", but ", v.Root)
		}
		if v.Tones == expected[i].Tones {
			t.Error("Expected:", expected[i].Tones,  ", but ", v.Tones)
		}
		if v.Name == expected[i].Name {
			t.Error("Expected:", expected[i].Name,  ", but ", v.Name)
		}
	}
}

func cM7() *code.Code {
	return major.NewMajorSeventh().GetCode(&tone.ScaleTones.C)
}