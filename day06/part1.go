package day06

import (
	"bufio"
	"fmt"
	"io"
)

func Part1(r io.Reader) {
	s := bufio.NewScanner(r)

	for s.Scan() {
		line := s.Bytes()
		for i := 0; i < len(line); i++ {
			if i < 3 {
				continue
			}
			k := map[byte]struct{}{line[i]: {}, line[i-1]: {}, line[i-2]: {}, line[i-3]: {}}
			if len(k) == 4 {
				fmt.Println(i + 1)
				break
			}
		}
	}
}
