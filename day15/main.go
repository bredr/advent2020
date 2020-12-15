package main

import "fmt"

func game(input []int, limit int) int {
	lastAppearances := make(map[int]int)
	for i, v := range input {
		lastAppearances[v] = i
	}

	turnNumber := len(input)
	lastTurn := input[len(input)-1]
	for {
		if turnNumber == limit {
			return lastTurn
		}
		i, ok := lastAppearances[lastTurn]
		if !ok {
			lastAppearances[lastTurn] = turnNumber - 1
			lastTurn = 0
		} else {
			lastAppearances[lastTurn] = turnNumber - 1
			lastTurn = turnNumber - i - 1
		}
		turnNumber++
	}
}

func main() {
	lastValue := game([]int{7, 12, 1, 0, 16, 2}, 2020)
	fmt.Printf("result1 = %d\n", lastValue)
	lastValue = game([]int{7, 12, 1, 0, 16, 2}, 30000000)
	fmt.Printf("result2 = %d\n", lastValue)
}
