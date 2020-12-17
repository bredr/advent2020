package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type limits struct {
	min int
	max int
}

func Limit(s string) limits {
	min, err := strconv.Atoi(strings.Split(s, "-")[0])
	if err != nil {
		panic(err)
	}
	max, err := strconv.Atoi(strings.Split(s, "-")[1])
	if err != nil {
		panic(err)
	}
	return limits{min: min, max: max}
}

type ticket []int

type input struct {
	rules         map[string][]limits
	yourTicket    ticket
	nearbyTickets []ticket
}

func notValidForAnyField(rules map[string][]limits, value int) bool {
	accepted := true
	for _, rule := range rules {
		for _, l := range rule {
			if value >= l.min && value <= l.max {
				accepted = false
			}
		}
	}
	return accepted
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")

	scannedInput := &input{rules: make(map[string][]limits), yourTicket: []int{}, nearbyTickets: []ticket{}}
	inRules := true
	inYourTicket := false
	inNearbyTickets := false
	for _, line := range lines {
		if line == "" {
			if inRules {
				inRules = false
				inYourTicket = true
			} else if inYourTicket {
				inYourTicket = false
				inNearbyTickets = true
			}
			continue
		}
		if inRules {
			field := strings.Split(line, ":")[0]
			limitsRaw := strings.Split(strings.Split(line, ": ")[1], " or ")
			limits := make([]limits, len(limitsRaw))
			for i, r := range limitsRaw {
				limits[i] = Limit(r)
			}
			scannedInput.rules[field] = limits
			continue
		}
		if !strings.Contains(line, "ticket") {
			raw := strings.Split(line, ",")
			v := make([]int, len(raw))
			for i, r := range raw {
				value, err := strconv.Atoi(r)
				if err != nil {
					panic(err)
				}
				v[i] = value
			}
			if inYourTicket {
				scannedInput.yourTicket = v
			} else if inNearbyTickets {
				scannedInput.nearbyTickets = append(scannedInput.nearbyTickets, v)
			}
		}
	}

	sum := 0
	var validTickets []ticket
	for _, t := range scannedInput.nearbyTickets {
		valid := true
		for _, v := range t {
			if notValidForAnyField(scannedInput.rules, v) {
				valid = false
				sum += v
			}
		}
		if valid {
			validTickets = append(validTickets, t)
		}
	}
	fmt.Printf("result1 = %d\n", sum)

	possibleColumns := make(map[string][]int)
	for field, rules := range scannedInput.rules {
		possibleColumns[field] = []int{}
		for column := range validTickets[0] {
			possible := true
			for _, t := range validTickets {
				v := t[column]
				valid := false
				for _, r := range rules {
					if v >= r.min && v <= r.max {
						valid = true
						break
					}
				}
				if !valid {
					possible = false
					break
				}
			}
			if possible {
				possibleColumns[field] = append(possibleColumns[field], column)
			}
		}
	}

	columnsMap := make(map[string]int)
	for {
		done := true
		for _, v := range possibleColumns {
			if len(v) != 1 {
				done = false
			}
		}
		if done {
			break
		}
		var selected int
		for k, v := range possibleColumns {
			if len(v) == 1 {
				columnsMap[k] = v[0]
				selected = v[0]
			}
		}
		for k, v := range possibleColumns {
			if len(v) > 1 {
				possibleColumns[k] = remove(v, selected)
			}
		}
	}

	result2 := 1
	for k, v := range columnsMap {
		if strings.Contains(k, "departure") {
			result2 *= scannedInput.yourTicket[v]
		}
	}
	fmt.Println(columnsMap)
	fmt.Printf("result2 = %d\n", result2)
}

func remove(xx []int, y int) (z []int) {
	for _, x := range xx {
		if x != y {
			z = append(z, x)
		}
	}
	return z
}
