package day05

import "os"

func Example_part2() {
	r, _ := os.Open("example.txt")

	Part2(r)

	// Output: MCD
}
