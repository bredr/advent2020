package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func activeToInactive(current map[[3]int]struct{}) map[[3]int]struct{} {
	next := make(map[[3]int]struct{})
	for k := range current {
		k := k
		activeNeighbours := 0
		for x := k[0] - 1; x <= k[0]+1; x++ {
			for y := k[1] - 1; y <= k[1]+1; y++ {
				for z := k[2] - 1; z <= k[2]+1; z++ {
					if !(x == k[0] && y == k[1] && z == k[2]) {
						if _, ok := current[[3]int{x, y, z}]; ok {
							activeNeighbours++
						}
					}
				}
			}
		}
		if activeNeighbours == 2 || activeNeighbours == 3 {
			next[k] = struct{}{}
		}
	}
	return next
}

func activeToInactive4D(current map[[4]int]struct{}) map[[4]int]struct{} {
	next := make(map[[4]int]struct{})
	for k := range current {
		k := k
		activeNeighbours := 0
		for x := k[0] - 1; x <= k[0]+1; x++ {
			for y := k[1] - 1; y <= k[1]+1; y++ {
				for z := k[2] - 1; z <= k[2]+1; z++ {
					for w := k[3] - 1; w <= k[3]+1; w++ {
						if !(x == k[0] && y == k[1] && z == k[2] && w == k[3]) {
							if _, ok := current[[4]int{x, y, z, w}]; ok {
								activeNeighbours++
							}
						}
					}
				}
			}
		}
		if activeNeighbours == 2 || activeNeighbours == 3 {
			next[k] = struct{}{}
		}
	}
	return next
}

func inactiveToActive(current map[[3]int]struct{}) map[[3]int]struct{} {
	next := make(map[[3]int]struct{})
	var minX, maxX, minY, maxY, minZ, maxZ int
	for k := range current {
		if k[0] < minX {
			minX = k[0]
		} else if k[0] > maxX {
			maxX = k[0]
		}
		if k[1] < minY {
			minY = k[1]
		} else if k[1] > maxY {
			maxY = k[1]
		}
		if k[2] < minZ {
			minZ = k[2]
		} else if k[2] > maxZ {
			maxZ = k[2]
		}
	}
	for x := minX - 1; x <= maxX+1; x++ {
		for y := minY - 1; y <= maxY+1; y++ {
			for z := minZ - 1; z <= maxZ+1; z++ {
				if _, ok := current[[3]int{x, y, z}]; !ok {
					activeNeighbours := 0
					for x1 := x - 1; x1 <= x+1; x1++ {
						for y1 := y - 1; y1 <= y+1; y1++ {
							for z1 := z - 1; z1 <= z+1; z1++ {
								if _, ok := current[[3]int{x1, y1, z1}]; ok {
									activeNeighbours++
								}
							}
						}
					}
					if activeNeighbours == 3 {
						next[[3]int{x, y, z}] = struct{}{}
					}
				}
			}
		}

	}
	return next
}

func inactiveToActive4D(current map[[4]int]struct{}) map[[4]int]struct{} {
	next := make(map[[4]int]struct{})
	var minX, maxX, minY, maxY, minZ, maxZ, minW, maxW int
	for k := range current {
		if k[0] < minX {
			minX = k[0]
		} else if k[0] > maxX {
			maxX = k[0]
		}
		if k[1] < minY {
			minY = k[1]
		} else if k[1] > maxY {
			maxY = k[1]
		}
		if k[2] < minZ {
			minZ = k[2]
		} else if k[2] > maxZ {
			maxZ = k[2]
		}
		if k[3] < minW {
			minW = k[3]
		} else if k[3] > maxW {
			maxW = k[3]
		}
	}
	for x := minX - 1; x <= maxX+1; x++ {
		for y := minY - 1; y <= maxY+1; y++ {
			for z := minZ - 1; z <= maxZ+1; z++ {
				for w := minW - 1; w <= maxW+1; w++ {
					if _, ok := current[[4]int{x, y, z, w}]; !ok {
						activeNeighbours := 0
						for x1 := x - 1; x1 <= x+1; x1++ {
							for y1 := y - 1; y1 <= y+1; y1++ {
								for z1 := z - 1; z1 <= z+1; z1++ {
									for w1 := w - 1; w1 <= w+1; w1++ {
										if _, ok := current[[4]int{x1, y1, z1, w1}]; ok {
											activeNeighbours++
										}
									}
								}
							}
						}
						if activeNeighbours == 3 {
							next[[4]int{x, y, z, w}] = struct{}{}
						}
					}
				}
			}
		}
	}
	return next
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	active := make(map[[3]int]struct{})

	lines := strings.Split(string(content), "\n")
	for y, line := range lines {
		for x, value := range []rune(line) {
			if value == '#' {
				active[[3]int{x, y, 0}] = struct{}{}
			}
		}
	}

	for i := 1; i <= 6; i++ {
		next1 := activeToInactive(active)
		next2 := inactiveToActive(active)
		active = make(map[[3]int]struct{})
		for k := range next1 {
			active[k] = struct{}{}
		}
		for k := range next2 {
			active[k] = struct{}{}
		}
	}
	fmt.Printf("result1 = %d\n", len(active))

	active2 := make(map[[4]int]struct{})
	for y, line := range lines {
		for x, value := range []rune(line) {
			if value == '#' {
				active2[[4]int{x, y, 0, 0}] = struct{}{}
			}
		}
	}

	for i := 1; i <= 6; i++ {
		next1 := activeToInactive4D(active2)
		next2 := inactiveToActive4D(active2)
		active2 = make(map[[4]int]struct{})
		for k := range next1 {
			active2[k] = struct{}{}
		}
		for k := range next2 {
			active2[k] = struct{}{}
		}
	}
	fmt.Printf("result2 = %d\n", len(active2))

}
