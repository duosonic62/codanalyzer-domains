package main

import (
	"github.com/duosonic62/codanalyzer-domains/internal/code"
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
	"strconv"
)

func toStructureBase(collection *[]CodeInput) (*[]code.StructureBase, error) {
	codeStructureBases := make([]code.StructureBase, len(*collection))
	for index, codeInput := range *collection {
		intervals, err := toIntervals(&codeInput.Intervals)
		if err != nil {
			return nil, err
		}

		codeStructureBase, err := code.NewCodeStructureBase(codeInput.Name, intervals)
		if err != nil {
			return nil, err
		}

		codeStructureBases[index] = *codeStructureBase
	}

	return &codeStructureBases, nil
}

func toIntervals(rowIntervals *[]string) (*[]tone.Interval, error) {
	domains := make([]tone.Interval, len(*rowIntervals))

	for index, rowInterval := range *rowIntervals {
		intInterval, err := strconv.Atoi(rowInterval)
		if err != nil {
			return nil, err
		}

		domains[index] = *tone.NewInterval(intInterval)
	}

	return &domains, nil
}
