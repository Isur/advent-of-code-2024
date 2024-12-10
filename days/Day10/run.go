package Day10

import (
	"fmt"

	"isur.dev/aoc2024/pkg"
)

type Point struct {
	x int
	y int
}

func checkPossible(point Point, sizeX int, sizeY int) bool {
	if point.x >= sizeX {
		return false
	}
	if point.x < 0 {
		return false
	}
	if point.y < 0 {
		return false
	}
	if point.y >= sizeY {
		return false
	}

	return true
}

func findDirs(hikeMap [][]int, point Point) []Point {
	sizeX := len(hikeMap)
	sizeY := len(hikeMap[0])
	dirs := []Point{}
	curr := hikeMap[point.x][point.y]

	left := Point{x: point.x, y: point.y - 1}
	possible := checkPossible(left, sizeX, sizeY)
	if possible {
		num := hikeMap[left.x][left.y]
		if num == curr+1 {
			dirs = append(dirs, left)
		}
	}
	right := Point{x: point.x, y: point.y + 1}
	possible = checkPossible(right, sizeX, sizeY)
	if possible {
		num := hikeMap[right.x][right.y]
		if num == curr+1 {
			dirs = append(dirs, right)
		}
	}

	top := Point{x: point.x - 1, y: point.y}
	possible = checkPossible(top, sizeX, sizeY)
	if possible {
		num := hikeMap[top.x][top.y]
		if num == curr+1 {
			dirs = append(dirs, top)
		}
	}

	bottom := Point{x: point.x + 1, y: point.y}
	possible = checkPossible(bottom, sizeX, sizeY)
	if possible {
		num := hikeMap[bottom.x][bottom.y]
		if num == curr+1 {
			dirs = append(dirs, bottom)
		}
	}
	return dirs
}

func findNext(hikeMap [][]int, path []Point) [][]Point {
	startPoint := path[len(path)-1]
	points := findDirs(hikeMap, startPoint)

	paths := [][]Point{}

	for _, p := range points {
		pa := make([]Point, len(path))
		copy(pa, path)
		if len(path) == 0 {
			pa = append(pa, startPoint, p)
		} else {
			pa = append(pa, p)
		}

		paths = append(paths, pa)
	}

	return paths
}

func findStartingPoints(hikeMap [][]string) []Point {
	points := []Point{}
	for x, row := range hikeMap {
		for y := range row {
			if hikeMap[x][y] == "0" {
				points = append(points, Point{x, y})
			}
		}
	}
	return points
}

func parseAllToInt(arr [][]string) [][]int {
	arrint := [][]int{}

	for i := range arr {
		ints := []int{}
		for _, y := range arr[i] {
			ints = append(ints, pkg.ParseToInt(y))
		}
		arrint = append(arrint, ints)
	}
	return arrint
}

func calcResult1(hikeMap [][]int, resPath [][]Point) int {
	sum := 0
	m := make(map[Point]bool)
	for _, p := range resPath {
		last := p[len(p)-1]
		val := hikeMap[last.x][last.y]
		if val == 9 {
			_, ok := m[last]
			if ok {
			} else {
				sum++
				m[last] = true
			}
		}
	}
	return sum
}
func calcResult2(hikeMap [][]int, resPath [][]Point) int {
	sum := 0
	for _, p := range resPath {
		last := p[len(p)-1]
		val := hikeMap[last.x][last.y]
		if val == 9 {
			sum++
		}
	}
	return sum
}

func reku(hikeMap [][]int, path []Point) [][]Point {
	all := [][]Point{}
	next := findNext(hikeMap, path)

	if len(next) == 0 {
		return [][]Point{path}
	}

	for _, n := range next {
		res := reku(hikeMap, n)
		all = append(all, res...)
	}

	return all
}

func Run(data []string) {
	arr := pkg.Array2D(data)

	startPoints := findStartingPoints(arr)

	arrInt := parseAllToInt(arr)

	sum1 := 0
	sum2 := 0

	for start := range startPoints {
		path := []Point{startPoints[start]}
		resPath := reku(arrInt, path)
		res1 := calcResult1(arrInt, resPath)
		res2 := calcResult2(arrInt, resPath)
		sum1 += res1
		sum2 += res2
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}
