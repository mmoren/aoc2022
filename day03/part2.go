package day03

import (
	"bufio"
	"fmt"
	"io"
)

func Part2(r io.Reader) {
	var sum int

	s := bufio.NewScanner(r)
outer:
	for {
		numBags := make(map[byte]int)

		for i := 0; i < 3; i++ {
			inBag := make(map[byte]struct{})
			if !s.Scan() {
				break outer
			}
			for _, c := range s.Bytes() {
				if _, ok := inBag[c]; !ok {
					numBags[c]++
					inBag[c] = struct{}{}
				}
			}
		}

		for c, freq := range numBags {
			if freq == 3 {
				sum += priority(c)
				break
			}
		}
	}

	fmt.Println(sum)
}
