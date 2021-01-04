package tone

import (
	"errors"
)

// 二音間の間隔
type Interval struct {
	value int
	name  string
}

var intervalMap = map[string]int{
	"R":   0,
	"m2":  1,
	"M2":  2,
	"m3":  3,
	"M3":  4,
	"P4":  5,
	"#4":  6,
	"b5":  6,
	"P5":  7,
	"#5":  8,
	"m6":  8,
	"M6":  9,
	"#6":  10,
	"m7":  10,
	"M7":  11,
	"b9":  1,
	"9":   2,
	"11":  5,
	"#11": 6,
	"b13": 8,
	"13":  9,
}

// NewInterval はIntervalインスタンスを生成
func NewInterval(interval int) (*Interval, error) {
	var intervalName string
	for name, value := range intervalMap {
		if interval == value {
			intervalName = name
			break
		}
	}

	if intervalName == "" {
		return nil, errors.New(intervalName + " is not found")
	}

	return &Interval{
		value: interval,
		name:  intervalName,
	}, nil
}

//NewIntervalFromName はインターバルを取得しました
func NewIntervalFromName(intervalName string) (*Interval, error) {
	interval, ok := intervalMap[intervalName]
	if ok {
		return &Interval{
			value: interval,
			name:  intervalName,
		}, nil
	}

	return nil, errors.New(intervalName + " is not found")
}

// String はインターバルを表示する
func (i Interval) String() string {
	return i.name
}

//IsEquivalent はインターバルどうしが等価であるか比較します
//nameが異なっていてもインターバルは等価であることがあるため比較はvalueで行います
func (i Interval) IsEquivalent(target *Interval) bool {
	return i.value == target.value
}
