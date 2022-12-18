package day02

import (
	"fmt"
	"io"
)

var moveScore2 = map[string]int{
	"AX": 0 + 3, "AY": 3 + 1, "AZ": 6 + 2,
	"BX": 0 + 1, "BY": 3 + 2, "BZ": 6 + 3,
	"CX": 0 + 2, "CY": 3 + 3, "CZ": 6 + 1,
}

func Part2(r io.Reader) {
	var score int
	for {
		var they, move string

		if _, err := fmt.Fscanln(r, &they, &move); err == io.EOF {
			break
		}

		score += moveScore2[they+move]
	}

	fmt.Println(score)
}
