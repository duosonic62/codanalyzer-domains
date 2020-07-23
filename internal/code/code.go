package code

import (
	"errors"
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
)

// コード構成に関するインターフェス
// コードの分析の窓口
type CodeStructure interface {
	// コード名
	Name() string

	// ルート音からのインターバルを取得する
	GetIntervals() *[]tone.Interval

	// 指定されたルート音からコードの構成音を取得する
	GetCode(root *tone.ScaleTone) *Code

	// 指定した音階のコードトーンを探す
	GetEquivalenceCodes() *[]Code
}

// トライアドコードの構造体
type TriadCode struct {
	Name  string
	Root  *tone.ScaleTone
	Tones *[]tone.ScaleTone
}

// トライアドコードを作成
func NewTriadCode(name string, root *tone.ScaleTone, tones *[]tone.ScaleTone) (*TriadCode, error) {
	if root == nil || tones == nil {
		return nil, errors.New("root and tones must not be null")
	}

	if len(*tones) != 3 {
		return nil, errors.New("triad code constituting tone must be 3 tones")
	}

	return &TriadCode{
		Name:  name,
		Root:  root,
		Tones: tones,
	}, nil
}

// コードに変換
func (t TriadCode) Code() *Code {
	return &Code{
		Name:  t.Name,
		Root:  t.Root,
		Tones: t.Tones,
	}
}

// コードの構造体
type Code struct {
	Name  string
	Root  *tone.ScaleTone
	Tones *[]tone.ScaleTone
}

// NewCode はインターバルとルート音からコードの構成を生成する
func NewCode(name string, intervals *[]tone.Interval, root *tone.ScaleTone) *Code {
	// 構成音をインターバル分用意する
	tones := make([]tone.ScaleTone, len(*intervals))

	// 高静音を取得する
	for i, interval := range *intervals {
		tones[i] = *root.CalculateTone(&interval)
	}
	
	return &Code{
		Name:  name,
		Root:  root,
		Tones: &tones,
	}
}
