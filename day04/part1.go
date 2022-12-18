package day04

import (
	"fmt"
	"io"
)

type interval [2]int

func (a *interval) contains(b interval) bool {
	return a[0] <= b[0] && a[1] >= b[1]
}

func Part1(r io.Reader) {
	var num int

	for {
		var g1, g2 interval
		_, err := fmt.Fscanf(r, "%d-%d,%d-%d\n", &g1[0], &g1[1], &g2[0], &g2[1])
		if err == io.EOF {
			break
		}
		if g1.contains(g2) || g2.contains(g1) {
			num++
		}
	}

	fmt.Println(num)
}
