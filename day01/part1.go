package day01

import (
	"fmt"
	"io"
)

func Part1(r io.Reader) {
	var total, max int

	for {
		var cals int
		n, err := fmt.Fscanln(r, &cals)
		total += cals
		if total > max {
			max = total
		}
		if err == io.EOF {
			break
		}
		if n == 0 {
			total = 0
		}
	}

	fmt.Println(max)
}
