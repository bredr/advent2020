package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type input struct {
	timestamp int
	buses     []int
}

func readInput(path string) (*input, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")

	timestamp, err := strconv.Atoi(lines[0])
	if err != nil {
		return nil, err
	}
	allBuses := strings.Split(lines[1], ",")
	var buses []int
	for _, v := range allBuses {
		if v != "x" {
			bus, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			buses = append(buses, bus)
		}
	}
	return &input{timestamp: timestamp, buses: buses}, nil
}

func readInputPart2(path string) ([]int, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	allBuses := strings.Split(lines[1], ",")
	buses := make([]int, len(allBuses))
	for i, v := range allBuses {
		if v != "x" {
			bus, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			buses[i] = bus
		} else {
			buses[i] = 0
		}
	}
	return buses, nil
}

type bestBus struct {
	bus         int
	waitingTime int
}

func BestBus(i *input, b bestBus, index int) bestBus {
	if index >= len(i.buses) {
		return b
	}
	bus := i.buses[index]
	waitingTime := bus*(1+i.timestamp/bus) - i.timestamp
	if waitingTime < b.waitingTime {
		b.bus = bus
		b.waitingTime = waitingTime
	}
	return BestBus(i, b, index+1)
}

// greatest common divisor (GCD) via Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)
	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}
	return result
}

func LowestTimestamp(buses []int) int {
	t0 := 1
	for {
		complete := true
		var matches []int
		for t, b := range buses {
			if b != 0 {
				if (t0+t)%b == 0 {
					matches = append(matches, b)
				} else {
					complete = false
				}
			}
		}
		if complete {
			return t0
		}
		switch len(matches) {
		case 0:
			t0++
		case 1:
			t0 += matches[0]
		case 2:
			t0 += lcm(matches[0], matches[1])
		default:
			t0 += lcm(matches[0], matches[1], matches[2:]...)
		}
	}
}

func main() {
	i, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	result1 := BestBus(i, bestBus{waitingTime: 1 << 22}, 0)

	fmt.Printf("result1 = %d\n", result1.bus*result1.waitingTime)

	i2, err := readInputPart2("input.txt")
	if err != nil {
		panic(err)
	}
	result2 := LowestTimestamp(i2)
	fmt.Printf("result2 = %d\n", result2)

}
