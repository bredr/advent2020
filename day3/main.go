package main

import (
	"bufio"
	"fmt"
	"os"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([][]rune, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, []rune(scanner.Text()))
	}
	return lines, scanner.Err()
}

func hits(data [][]rune, right, down int) int {
	trees := 0
	for i, r := range data {
		if i%down == 0 {
			x := ((i / down) * right) % len(r)
			if r[x] == '#' {
				trees++
			}
		}
	}
	fmt.Printf("Right %d, down %d -> Trees = %d\n", right, down, trees)

	return trees
}

func main() {

	data, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}
	o11 := hits(data, 1, 1)
	o31 := hits(data, 3, 1)
	o51 := hits(data, 5, 1)
	o71 := hits(data, 7, 1)
	o12 := hits(data, 1, 2)
	fmt.Printf("result = %d\n", o11*o31*o51*o71*o12)
}
