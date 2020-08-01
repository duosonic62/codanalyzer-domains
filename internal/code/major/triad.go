package major

import (
	"github.com/duosonic62/codanalyzer-domains/internal/code"
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
)

// MajorTriad はメジャートライアドコードを表す
type MajorTriad struct {
	name      string
	intervals *[]tone.Interval
}

// NewMajorTriad MajorTriadのコンストラクタ
func NewMajorTriad() MajorTriad {
	// 1度、3度、5度のインターバル
	intervals := &[]tone.Interval{
		*tone.NewInterval(0),
		*tone.NewInterval(4),
		*tone.NewInterval(7),
	}
	return MajorTriad{
		name:      "Major Triad",
		intervals: intervals,
	}
}

// Name はコード名を表示する
func (mt MajorTriad) Name() string {
	return mt.name
}

// GetIntervalsはコード構成音の間隔を取得する
func (mt MajorTriad) GetIntervals() *[]tone.Interval {
	return mt.intervals
}

// GetCode は指定されたルート音でのコード構成を返す
func (mt MajorTriad) GetCode(root *tone.ScaleTone) *code.Code {
	newCode, err := code.NewCode(mt.name, mt.intervals, root)

	// 本来エラーは起こり得ないので、エラーになったらpanicを起こす
	if err != nil {
		panic(err)
	}

	return newCode
}

func (mt MajorTriad) GetTriadCode(root *tone.ScaleTone) *code.TriadCode {
	triadCode, err := code.NewTriadCode(mt.name, mt.intervals, root)

	// 本来エラーは起こり得ないので、エラーになったらpanicを起こす
	if err != nil {
		panic(err)
	}

	return triadCode
}