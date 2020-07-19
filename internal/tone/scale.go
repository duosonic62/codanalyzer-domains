package tone

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

// 二音間の間隔を計算する
func (baseTone ScaleTone) CalculateInterval(compareTone ScaleTone) *Interval {
	if compareTone.interval > baseTone.interval {
		return NewInterval(compareTone.interval - baseTone.interval)
	}
	return NewInterval(compareTone.interval + (scaleToneCount - baseTone.interval))
}