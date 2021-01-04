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

var Intervals = struct {
	R         Interval
	Minor2    Interval
	Major2    Interval
	Minor3    Interval
	Major3    Interval
	Perfect4  Interval
	Sharp4    Interval
	Flat5     Interval
	Perfect5  Interval
	Sharp5    Interval
	Minor6    Interval
	Major6    Interval
	Sharp6    Interval
	Minor7    Interval
	Major7    Interval
	Flat9     Interval
	Natural9  Interval
	Natural11 Interval
	Sharp11   Interval
	Flat13    Interval
	Natural13 Interval
}{
	R: Interval{
		value: intervalMap["R"],
		name:  "R",
	},
	Minor2: Interval{
		value: intervalMap["m2"],
		name:  "m2",
	},
	Major2: Interval{
		value: intervalMap["M2"],
		name:  "M2",
	},
	Minor3: Interval{
		value: intervalMap["m3"],
		name:  "m3",
	},
	Major3: Interval{
		value: intervalMap["M3"],
		name:  "M3",
	},
	Perfect4: Interval{
		value: intervalMap["P4"],
		name:  "P4",
	},
	Sharp4: Interval{
		value: intervalMap["#4"],
		name:  "#4",
	},
	Flat5: Interval{
		value: intervalMap["b5"],
		name:  "b5",
	},
	Perfect5: Interval{
		value: intervalMap["P5"],
		name:  "P5",
	},
	Sharp5: Interval{
		value: intervalMap["#5"],
		name:  "#5",
	},
	Minor6: Interval{
		value: intervalMap["m6"],
		name:  "m6",
	},
	Major6: Interval{
		value: intervalMap["M6"],
		name:  "M6",
	},
	Sharp6: Interval{
		value: intervalMap["#6"],
		name:  "#6",
	},
	Minor7: Interval{
		value: intervalMap["m7"],
		name:  "m7",
	},
	Major7: Interval{
		value: intervalMap["M7"],
		name:  "M7",
	},
	Flat9: Interval{
		value: intervalMap["b9"],
		name:  "b9",
	},
	Natural9: Interval{
		value: intervalMap["9"],
		name:  "9",
	},
	Natural11: Interval{
		value: intervalMap["11"],
		name:  "11",
	},
	Sharp11: Interval{
		value: intervalMap["#11"],
		name:  "#11",
	},
	Flat13: Interval{
		value: intervalMap["b13"],
		name:  "b13",
	},
	Natural13: Interval{
		value: intervalMap["13"],
		name:  "13",
	},
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
