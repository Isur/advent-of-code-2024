package Day06

import (
	"fmt"

	"isur.dev/aoc2024/pkg"
)

type ELEM string

const (
	UP      ELEM = "^"
	DOWN         = "v"
	LEFT         = "<"
	RIGHT        = ">"
	NOTHING      = "."
	BLOCK        = "#"
)

type Guard struct {
	x   int
	y   int
	dir ELEM
}

func findGuard(data [][]ELEM) Guard {
	for y := range data {
		for x := range data[y] {
			if data[y][x] == UP {
				return Guard{x, y, UP}
			}
		}
	}

	return Guard{}
}

func Go(data [][]ELEM, guard *Guard) bool {
	data[guard.y][guard.x] = "X"
	switch guard.dir {
	case UP:
		if guard.y < 1 {
			return false
		} else if data[guard.y-1][guard.x] == BLOCK {
			guard.dir = RIGHT
		} else {
			guard.y--
		}
	case DOWN:
		if guard.y == len(data)-1 {
			return false
		} else if data[guard.y+1][guard.x] == BLOCK {
			guard.dir = LEFT
		} else {
			guard.y++
		}
	case LEFT:
		if guard.x < 1 {
			return false
		} else if data[guard.y][guard.x-1] == BLOCK {
			guard.dir = UP
		} else {
			guard.x--
		}
	case RIGHT:
		if guard.x == len(data[0])-1 {
			return false
		} else if data[guard.y][guard.x+1] == BLOCK {
			guard.dir = DOWN
		} else {
			guard.x++
		}
	}
	data[guard.y][guard.x] = guard.dir
	return true
}

func RunRoute(data [][]ELEM) ([][]ELEM, bool) {
	guard := findGuard(data)
	var pos map[string]bool = make(map[string]bool)

	for Go(data, &guard) {
		p := fmt.Sprintf("%d,%d,%s", guard.x, guard.y, guard.dir)
		if pos[p] {
			return data, true
		} else {
			pos[p] = true
		}
	}

	return data, false
}

func parseInput(arr [][]string) [][]ELEM {
	var res [][]ELEM
	for _, line := range arr {
		var resLine []ELEM
		for _, elem := range line {
			resLine = append(resLine, ELEM(elem))
		}
		res = append(res, resLine)
	}
	return res
}

func Count(data [][]ELEM) int {
	res := 0
	for _, line := range data {
		for _, elem := range line {
			if elem == "X" {
				res++
			}
		}
	}
	return res
}

func cp(arr [][]ELEM) [][]ELEM {
	var res [][]ELEM
	for _, line := range arr {
		var resLine []ELEM
		for _, elem := range line {
			resLine = append(resLine, elem)
		}
		res = append(res, resLine)
	}
	return res
}

func Run(data []string) {
	arr2D := pkg.Array2D(data)
	arrO := parseInput(arr2D)
	cnt := 0

	for y := range arrO {
		for x := range arrO[y] {
			arr := cp(arrO)
			if arr[y][x] == NOTHING {
				arr[y][x] = BLOCK
			}
			_, loop := RunRoute(arr)
			if loop {
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}
