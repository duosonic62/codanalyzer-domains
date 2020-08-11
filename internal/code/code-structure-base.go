package code

import (
	"errors"
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
)

type StructureBase struct {
	name      string
	intervals *[]tone.Interval
}

//NewCodeStructureBase はStructureBaseのインスタンスを生成します
func NewCodeStructureBase(codeBaseName string, intervals *[]tone.Interval) (*StructureBase, error) {
	if len(*intervals) < 3 {
		return nil, errors.New("need at least 3 chord tones")
	}

	return &StructureBase{
		name:      codeBaseName,
		intervals: intervals,
	}, nil
}

//Name はコード名を表示します
func (sb StructureBase) Name() string {
	return sb.name
}

//GetIntervals はコード構成音の間隔を取得します
func (sb StructureBase) GetIntervals() *[]tone.Interval {
	return sb.intervals
}

//GetCode は指定されたルート音でのコード構成を返します
func (sb StructureBase) GetCode (root *tone.ScaleTone) *Code {
	code, err := NewCode(sb.name, sb.intervals, root)

	// 本来エラーは起こり得ないので、エラーになったらpanicを起こす
	if err != nil {
		panic(err)
	}

	return code
}
