package day06

import (
	"bufio"
	"fmt"
	"io"
)

func Part2(r io.Reader) {
	s := bufio.NewScanner(r)

	for s.Scan() {
		line := s.Bytes()
		for i := 0; i < len(line); i++ {
			if i < 13 {
				continue
			}
			k := make(map[byte]struct{})
			for j := 0; j < 14; j++ {
				k[line[i-j]] = struct{}{}
			}
			if len(k) == 14 {
				fmt.Println(i + 1)
				break
			}
		}
	}
}
