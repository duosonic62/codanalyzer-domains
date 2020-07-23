package analyze

import "github.com/duosonic62/codanalyzer-domains/internal/code"

type Analyzer struct {

}

//GetTriadWithContainedSameTonesInCode は指定されたコード内に含まれる音と同じ構成のトライアドを取得する
func (a Analyzer) GetTriadWithContainedSameTonesInCode(code *code.Code) *[]code.TriadCode {

}