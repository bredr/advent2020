package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type position struct {
	x     int
	y     int
	theta int
}

type instruction struct {
	op  rune
	arg int
}

func readLines(path string) ([]instruction, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	instructions := make([]instruction, len(lines))
	for i, line := range lines {
		runes := []rune(line)
		arg, err := strconv.Atoi(string(runes[1:]))
		if err != nil {
			return nil, err
		}
		instructions[i] = instruction{
			op:  runes[0],
			arg: arg,
		}
	}
	return instructions, nil
}

func Move(instructions []instruction, pos position, index int) position {
	if index >= len(instructions) {
		return pos
	}
	i := instructions[index]
	switch i.op {
	case 'N':
		pos.y += i.arg
	case 'S':
		pos.y -= i.arg
	case 'E':
		pos.x += i.arg
	case 'W':
		pos.x -= i.arg
	case 'L':
		pos.theta += i.arg
	case 'R':
		pos.theta -= i.arg
	case 'F':
		rads := float64(pos.theta) * math.Pi / 180.0
		pos.x += int(float64(i.arg) * math.Cos(rads))
		pos.y += int(float64(i.arg) * math.Sin(rads))
	default:
		panic("unknown operation")
	}

	return Move(instructions, pos, index+1)
}

func MoveWithWaypoint(instructions []instruction, ship position, waypoint position, index int) position {
	if index >= len(instructions) {
		return ship
	}
	i := instructions[index]
	switch i.op {
	case 'N':
		waypoint.y += i.arg
	case 'S':
		waypoint.y -= i.arg
	case 'E':
		waypoint.x += i.arg
	case 'W':
		waypoint.x -= i.arg
	case 'L':
		switch i.arg {
		case 90:
			waypoint.x, waypoint.y = -waypoint.y, waypoint.x
		case 180:
			waypoint.x, waypoint.y = -waypoint.x, -waypoint.y
		case 270:
			waypoint.x, waypoint.y = waypoint.y, -waypoint.x
		default:
			panic("unknown angle")
		}
	case 'R':
		switch i.arg {
		case 90:
			waypoint.x, waypoint.y = waypoint.y, -waypoint.x
		case 180:
			waypoint.x, waypoint.y = -waypoint.x, -waypoint.y
		case 270:
			waypoint.x, waypoint.y = -waypoint.y, waypoint.x
		default:
			panic("unknown angle")
		}
	case 'F':
		ship.x += i.arg * waypoint.x
		ship.y += i.arg * waypoint.y
	default:
		panic("unknown operation")
	}
	return MoveWithWaypoint(instructions, ship, waypoint, index+1)
}

func main() {
	data, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}
	result1 := Move(data, position{x: 0, y: 0, theta: 0}, 0)
	fmt.Printf("result 1 = %d\n", abs(result1.x)+abs(result1.y))

	result2 := MoveWithWaypoint(data, position{x: 0, y: 0}, position{x: 10, y: 1}, 0)
	fmt.Printf("result 2 = %d\n", abs(result2.x)+abs(result2.y))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
