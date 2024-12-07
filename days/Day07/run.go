package Day07

import (
	"fmt"
	"strconv"
	"strings"

	"isur.dev/aoc2024/pkg"
)

type Operator string

const (
	ADD Operator = "+"
	MUL Operator = "*"
	CON Operator = "||"
)

func parseLine(line string) (int, []int) {
	fragments := strings.Split(line, ":")
	result := pkg.ParseToInt(fragments[0])
	numbers := strings.Split(fragments[1], " ")
	var nums []int
	for _, n := range numbers {
		if n == "" {
			continue
		}
		nn := pkg.ParseToInt(n)
		nums = append(nums, nn)
	}
	return result, nums
}

func concat(i int, s int) int {
	str1 := strconv.Itoa(i)
	str2 := strconv.Itoa(s)
	result := str1 + str2
	x, _ := strconv.Atoi(result)
	return x
}

var cache map[int][][]Operator = make(map[int][][]Operator)

func checkLine(result int, numbers []int) int {
	ops := []Operator{ADD, MUL, CON}
	p := len(numbers) - 1
	var operators [][]Operator
	operators, ok := cache[p]
	if ok == false {
		operators = pkg.AllPermutations(ops, p)
		cache[p] = operators
	}
	res := numbers[0]
	for _, perm := range operators {
		for i, o := range perm {
			if o == MUL {
				res *= numbers[i+1]
			} else if o == ADD {
				res += numbers[i+1]
			} else if o == CON {
				res = concat(res, numbers[i+1])
			}
		}

		if res == result {
			return result
		} else {
			res = numbers[0]
		}
	}

	return 0
}

func Run(data []string) {
	sum := 0
	for _, line := range data {
		res, nums := parseLine(line)
		sum += checkLine(res, nums)
	}
	fmt.Println(sum)
}
