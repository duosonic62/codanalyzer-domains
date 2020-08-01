package code

import (
	"errors"
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
)

//TriadPattern トライアドの解析パターン
type TriadPattern struct {
	Root      *tone.ScaleTone
	Intervals *[]tone.Interval
}

//NewTriadPattern は三音構成のトーンからTriadPatternのインスタンスを生成します
//三音構成のためルートが３つ生まれるので、三つのTriadPatternが生成されます
func NewTriadPattern(triads *[]tone.ScaleTone) (*[]TriadPattern, error) {
	triadTones := *triads
	if len(triadTones) != 3 {
		return nil, errors.New("triad tones needs 3 code tones")
	}

	patterns := make([]TriadPattern, len(triadTones))
	for i, triadTone := range triadTones {
		root := triadTone
		intervals := make([]tone.Interval, len(triadTones))

		//0->1->2, 1->2->0, 2->0->1の順番に並び替えてインターバルをとる
		for j := 0; j <len(triadTones) ; j++ {
			interval := root.CalculateInterval(triadTones[(i + j) % len(triadTones)])
			intervals[j] = *interval
		}

		pattern := TriadPattern{
			Root:      &root,
			Intervals: &intervals,
		}
		patterns[i] = pattern
	}

	return &patterns, nil
}

