package code

import (
	"errors"
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
)

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

//ExtractTriadPatterns コードから3音構成を抽出する
func (c Code) ExtractTriadPatterns() (*[]TriadPattern, error) {
	//三音の組み合わせを取得する
	codeTones := *c.Tones
	triadCombinations := combination(codeTones, 3)

	//三音の組みあわあせからトライアドの構成音を分析する
	triadPatterns := make([]TriadPattern, 0)
	for _, triad := range triadCombinations {
		triadPattern, err := NewTriadPattern(&triad)
		if err != nil {
			return nil, err
		}
		triadPatterns = append(triadPatterns, *triadPattern...)
	}

	return &triadPatterns, nil
}

//combination トーンの組み合わせを取得
func combination(codeTones []tone.ScaleTone, r int) (result [][]tone.ScaleTone) {
	// 組み合わるの結果の個数が1の場合はそれぞれの要素を結果の配列に詰めて終わり
	if r == 1 {
		for _, codeTone := range codeTones {
			result = append(result, []tone.ScaleTone{codeTone})
		}
		return result
	}

	// 組み合わせ数と元の配列長が同じ場合はその配列を結果に詰めて終わり
	if r == len(codeTones) {
		result = append(result, codeTones)
		return
	}

	// 組み合わせ数より元の配列長の方が大きい場合は
	// 一個目の要素を固定して、配列の要素数を減らして組み合わせを再度行う
	reducedAnswersCh := make(chan [][]tone.ScaleTone)
	go func() {
		defer close(reducedAnswersCh)
		reducedAnswersCh <- combination(codeTones[1:], r-1)
	}()
	for _, reducedAnswer := range <-reducedAnswersCh {
		// 一つ目の要素を固定した結果をresultにつめる
		result = append(result, append([]tone.ScaleTone{codeTones[0]}, reducedAnswer...))
	}

	// 一つ目の要素を排除して再度組み合わせを行う
	pickedAnswerCh := make(chan [][]tone.ScaleTone)
	go func() {
		defer close(pickedAnswerCh)
		pickedAnswerCh <- combination(codeTones[1:], r)
	}()
	result = append(result, <-pickedAnswerCh...)

	return
}

