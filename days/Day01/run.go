package Day01

import (
	"fmt"
	"sort"
	"strings"

	"isur.dev/aoc2024/pkg"
)

func Run(data []string) {
	var firstList []int
	var secondList []int

	for _, line := range data {
		res := strings.Split(line, "   ")
		firstNum := pkg.ParseToInt(res[0])
		secondNum := pkg.ParseToInt(res[1])

		firstList = append(firstList, firstNum)
		secondList = append(secondList, secondNum)
	}

	sort.Slice(firstList, func(i, j int) bool {
		return firstList[i] < firstList[j]
	})
	sort.Slice(secondList, func(i, j int) bool {
		return secondList[i] < secondList[j]
	})
	sum := 0
	for i := 0; i < len(firstList); i++ {
		a := firstList[i]
		b := secondList[i]
		if a < b {
			sum += b - a
		} else {
			sum += a - b
		}
	}
	var m map[int]int = make(map[int]int)

	for i := 0; i < len(secondList); i++ {
		a := secondList[i]
		val, ok := m[a]
		if ok {
			m[a] = val + 1
		} else {
			m[a] = 1
		}
	}

	diff := 0

	for i := 0; i < len(firstList); i++ {
		a := firstList[i]
		val, ok := m[a]
		if ok {
			diff = diff + a*val
		}

	}

	fmt.Println("Sum:  ", sum)
	fmt.Println("Diff: ", diff)
}
