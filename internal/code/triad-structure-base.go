package code

import (
	"errors"
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
)

type TriadStructureBase struct {
	name      string
	intervals *[]tone.Interval
}

//NewTriadStructureBase はTriadStructureBaseのインスタンスを生成します
func NewTriadStructureBase(name string, intervals *[]tone.Interval) (*TriadStructureBase, error) {
	if len(*intervals) != 3 {
		return nil, errors.New("need 3 chord tones")
	}

	return &TriadStructureBase{
		name:      name,
		intervals: intervals,
	}, nil
}

//Name はコード名を表示します
func (tsb TriadStructureBase) Name() string {
	return tsb.name
}

//GetIntervals はコード構成音の間隔を取得します
func (tsb TriadStructureBase) GetIntervals() *[]tone.Interval {
	return tsb.intervals
}

//GetCode は指定されたルート音でのコード構成を返します
func (tsb TriadStructureBase) GetCode(root *tone.ScaleTone) *Code {
	newCode, err := NewCode(tsb.name, tsb.intervals, root)

	// 本来エラーは起こり得ないので、エラーになったらpanicを起こす
	if err != nil {
		panic(err)
	}

	return newCode
}

//GetCode は指定されたルート音でのトライアドコード構成を返します
func (tsb TriadStructureBase) GetTriadCode(root *tone.ScaleTone) *TriadCode {
	triadCode, err := NewTriadCode(tsb.name, tsb.intervals, root)

	// 本来エラーは起こり得ないので、エラーになったらpanicを起こす
	if err != nil {
		panic(err)
	}

	return triadCode
}