package day06

import "os"

func Example_part1() {
	r, _ := os.Open("example.txt")

	Part1(r)

	// Output:
	// 7
	// 5
	// 6
	// 10
	// 11
}
