package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	var setMask uint64
	var clrMask uint64
	memory := make(map[uint64]uint64)
	for _, line := range lines {
		if strings.Contains(line, "mask") {
			line = strings.Replace(line, "mask = ", "", -1)
			line = strings.TrimSpace(line)
			setMask, clrMask = processMask(line)
		}
		if strings.Contains(line, "mem") {
			parts := strings.Split(line, "] = ")
			value, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			parts[0] = strings.Replace(parts[0], "mem[", "", -1)
			address, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			v := uint64(value)
			v |= setMask
			v &= clrMask
			memory[uint64(address)] = v
		}
	}
	sum := uint64(0)
	for _, v := range memory {
		sum += v
	}
	fmt.Printf("result1 = %d\n", sum)
}

func part2() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	var mask string
	memory := make(map[uint64]uint64)
	for _, line := range lines {
		if strings.Contains(line, "mask") {
			line = strings.Replace(line, "mask = ", "", -1)
			mask = strings.TrimSpace(line)
		}
		if strings.Contains(line, "mem") {
			parts := strings.Split(line, "] = ")
			value, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			parts[0] = strings.Replace(parts[0], "mem[", "", -1)
			address, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			addresses := []uint64{0}
			for i := 0; i < 36; i++ {
				bitIndex := 35 - i
				switch mask[i] {
				case '0':
					for j := range addresses {
						if bitAt(uint64(address), bitIndex) == 1 {
							addresses[j] = setBit(addresses[j], bitIndex)
						}
					}
				case '1':
					for j := range addresses {
						addresses[j] = setBit(addresses[j], bitIndex)
					}
				case 'X':
					for j := range addresses {
						addresses = append(addresses, addresses[j])
						addresses[j] = setBit(addresses[j], bitIndex)
					}
				}
			}
			for _, address := range addresses {
				memory[address] = uint64(value)
			}
		}
	}
	sum := uint64(0)
	for _, v := range memory {
		sum += v
	}
	fmt.Printf("result2 = %d\n", sum)
}

func processMask(s string) (setMask uint64, clrMask uint64) {
	for i := range s {
		c := s[len(s)-1-i]
		switch c {
		case '1':
			setMask |= (1 << i)
		case '0':
			clrMask |= (1 << i)
		}
	}
	return setMask, ^clrMask
}

func bitAt(val uint64, index int) uint64 {
	mask := uint64(1) << index
	return (val & mask) >> index
}

func setBit(val uint64, index int) uint64 {
	return val | (1 << index)
}
