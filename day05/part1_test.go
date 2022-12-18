package day05

import "os"

func Example_part1() {
	r, _ := os.Open("example.txt")

	Part1(r)

	// Output: CMZ
}
