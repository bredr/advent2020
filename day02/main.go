package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type password struct {
	min  int
	max  int
	char rune
	pwd  []rune
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]password, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []password
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		limits := strings.Split(parts[0], "-")

		min, err := strconv.Atoi(limits[0])
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(limits[1])
		if err != nil {
			panic(err)
		}
		lines = append(lines, password{min: min, max: max, char: []rune(parts[1])[0], pwd: []rune(parts[2])})
	}
	return lines, scanner.Err()
}

func main() {
	data, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}
	correctPwds := 0
	for _, x := range data {
		count := 0
		for _, c := range x.pwd {
			if c == x.char {
				count++
			}
		}
		if count >= x.min && count <= x.max {
			correctPwds++
		}
	}
	fmt.Printf("Old job correct passwords = %d\n", correctPwds)

	newCorrectPwds := 0
	for _, x := range data {
		first := x.pwd[x.min-1]
		second := x.pwd[x.max-1]
		if (first == x.char && second != x.char) || (first != x.char && second == x.char) {
			newCorrectPwds++
		}
	}
	fmt.Printf("New job correct passwords = %d\n", newCorrectPwds)
}
