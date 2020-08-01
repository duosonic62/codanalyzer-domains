package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := []int{0, 1, 2}
	for i := range a {
		fmt.Println(strconv.Itoa(i))
		for j := 0; j < len(a); j++ {
			fmt.Println(strconv.Itoa(j) + ": " + strconv.Itoa((i+j)%3))
		}
	}
}
