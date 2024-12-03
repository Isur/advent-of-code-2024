package Day02

import (
	"fmt"
	"strings"

	"isur.dev/aoc2024/pkg"
)

func Run(data []string) {
	sum := 0
	for _, line := range data {
		numbers := LineParse(line)
		ok := CheckIfOk(numbers)
		if ok {
			sum = sum + 1
		} else {
			for i := range numbers {
				kek := pkg.RemoveFromSlice(numbers, i)
				xd := CheckIfOk(kek)
				if xd {
					sum = sum + 1
					break
				}
			}
		}
	}

	fmt.Printf("Result: %d\n", sum)
}

func LineParse(line string) []int {
	var numbers []int
	for _, element := range strings.Split(line, " ") {
		num := pkg.ParseToInt(element)
		numbers = append(numbers, num)
	}

	return numbers
}

func CheckIfOk(numbers []int) bool {
	var prev int
	var growing bool

	for index, num := range numbers {
		if index == 0 {
			prev = num
			continue
		} else {
			if index == 1 {
				if num > prev {
					growing = true
				} else if prev > num {
					growing = false
				} else {
					return false
				}
			} else if num > prev && growing == false {
				return false
			} else if num < prev && growing == true {
				return false
			}

			diff := pkg.Abs(prev, num)
			if diff > 3 || diff < 1 {
				return false
			}

			prev = numbers[index]
		}
	}
	return true
}
