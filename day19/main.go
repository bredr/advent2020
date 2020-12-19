package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func rule0part1(rules map[int]string) *regexp.Regexp {
	re := regexp.MustCompile(`\d+`)
	for {
		done := true
		for k := range rules {
			rules[k] = re.ReplaceAllStringFunc(rules[k], func(s string) string {
				done = false
				idx, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				return rules[idx]
			})
		}
		if done {
			break
		}
	}
	regexRules := map[int]*regexp.Regexp{}

	for k := range rules {
		replacer := strings.NewReplacer("\"", "", " ", "")
		regexRules[k] = regexp.MustCompile("^" + replacer.Replace(rules[k]) + "$")
	}

	return regexRules[0]
}

func rule0part2(rules map[int]string) *regexp.Regexp {
	rules[8] = "(42 | 42 8)"
	rules[11] = "(42 31 | 42 11 31)"

	count8 := 0
	count11 := 0
	maxRecursion := 10
	re := regexp.MustCompile(`\d+`)
	for {
		done := true
		for k := range rules {
			rules[k] = re.ReplaceAllStringFunc(rules[k], func(s string) string {
				done = false
				idx, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				if idx == 8 {
					count8++
					if count8 > maxRecursion {
						return "(42)"
					}
				}
				if idx == 11 {
					count11++
					if count11 > maxRecursion {
						return "(42 31)"
					}
				}
				return rules[idx]
			})
		}
		if done {
			break
		}
	}
	regexRules := map[int]*regexp.Regexp{}

	for k := range rules {
		replacer := strings.NewReplacer("\"", "", " ", "")
		regexRules[k] = regexp.MustCompile("^" + replacer.Replace(rules[k]) + "$")
	}

	return regexRules[0]
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(string(contents), "\n\n")
	rules := make(map[int]string)
	for _, line := range strings.Split(parts[0], "\n") {
		lPart := strings.Split(line, ": ")
		idx, err := strconv.Atoi(lPart[0])
		if err != nil {
			panic(err)
		}

		char := "(" + strings.TrimSpace(lPart[1]) + ")"
		rules[idx] = char
	}
	rulePart1 := rule0part1(rules)

	rules = make(map[int]string)
	for _, line := range strings.Split(parts[0], "\n") {
		lPart := strings.Split(line, ": ")
		idx, err := strconv.Atoi(lPart[0])
		if err != nil {
			panic(err)
		}

		char := "(" + strings.TrimSpace(lPart[1]) + ")"
		rules[idx] = char
	}
	rulePart2 := rule0part2(rules)
	countPart1 := 0
	countPart2 := 0

	for _, line := range strings.Split(parts[1], "\n") {
		if rulePart1.MatchString(line) {
			countPart1++
		}
		if rulePart2.MatchString(line) {
			countPart2++
		}
	}
	fmt.Printf("result1 = %d\n", countPart1)
	fmt.Printf("result2 = %d\n", countPart2)
}
