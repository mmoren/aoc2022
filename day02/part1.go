package day02

import (
	"fmt"
	"io"
)

var moveScore1 = map[string]int{
	"AX": 3 + 1, "AY": 6 + 2, "AZ": 0 + 3,
	"BX": 0 + 1, "BY": 3 + 2, "BZ": 6 + 3,
	"CX": 6 + 1, "CY": 0 + 2, "CZ": 3 + 3,
}

func Part1(r io.Reader) {
	var score int
	for {
		var they, you string

		if _, err := fmt.Fscanln(r, &they, &you); err == io.EOF {
			break
		}

		score += moveScore1[they+you]
	}

	fmt.Println(score)
}
