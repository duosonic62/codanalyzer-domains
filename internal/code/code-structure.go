package code

import "github.com/duosonic62/codanalyzer-domains/internal/tone"

type (

	//CodeStructure コード構成に関するインターフェス
	CodeStructure interface {
		// コード名
		Name() string

		// ルート音からのインターバルを取得する
		GetIntervals() *[]tone.Interval

		// 指定されたルート音からコードの構成音を取得する
		GetCode(root *tone.ScaleTone) *Code
	}

	//TriadCodeStructure はトライアドコードの構成を表すインターフェース
	TriadCodeStructure interface {
		// コード名
		Name() string

		// ルート音からのインターバルを取得する
		GetIntervals() *[]tone.Interval

		// 指定されたルート音からコードの構成音を取得する
		GetTriadCode(root *tone.ScaleTone) *TriadCode
	}
)
