package day06

import "os"

func Example_part2() {
	r, _ := os.Open("example.txt")

	Part2(r)

	// Output:
	// 19
	// 23
	// 23
	// 29
	// 26
}
