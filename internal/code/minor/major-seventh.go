package minor

import (
	"github.com/duosonic62/codanalyzer-domains/internal/code"
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
)

type MajorSeventh struct {
	name      string
	intervals *[]tone.Interval
}

//NewMajorSeventh MajorSeventhのコンストラクタ
func NewMajorSeventh() MajorSeventh {
	// 1度、m3度、5度、M7度のインターバル
	intervals := &[]tone.Interval{
		*tone.NewInterval(0),
		*tone.NewInterval(3),
		*tone.NewInterval(7),
		*tone.NewInterval(11),
	}
	return MajorSeventh{
		name:      "Minor Major Seventh",
		intervals: intervals,
	}
}

// Name はコード名を表示する
func (ms MajorSeventh) Name() string {
	return ms.name
}

// GetIntervalsはコード構成音の間隔を取得する
func (ms MajorSeventh) GetIntervals() *[]tone.Interval {
	return ms.intervals
}

// GetCode は指定されたルート音でのコード構成を返す
func (ms MajorSeventh) GetCode(root *tone.ScaleTone) *code.Code {
	code, err := code.NewCode(ms.name, ms.intervals, root)

	// 本来エラーは起こり得ないので、エラーになったらpanicを起こす
	if err != nil {
		panic(err)
	}

	return code
}