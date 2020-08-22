package main

import (
	"fmt"
)

func main() {
	codes, err := readFromJson("codes.json")
	if err != nil {
		fmt.Println(err)
	}

	codeStructureBases, err := toStructureBase(codes)
	if err != nil {
		fmt.Println(err)
	}

	for _, code := range *codeStructureBases {
		fmt.Println(code)
	}
}
