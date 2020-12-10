package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func readLines(path string) ([]int, error) {

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	numbers := make([]int, len(lines))
	for i, line := range lines {
		numbers[i], err = strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
	}
	return numbers, nil
}

func NextAdapter(data []int, acc map[int]int, previousJolt int, index int) map[int]int {
	if index == len(data) {
		acc[3]++
		return acc
	}
	jump := data[index] - previousJolt
	if jump > 3 {
		acc[3]++
		return acc
	}
	acc[jump]++
	return NextAdapter(data, acc, data[index], index+1)
}

func Arrangements(data []int, acc int, currentRun int, previousJolt int, index int) int {
	if index == len(data) {
		if currentRun != 0 {
			return acc * Trib(currentRun)
		}
		return acc
	}
	jump := data[index] - previousJolt
	if jump > 3 {
		if currentRun != 0 {
			return acc * Trib(currentRun)
		}
		return acc
	}
	if jump != 1 {
		if currentRun != 0 {
			acc *= Trib(currentRun)
			currentRun = 0
		}
	} else {
		currentRun++
	}
	return Arrangements(data, acc, currentRun, data[index], index+1)
}

func Trib(i int) int {
	switch i {
	case 0:
		return 1
	case 1:
		return 1
	case 2:
		return 2
	default:
		return Trib(i-1) + Trib(i-2) + Trib(i-3)
	}
}

func main() {
	data, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}
	sort.Ints(data)
	result1 := NextAdapter(data, make(map[int]int), 0, 0)
	fmt.Printf("result1 = %d\n", result1[1]*result1[3])

	total := Arrangements(data, 1, 0, 0, 0)
	fmt.Printf("result2 = %d\n", total)
}
