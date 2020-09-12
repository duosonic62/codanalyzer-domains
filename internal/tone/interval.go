package tone

import (
	"errors"
	"strconv"
)

// 二音間の間隔
type Interval struct {
	value int
	name  string
}

var intervalMap = map[string]int{
	"R":  0,
	"m2": 1,
	"M2": 2,
	"m3": 3,
	"M3": 4,
	"P4": 5,
	"#4": 6,
	"b5": 6,
	"P5": 7,
	"#5": 8,
	"m6": 8,
	"M6": 9,
	"#6": 10,
	"m7": 10,
	"M7": 11,
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

// String はインターバルを表示するメソッド
// ユーザが理解しやすいように全音を１に直す
func (i Interval) String() string {
	wholeTone := i.value / 2
	halfTone := i.value%2 == 1

	// 単に割る2をしても良いが、不動小数点で誤差が生まれそうなので
	// 2で割り切れなかったら0.5の表記を足す
	if halfTone {
		return strconv.Itoa(wholeTone) + ".5"
	}

	return strconv.Itoa(wholeTone)
}
