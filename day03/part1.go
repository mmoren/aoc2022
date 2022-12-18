package day03

import (
	"bufio"
	"fmt"
	"io"
)

func Part1(r io.Reader) {
	var sum int

	s := bufio.NewScanner(r)
	for s.Scan() {
		all := s.Text()

		half := len(all) / 2
		a, b := all[:half], all[half:]

	search:
		for i := range a {
			for j := range b {
				if a[i] == b[j] {
					sum += priority(a[i])
					break search
				}
			}
		}
	}

	fmt.Println(sum)
}

func priority(item byte) int {
	if item > 'Z' {
		return 1 + int(item-'a')
	}
	return 27 + int(item-'A')
}
