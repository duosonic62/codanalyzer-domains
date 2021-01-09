package main

import (
	"fmt"
	"github.com/duosonic62/codanalyzer-domains/cmd/factory"
	"github.com/duosonic62/codanalyzer-domains/internal/code"
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
)

func main() {
	codeStructureBases, err := factory.LoadCodes("codes.json")
	if err != nil {
		fmt.Println(err)
	}
	for _, code := range *codeStructureBases {
		fmt.Println(code.Name())
		fmt.Println(code.GetIntervals())
	}

	triadStructureBases, err := factory.LoadTriads("triad.json")
	if err != nil {
		fmt.Println(err)
	}
	for _, code := range *triadStructureBases {
		fmt.Println(code.Name())
		fmt.Println(code.GetIntervals())
	}

	tones := [] tone.Interval{
		tone.Intervals.R,
		tone.Intervals.Major3,
		tone.Intervals.Major7,
		tone.Intervals.Natural9,
		tone.Intervals.Sharp11,
		tone.Intervals.Natural13,
	}
	seventhAddUpper, _ := code.NewCodeStructureBase("9", &tones)
	fmt.Println(seventhAddUpper.Name())
	fmt.Println(seventhAddUpper.GetIntervals())
}
