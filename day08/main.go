package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type op string

const acc op = "acc"
const jmp op = "jmp"
const nop op = "nop"

type Operation struct {
	op      op
	arg     int
	visited bool
}

func readOperations(path string) ([]Operation, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	ops := make([]Operation, len(lines))
	for i, line := range lines {
		s := strings.Split(line, " ")
		op := nop
		switch s[0] {
		case "acc":
			op = acc
		case "jmp":
			op = jmp
		case "nop":
			op = nop
		}

		arg, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		ops[i] = Operation{
			op:      op,
			arg:     arg,
			visited: false,
		}
	}
	return ops, nil
}

func Compute(operations []*Operation, accumulator, index int) (finalAccumulator int, finalIndex int) {
	if index >= len(operations) {
		return accumulator, len(operations)
	}
	op := operations[index]
	if op.visited {
		return accumulator, index
	}
	op.visited = true
	switch op.op {
	case acc:
		return Compute(operations, accumulator+op.arg, index+1)
	case jmp:
		return Compute(operations, accumulator, index+op.arg)
	case nop:
		return Compute(operations, accumulator, index+1)
	default:
		panic("unknown operation")
	}
}

func GenerateOperations(ops []Operation, index *int) []*Operation {
	o := make([]*Operation, len(ops))
	for i, operation := range ops {
		if index != nil && i == *index {
			switch operation.op {
			case jmp:
				o[i] = &Operation{op: nop, arg: operation.arg, visited: false}
			case nop:
				o[i] = &Operation{op: jmp, arg: operation.arg, visited: false}
			default:
				panic("shouldn't try swapping")
			}
		} else {
			o[i] = &Operation{op: operation.op, arg: operation.arg, visited: false}
		}
	}
	return o
}

func main() {
	operations, err := readOperations("input.txt")
	if err != nil {
		panic(err)
	}
	rawOps := GenerateOperations(operations, nil)
	result, _ := Compute(rawOps, 0, 0)
	fmt.Printf("result1 = %d\n", result)

	for i, op := range operations {
		if op.op == jmp || op.op == nop {
			acc, finalIndex := Compute(GenerateOperations(operations, &i), 0, 0)
			if finalIndex == len(operations) {
				fmt.Printf("result2 = %d\n", acc)
				break
			}
		}
	}

}
