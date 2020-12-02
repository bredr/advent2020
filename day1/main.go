package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if !(err != nil) {
			lines = append(lines, i)
		}
	}
	return lines, scanner.Err()
}

func main() {
	data, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}
	for _, x := range data {
		for _, y := range data {
			if x+y == 2020 {
				fmt.Printf("%d * %d = %d\n", x, y, x*y)
				return
			}
		}
	}
}
