package analyze

import (
	"github.com/duosonic62/codanalyzer-domains/internal/code"
	"reflect"
)

type Analyzer struct {
	triadCodeStructures *[]code.TriadCodeStructure
}

//NewTriadCodeAnalyzer はコードアナライザーのインスタンスを生成します
func NewTriadCodeAnalyzer() *Analyzer {
	//トライアドコードはメジャー、マイナー、オーギュメント、ディミニッシュ
	triadCodeStructures := make([]code.TriadCodeStructure, 4)

	return &Analyzer{triadCodeStructures: &triadCodeStructures}
}

//GetTriadWithContainedSameTonesInCode は指定されたコード内に含まれる音と同じ構成のトライアドを取得します
func (a Analyzer) GetTriadWithContainedSameTonesInCode(target *code.Code) (*[]code.TriadCode, error) {
	triadCodes := make([]code.TriadCode, 0)
	triadPatterns, err := target.ExtractTriadPatterns()
	if err != nil {
		return nil, err
	}

	for _, triadPattern := range *triadPatterns {
		for _, structure := range *a.triadCodeStructures {
			codeOrNil := getEquivalentCodeOrNil(&triadPattern, &structure)
			if codeOrNil != nil {
				triadCodes = append(triadCodes, *codeOrNil)
			}
		}
	}

	return &triadCodes, nil
}

//getEquivalentCodeOrNil はTriadPatternに等価なコードがあれば取得するメソッドです
func getEquivalentCodeOrNil(pattern *code.TriadInversion, structure *code.TriadCodeStructure) *code.TriadCode {
	s := *structure
	//インターバルが等価であればコードを取得する
	if reflect.DeepEqual(*pattern.Intervals, *s.GetIntervals()) {
		return s.GetTriadCode(pattern.Root)
	}

	return nil
}