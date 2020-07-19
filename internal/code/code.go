package code

import "github.com/duosonic62/codanalyzer-domains/internal/tone"

// コード構成に関するインターフェス
// コードの分析の窓口
type CodeStructure interface {
	// コード名
	Name() string

	// ルート音からのインターバルを取得する
	GetIntervals() *[]tone.Interval

	// 指定されたルート音からコードの構成音を取得する
	GetCode(root tone.ScaleTone) *Code

	// 指定した音階と
	Equals()
}

// コードのストラクチャ
type Code struct {
	Name  string
	Root  *tone.ScaleTone
	Tones *[]tone.ScaleTone
}
