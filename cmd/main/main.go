package main

import "fmt"

func main() {
	codes, err := readFromJson("codes.json")
	if err != nil {
		fmt.Println(err)
	}

	for _, code := range *codes {
		fmt.Println(code)
	}
}
