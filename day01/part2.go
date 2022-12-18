package day01

import (
	"fmt"
	"io"
	"sort"
)

func Part2(r io.Reader) {
	var totals []int
	var total int

	for {
		var cals int
		n, err := fmt.Fscanln(r, &cals)
		total += cals
		if n == 0 {
			totals = append(totals, total)
			total = 0
		}
		if err == io.EOF {
			break
		}
	}
	sort.Ints(totals)

	var sum int
	for _, total := range totals[len(totals)-3:] {
		sum += total
	}

	fmt.Println(sum)
}
