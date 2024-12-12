package Day11

import (
	"fmt"
	"strconv"
	"strings"

	"isur.dev/aoc2024/pkg"
)

func parseToInt(s []string) []int {
	arr := []int{}
	for i := range s {
		arr = append(arr, pkg.ParseToInt(s[i]))
	}
	return arr
}

func blink(line []int) []int {
	arr := []int{}
	for _, item := range line {
		s := strconv.Itoa(item)
		if item == 0 {
			arr = append(arr, 1)
		} else if len(s)%2 == 0 {
			half := len(s) / 2
			a := pkg.ParseToInt(s[:half])
			b := pkg.ParseToInt(s[half:])
			arr = append(arr, a, b)
		} else {
			arr = append(arr, item*2024)
		}
	}
	return arr
}

type Point struct {
	num   int
	depth int
}

var MAP = make(map[Point]int)
var K = 0

func mapper(item int, depth int) int {
	point := Point{num: item, depth: depth}
	str := strconv.Itoa(item)
	p, ok := MAP[point]

	if ok {
		return p
	}

	if depth == 0 {
		return 1
	}

	arr := []int{}
	if item == 0 {
		arr = append(arr, 1)
	} else if len(str)%2 == 0 {
		half := len(str) / 2
		a := pkg.ParseToInt(str[:half])
		b := pkg.ParseToInt(str[half:])
		arr = append(arr, a, b)
	} else {
		arr = append(arr, item*2024)
	}

	l := mapper(arr[0], depth-1)
	MAP[point] = l
	if len(arr) > 1 {
		r := mapper(arr[1], depth-1)
		MAP[point] += r
	}

	return MAP[point]
}

func blink2(line []int, depth int) int {
	sum := 0
	for i := range line {
		sum += mapper(line[i], depth)
	}
	return sum
}

func Run(data []string) {
	num := pkg.ParseToInt("11")
	fmt.Println("Day ", num)

	line := parseToInt(strings.Split(data[0], " "))
	str := data[0]
	fmt.Println(str)
	x := blink2(line, 75)
	fmt.Println(x)
}
