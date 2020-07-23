package code

import (
	"errors"
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
)

// トライアドコードの構造体
type TriadCode struct {
	Name  *CodeName
	Root  *tone.ScaleTone
	Tones *[]tone.ScaleTone
}

// トライアドコードを作成
func NewTriadCode(structureName string, intervals *[]tone.Interval, root *tone.ScaleTone) (*TriadCode, error) {
	if root == nil || intervals == nil {
		return nil, errors.New("root and tones must not be null")
	}

	if len(*intervals) != 3 {
		return nil, errors.New("triad code constituting tone must be 3 tones")
	}

	// 構成音をインターバル分用意する
	tones := make([]tone.ScaleTone, len(*intervals))

	// 構成音を取得する
	for i, interval := range *intervals {
		tones[i] = *root.GetToneWithApartInterval(&interval)
	}

	return &TriadCode{
		Name:  NewCodeName(root, structureName),
		Root:  root,
		Tones: &tones,
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
	Name  *CodeName
	Root  *tone.ScaleTone
	Tones *[]tone.ScaleTone
}

// NewCode はインターバルとルート音からコードの構成を生成する
func NewCode(structureName string, intervals *[]tone.Interval, root *tone.ScaleTone) (*Code, error) {
	if root == nil || intervals == nil {
		return nil, errors.New("root and tones must not be null")
	}

	if len(*intervals) < 3 {
		return nil, errors.New("need at least 3 chord tones")
	}

	// 構成音をインターバル分用意する
	tones := make([]tone.ScaleTone, len(*intervals))

	// 構成音を取得する
	for i, interval := range *intervals {
		tones[i] = *root.GetToneWithApartInterval(&interval)
	}

	return &Code{
		Name:  NewCodeName(root, structureName),
		Root:  root,
		Tones: &tones,
	}, nil
}
