package day05

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

func (o *operation) execute9001(stacks [][]byte) {
	stacks[o.dst] = append(make([]byte, o.n), stacks[o.dst]...)
	copy(stacks[o.dst], stacks[o.src][:o.n])
	stacks[o.src] = stacks[o.src][o.n:len(stacks[o.src])]
}

func Part2(r io.Reader) {
	var stacks [][]byte
	var ops []operation

	s := bufio.NewScanner(r)
	for s.Scan() {
		l := s.Bytes()
		if bytes.IndexByte(l, '[') == -1 {
			break
		}
		numStacks := (len(l) + 1) / 4
		if stacks == nil {
			stacks = make([][]byte, numStacks)
		}
		for n := 0; n < numStacks; n++ {
			box := l[1+n*4]
			if box != ' ' {
				stacks[n] = append(stacks[n], box)
			}
		}
	}

	for s.Scan() {
		var op operation
		_, err := fmt.Sscanf(s.Text(), "move %d from %d to %d", &op.n, &op.src, &op.dst)
		op.src--
		op.dst--
		if err != nil {
			if err == io.EOF {
				break
			}
			continue
		}
		ops = append(ops, op)
	}

	for _, o := range ops {
		o.execute9001(stacks[:])
	}

	for _, stack := range stacks {
		fmt.Printf("%c", stack[0])
	}
	fmt.Println()
}
