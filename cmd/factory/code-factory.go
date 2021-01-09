package factory

import (
	"github.com/duosonic62/codanalyzer-domains/internal/code"
	"github.com/duosonic62/codanalyzer-domains/internal/tone"
)

//makeCodeStructureBase はコード入力用のDTOからコード構造クラスを生成します
func makeCodeStructureBase(collection *[]CodeInput) (*[]code.StructureBase, error) {
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

//makeTriadStructureBase  はコード入力用のDTOからトライアド構造クラスを生成します
func makeTriadStructureBase(collection *[]CodeInput) (*[]code.TriadStructureBase, error) {
	triadStructureBases := make([]code.TriadStructureBase, len(*collection))
	for index, codeInput := range *collection {
		intervals, err := toIntervals(&codeInput.Intervals)
		if err != nil {
			return nil, err
		}

		triadStructureBase, err := code.NewTriadStructureBase(codeInput.Name, intervals)
		if err != nil {
			return nil, err
		}

		triadStructureBases[index] = *triadStructureBase
	}

	return &triadStructureBases, nil
}

func toIntervals(rowIntervals *[]string) (*[]tone.Interval, error) {
	domains := make([]tone.Interval, len(*rowIntervals))

	for index, rowInterval := range *rowIntervals {
		interval, err := tone.NewIntervalFromName(rowInterval)
		if err != nil {
			return nil, err
		}

		domains[index] = *interval
	}

	return &domains, nil
}
