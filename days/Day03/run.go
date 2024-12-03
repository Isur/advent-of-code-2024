package Day03

import (
	"fmt"
	"regexp"

	"isur.dev/aoc2024/pkg"
)

func Run(data []string) {
	sum := 0
	for _, value := range data {
		res := parseData(value)
		sum = sum + res
	}

	fmt.Println(sum)
}

func parseData(str string) int {
	sum := 0
	do := true
	for i := 0; i < len(str); i++ {
		dontExp, _ := regexp.Compile("^don't()")

		if dontExp.MatchString(str[i:]) {
			do = false
			continue
		}
		doExp, _ := regexp.Compile("^do()")
		if doExp.MatchString(str[i:]) {
			do = true
			continue
		}

		if do {
			r, _ := regexp.Compile("^mul\\(([\\d]{1,3}),([\\d]{1,3})\\)")
			matches := r.FindStringSubmatch(str[i:])
			if len(matches) == 3 {
				num1 := pkg.ParseToInt(matches[1])
				num2 := pkg.ParseToInt(matches[2])
				sum = sum + num1*num2
			}
		}
	}
	return sum
}
