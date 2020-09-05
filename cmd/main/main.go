package main

import (
	"fmt"
	"github.com/duosonic62/codanalyzer-domains/cmd/factory"
)

func main() {
	codes, err := factory.ReadCodesFromJson("codes.json")
	if err != nil {
		fmt.Println(err)
	}

	triads, err := factory.ReadCodesFromJson("triad.json")
	if err != nil {
		fmt.Println(err)
	}

	codeStructureBases, err := factory.MakeCodeStructureBase(codes)
	if err != nil {
		fmt.Println(err)
	}
	for _, code := range *codeStructureBases {
		fmt.Println(code.Name())
		fmt.Println(code.GetIntervals())
	}

	triadStructureBases, err := factory.MakeTriadStructureBase(triads)
	if err != nil {
		fmt.Println(err)
	}
	for _, code := range *triadStructureBases {
		fmt.Println(code.Name())
		fmt.Println(code.GetIntervals())
	}
}
