package code

import "github.com/duosonic62/codanalyzer-domains/internal/tone"

type CodeName struct {
	Value string
}

//NewCodeName コード名を生成する
func NewCodeName(root *tone.ScaleTone, structureName string) *CodeName {
	name := root.String() + " " + structureName
	return &CodeName{Value: name}
}
