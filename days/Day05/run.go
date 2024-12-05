package Day05

import (
	"fmt"
	"strings"

	"isur.dev/aoc2024/pkg"
)

type Rule struct {
	before int
	after  int
}

func parseData(data []string) ([]Rule, [][]int) {
	var rules []Rule
	var updates [][]int
	checkRule := true
	for _, line := range data {
		if line == "" {
			checkRule = false
			continue
		}
		if checkRule {
			nums := strings.Split(line, "|")
			before := pkg.ParseToInt(nums[0])
			after := pkg.ParseToInt(nums[1])
			rules = append(rules, Rule{
				before: before,
				after:  after,
			})
		} else {
			nums := strings.Split(line, ",")
			var row []int
			for _, num := range nums {
				row = append(row, pkg.ParseToInt(num))
			}
			updates = append(updates, row)
		}
	}

	return rules, updates

}

func checkRules(row []int, rules []Rule) (int, Rule) {
	seen := []int{}
	for _, item := range row {
		for _, rule := range rules {
			if item == rule.before {
				for _, s := range seen {
					if s == rule.after {
						return 0, rule
					}
				}
			}
		}
		seen = append(seen, item)
	}

	return row[len(row)/2], rules[0]
}

func Swap(slice []int, a int, b int) []int {
	ia := 0
	ib := 0
	for index, item := range slice {
		if item == a {
			ia = index
		}
		if item == b {
			ib = index
		}
	}
	slice[ia], slice[ib] = slice[ib], slice[ia]
	return slice
}

func Run(data []string) {
	rules, updates := parseData(data)
	sum := 0
	incSum := 0
	for _, row := range updates {
		res, failedRule := checkRules(row, rules)
		sum += res

		for res == 0 {
			Swap(row, failedRule.before, failedRule.after)
			res, failedRule = checkRules(row, rules)
			incSum += res
		}
	}

	fmt.Println(sum)
	fmt.Println(incSum)
}
