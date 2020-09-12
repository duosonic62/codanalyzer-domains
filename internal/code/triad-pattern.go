package code

import (
	"errors"
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
)

//TriadInversion トライアドの解析パターン
type TriadInversion struct {
	Root      *tone.ScaleTone
	Intervals *[]tone.Interval
}

//NewTriadInversions は三音構成のトーンから３つの転回形のコードパターン、TriadInversionのインスタンスを生成します
//三音構成のためルートが３つ生まれるので、三つのTriadInversionが生成されます
func NewTriadInversions(triads *[]tone.ScaleTone) (*[]TriadInversion, error) {
	triadTones := *triads
	if len(triadTones) != 3 {
		return nil, errors.New("triad tones needs 3 code tones")
	}

	patterns := make([]TriadInversion, len(triadTones))
	for i, triadTone := range triadTones {
		root := triadTone
		intervals := make([]tone.Interval, len(triadTones))

		//0->1->2, 1->2->0, 2->0->1の順番に並び替えてインターバルをとる
		for j := 0; j <len(triadTones) ; j++ {
			interval, err := root.CalculateInterval(triadTones[(i + j) % len(triadTones)])
			if err != nil {
				return nil, err
			}
			intervals[j] = *interval
		}

		pattern := TriadInversion{
			Root:      &root,
			Intervals: &intervals,
		}
		patterns[i] = pattern
	}

	return &patterns, nil
}

