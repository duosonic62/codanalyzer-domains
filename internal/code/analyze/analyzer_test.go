package analyze

import (
	"fmt"
	"github.com/duosonic62/codanalyzer-domains/internal/code"
	"github.com/duosonic62/codanalyzer-domains/internal/code/major"
	"github.com/duosonic62/codanalyzer-domains/internal/code/minor"
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
	"testing"
)

func TestAnalyzer_GetTriadWithContainedSameTonesInCode(t *testing.T) {
	//メジャートライアドのアナライザを作成
	analyuzer := Analyzer{triadCodeStructures: &[]code.TriadCodeStructure{
		major.NewMajorTriad(),
		minor.NewMinorTriad(),
	}}

	//CM7を解析する
	actual, err := analyuzer.GetTriadWithContainedSameTonesInCode(cM7())
	if err != nil {
		t.Error("Expected: error is nil, but ", err)
	}

	for _, v := range *actual {
		t.Log(v)
		fmt.Println(v)
	}
}

func cM7() *code.Code {
	return major.NewMajorMajorSeventh().GetCode(&tone.ScaleTones.C)
}