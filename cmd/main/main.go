package main

import (
	"fmt"
)

func main() {
	codes, err := readCodesFromJson("codes.json")
	if err != nil {
		fmt.Println(err)
	}

	triads, err := readCodesFromJson("triad.json")
	if err != nil {
		fmt.Println(err)
	}

	codeStructureBases, err := toStructureBase(codes)
	if err != nil {
		fmt.Println(err)
	}
	for _, code := range *codeStructureBases {
		fmt.Println(code.Name())
		fmt.Println(code.GetIntervals())
	}

	triadStructureBases, err := toTriadStructureBase(triads)
	if err != nil {
		fmt.Println(err)
	}
	for _, code := range *triadStructureBases {
		fmt.Println(code.Name())
		fmt.Println(code.GetIntervals())
	}
}
