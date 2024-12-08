package Day08

import (
	"fmt"

	"isur.dev/aoc2024/pkg"
)

const EMPTY = "."
const ANTINODE = "#"

type Position struct {
	x int
	y int
}

type AllPositions struct {
	next      Position
	positions []Position
}

func findPositions(arr [][]string) map[string][]Position {
	positionMap := make(map[string][]Position)

	for y := range arr {
		for x := range arr[y] {
			if arr[x][y] != EMPTY {
				item := arr[x][y]
				pos := Position{x, y}
				posArr, ok := positionMap[item]
				if ok {
					posArr = append(posArr, pos)
					positionMap[item] = posArr
				} else {
					positionMap[item] = []Position{pos}
				}
			}
		}
	}

	return positionMap
}

func filterNodes(nodes []Position, size Position) []Position {
	filtered := []Position{}
	for _, pos := range nodes {
		if pos.x < 0 || pos.x > size.x {
			continue
		}
		if pos.y < 0 || pos.y > size.y {
			continue
		}

		filtered = append(filtered, pos)
	}

	return filtered
}

func findAntinodes(positions map[string][]Position, size Position) []Position {
	antinodes := make(map[Position]bool)
	for _, value := range positions {
		if len(value) <= 1 {
			continue
		}

		for i, first := range value {
			for j, pair := range value {
				if i == j {
					continue
				}

				distance := Position{x: pair.x - first.x, y: pair.y - first.y}

				i := 0
				for {
					dist := Position{
						x: distance.x * i,
						y: distance.y * i,
					}
					if pair.x+dist.x <= size.x && pair.x+dist.x >= 0 && pair.y+dist.y <= size.y && pair.y+dist.y >= 0 {
						antinode := Position{x: pair.x + dist.x, y: pair.y + dist.y}
						_, ok := antinodes[antinode]
						if ok == false {
							antinodes[antinode] = true
						}

						dist = Position{x: pair.x - first.x, y: pair.y - first.y}
						i++
					} else {
						break
					}
				}

				i = 0
				for {
					dist := Position{
						x: distance.x * i,
						y: distance.y * i,
					}
					if first.x-dist.x <= size.x && first.x-dist.x >= 0 && first.y-dist.y <= size.y && first.y-dist.y >= 0 {
						antinode := Position{x: first.x - dist.x, y: first.y - dist.y}
						_, ok := antinodes[antinode]
						if ok == false {
							antinodes[antinode] = true
						}
						i++
					} else {
						break
					}
				}
			}
		}
	}

	nodes := []Position{}
	for key, _ := range antinodes {
		nodes = append(nodes, key)
	}

	return nodes
}

func Run(data []string) {
	arr := pkg.Array2D(data)
	for i := range arr {
		for _, y := range arr[i] {
			fmt.Print(y, " ")
		}
		fmt.Println()
	}

	positions := findPositions(arr)
	size := Position{
		x: len(arr) - 1,
		y: len(arr[0]) - 1,
	}
	nodes := findAntinodes(positions, size)

	for _, node := range nodes {
		fmt.Println(node)
	}
	fmt.Println(len(nodes))
}
