package tone

import "strconv"

// 二音間の間隔
type Interval struct {
	value int
}

// Intervalインスタンスを生成
func NewInterval(interval int) *Interval {
	return &Interval{value: interval}
}

// インターバルを表示するメソッド
// ユーザが理解しやすいように全音を１に直す
func (i Interval) String() string {
	wholeTone := i.value / 2
	halfTone := i.value % 2 == 1

	// 単に割る2をしても良いが、不動小数点で誤差が生まれそうなので
	// 2で割り切れなかったら0.5の表記を足す
	if halfTone {
		return strconv.Itoa(wholeTone) + ".5"
	}

	return strconv.Itoa(wholeTone)
}
