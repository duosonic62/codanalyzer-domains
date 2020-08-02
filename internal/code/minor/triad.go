package minor

import (
	"github.com/duosonic62/codanalyzer-domains/internal/code"
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
)

// Triad はマイナートライアドコードを表す
type Triad struct {
	name      string
	intervals *[]tone.Interval
}

//NewTriad MinorTriadのコンストラクタ
func NewTriad() Triad {
	// 1度、3度、5度のインターバル
	intervals := &[]tone.Interval{
		*tone.NewInterval(0),
		*tone.NewInterval(3),
		*tone.NewInterval(7),
	}
	return Triad{
		name:      "Minor Triad",
		intervals: intervals,
	}
}

// Name はコード名を表示する
func (mt Triad) Name() string {
	return mt.name
}

// GetIntervalsはコード構成音の間隔を取得する
func (mt Triad) GetIntervals() *[]tone.Interval {
	return mt.intervals
}

// GetCode は指定されたルート音でのコード構成を返す
func (mt Triad) GetCode(root *tone.ScaleTone) *code.Code {
	code, err := code.NewCode(mt.name, mt.intervals, root)

	// 本来エラーは起こり得ないので、エラーになったらpanicを起こす
	if err != nil {
		panic(err)
	}

	return code
}

//GetTriadCode は指定されたルート音でのコード構成を返す
func (mt Triad) GetTriadCode(root *tone.ScaleTone) *code.TriadCode {
	triadCode, err := code.NewTriadCode(mt.name, mt.intervals, root)

	// 本来エラーは起こり得ないので、エラーになったらpanicを起こす
	if err != nil {
		panic(err)
	}

	return triadCode
}