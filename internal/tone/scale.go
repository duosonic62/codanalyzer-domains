package tone

import "errors"

type ScaleTone struct {
	value    string
	interval int
}

// クロマチックスケールの音数
var scaleToneCount = 12

// 音階内に含まれるトーンの定義
var ScaleTones = struct {
	C  ScaleTone
	CS ScaleTone
	D  ScaleTone
	DS ScaleTone
	E  ScaleTone
	F  ScaleTone
	FS ScaleTone
	G  ScaleTone
	GS ScaleTone
	A  ScaleTone
	AS ScaleTone
	B  ScaleTone
}{
	C: ScaleTone{
		value:    "C",
		interval: 0,
	},
	CS: ScaleTone{
		value:    "C#",
		interval: 1,
	},
	D: ScaleTone{
		value:    "D",
		interval: 2,
	},
	DS: ScaleTone{
		value:    "D#",
		interval: 3,
	},
	E: ScaleTone{
		value:    "E",
		interval: 4,
	},
	F: ScaleTone{
		value:    "F",
		interval: 5,
	},
	FS: ScaleTone{
		value:    "F#",
		interval: 6,
	},
	G: ScaleTone{
		value:    "G",
		interval: 7,
	},
	GS: ScaleTone{
		value:    "G#",
		interval: 8,
	},
	A: ScaleTone{
		value:    "A",
		interval: 9,
	},
	AS: ScaleTone{
		value:    "A#",
		interval: 10,
	},
	B: ScaleTone{
		value:    "B",
		interval: 11,
	},
}

//CalculateInterval は二音間のインターバルを計算する
func (sc ScaleTone) CalculateInterval(compareTone ScaleTone) *Interval {
	if compareTone.interval > sc.interval {
		return NewInterval(compareTone.interval - sc.interval)
	}
	return NewInterval(compareTone.interval + (scaleToneCount - sc.interval))
}

//GetToneWithApartInterval は基音からインターバル分離れた音を取得する
func (sc ScaleTone) GetToneWithApartInterval(interval *Interval) *ScaleTone {
	toneInterval := sc.interval + interval.value
	if toneInterval >= scaleToneCount {
		toneInterval = toneInterval - scaleToneCount
	}

	// エラーは本来起こり得ないので発生したらpanicを起こさせる
	scaleTone, err := scaleToneFromInterval(toneInterval)
	if err != nil {
		panic(err)
	}

	return scaleTone
}

func (sc ScaleTone) String() string {
	return sc.value
}

//ScaleToneFromInterval はCからのインターバルからスケール音を取得する
func scaleToneFromInterval(interval int) (*ScaleTone, error) {
	switch interval {
	case 0:
		return &ScaleTones.C, nil
	case 1:
		return &ScaleTones.CS, nil
	case 2:
		return &ScaleTones.D, nil
	case 3:
		return &ScaleTones.DS, nil
	case 4:
		return &ScaleTones.E, nil
	case 5:
		return &ScaleTones.F, nil
	case 6:
		return &ScaleTones.FS, nil
	case 7:
		return &ScaleTones.G, nil
	case 8:
		return &ScaleTones.GS, nil
	case 9:
		return &ScaleTones.A, nil
	case 10:
		return &ScaleTones.AS, nil
	case 11:
		return &ScaleTones.B, nil
	default:
		return nil, errors.New("scale tone interval must be 0 - 11")
	}
}