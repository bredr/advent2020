package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
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

func FindError(series []int, selection int, index int) *int {
	if index == len(series) {
		return nil
	}
	if index+1 <= selection {
		return FindError(series, selection, index+1)
	}
	target := series[index]
	for i := index - selection; i < index; i++ {
		for j := index - selection; j < index; j++ {
			if i != j {
				if series[i]+series[j] == target {
					return FindError(series, selection, index+1)
				}
			}
		}
	}
	return &target
}

func FindContiguousSet(series []int, target int, index int) []int {
	total := 0
	var set []int
	for i := index; i < len(series); i++ {
		total += series[i]
		set = append(set, series[i])
		if total == target {
			fmt.Printf("result2 set= %v total=%d\n", set, total)
			return set
		}
		if total > target {
			return FindContiguousSet(series, target, index+1)
		}
	}
	panic("no set found")
}

func main() {
	data, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}
	result1 := FindError(data, 25, 0)
	if result1 != nil {
		fmt.Printf("result1 = %d\n", *result1)
	}
	result2 := FindContiguousSet(data, *result1, 0)
	min, max := findMinAndMax(result2)
	fmt.Printf("result2 min=%d max=%d weakness = %d\n", min, max, min+max)

}

func findMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
