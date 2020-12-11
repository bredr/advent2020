package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readLines(path string) ([][]rune, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	rows := make([][]rune, len(lines))
	for i, line := range lines {
		rows[i] = []rune(line)
	}
	return rows, nil
}

type Rule struct {
	change rune
	r      func([][]rune) [][]rune
}

func merged(input [][]rune, x [][][]rune, rules ...Rule) [][]rune {
	newState := make([][]rune, len(input))
	for i, row := range input {
		newState[i] = make([]rune, len(row))
	}
	for i, row := range input {
		for j := range row {
			newState[i][j] = input[i][j]
			for r, d := range x {
				if rules[r].change == d[i][j] {
					newState[i][j] = d[i][j]
				}
			}
		}
	}
	return newState
}

func isEqual(a [][]rune, b [][]rune) bool {
	for i, row := range a {
		for j, valA := range row {
			if valA != b[i][j] {
				return false
			}
		}
	}
	return true
}

func Simulation(input [][]rune, rules ...Rule) [][]rune {
	intermediateStates := make([][][]rune, len(rules))
	for i, rule := range rules {
		intermediateStates[i] = rule.r(input)
	}

	result := merged(input, intermediateStates, rules...)
	if isEqual(result, input) {
		return input
	}
	return Simulation(result, rules...)
}

const occupied rune = '#'
const empty rune = 'L'

func main() {
	data, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}

	rule1 := Rule{
		change: occupied, r: func(x [][]rune) [][]rune {
			newState := make([][]rune, len(x))
			for i, row := range x {
				newState[i] = make([]rune, len(row))
			}
			for i, row := range x {
				for j := range row {
					if x[i][j] == empty {
						change := true
						for k := i - 1; k <= i+1; k++ {
							for l := j - 1; l <= j+1; l++ {
								if k >= 0 && k < len(x) && l >= 0 && l < len(row) {
									if x[k][l] == occupied {
										change = false
									}
								}
							}
						}
						if change {
							newState[i][j] = occupied
						}
					}
				}
			}
			return newState
		},
	}

	rule1b := Rule{
		change: occupied, r: func(x [][]rune) [][]rune {
			newState := make([][]rune, len(x))
			for i, row := range x {
				newState[i] = make([]rune, len(row))
			}
			maxRadius := len(x)
			if len(x[0]) > maxRadius {
				maxRadius = len(x[0])
			}
			for i, row := range x {
				for j := range row {
					if x[i][j] == empty {
						change := true
						directions := []func(int) (int, int){
							func(r int) (int, int) {
								return i, j + r
							},
							func(r int) (int, int) {
								return i + r, j + r
							},
							func(r int) (int, int) {
								return i + r, j
							},
							func(r int) (int, int) {
								return i + r, j - r
							},
							func(r int) (int, int) {
								return i, j - r
							},
							func(r int) (int, int) {
								return i - r, j - r
							},
							func(r int) (int, int) {
								return i - r, j + r
							},
							func(r int) (int, int) {
								return i - r, j
							},
						}
						for _, posFn := range directions {
							for r := 1; r <= maxRadius; r++ {
								k, l := posFn(r)
								if k >= 0 && k < len(x) && l >= 0 && l < len(row) {
									if x[k][l] == empty {
										break
									}
									if x[k][l] == occupied {
										change = false
									}
								} else {
									break
								}
							}
						}
						if change {
							newState[i][j] = occupied
						}
					}
				}
			}
			return newState
		},
	}

	rule2 := Rule{
		change: empty,
		r: func(x [][]rune) [][]rune {
			newState := make([][]rune, len(x))
			for i, row := range x {
				newState[i] = make([]rune, len(row))
			}
			for i, row := range x {
				for j := range row {
					if x[i][j] == occupied {
						adjacentOccupied := 0
						for k := i - 1; k <= i+1; k++ {
							for l := j - 1; l <= j+1; l++ {
								if k >= 0 && k < len(x) && l >= 0 && l < len(row) {
									if x[k][l] == occupied {
										adjacentOccupied++
									}
								}
							}
						}
						if adjacentOccupied >= 5 {
							newState[i][j] = empty
						}
					}
				}
			}
			return newState
		},
	}

	rule2b := Rule{
		change: empty,
		r: func(x [][]rune) [][]rune {
			newState := make([][]rune, len(x))
			for i, row := range x {
				newState[i] = make([]rune, len(row))
			}
			maxRadius := len(x)
			if len(x[0]) > maxRadius {
				maxRadius = len(x[0])
			}
			for i, row := range x {
				for j := range row {
					if x[i][j] == occupied {
						adjacentOccupied := 0
						directions := []func(int) (int, int){
							func(r int) (int, int) {
								return i, j + r
							},
							func(r int) (int, int) {
								return i + r, j + r
							},
							func(r int) (int, int) {
								return i + r, j
							},
							func(r int) (int, int) {
								return i + r, j - r
							},
							func(r int) (int, int) {
								return i, j - r
							},
							func(r int) (int, int) {
								return i - r, j - r
							},
							func(r int) (int, int) {
								return i - r, j + r
							},
							func(r int) (int, int) {
								return i - r, j
							},
						}
						for _, posFn := range directions {
							for r := 1; r <= maxRadius; r++ {
								k, l := posFn(r)
								if k >= 0 && k < len(x) && l >= 0 && l < len(row) {
									if x[k][l] == occupied {
										adjacentOccupied++
										break
									}
									if x[k][l] == empty {
										break
									}
								} else {
									break
								}
							}
						}
						if adjacentOccupied >= 5 {
							newState[i][j] = empty
						}
					}
				}
			}
			return newState
		},
	}

	result1 := Simulation(data, rule1, rule2)
	fmt.Printf("result1 = %d\n", countOccupied(result1))

	result2 := Simulation(data, rule1b, rule2b)
	fmt.Printf("result2 = %d\n", countOccupied(result2))
}

func countOccupied(x [][]rune) int {
	count := 0
	for _, row := range x {
		for _, seat := range row {
			if seat == occupied {
				count++
			}
		}
	}
	return count
}
