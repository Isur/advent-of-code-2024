package Day09

import (
	"fmt"
	"sort"
	"strconv"

	"isur.dev/aoc2024/pkg"
)

func decode(line string) []string {
	decoded := []string{}
	free := false
	id := 0

	for _, c := range line {
		char := string(c)
		num := pkg.ParseToInt(char)

		for range num {
			if free == false {
				decoded = append(decoded, strconv.Itoa(id))
			} else {
				decoded = append(decoded, ".")
			}

		}

		if free == false {
			id++
		}
		free = !free
	}

	return decoded

}

func Swap[T any](arr []T, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func defragment1(line []string) []string {
	j := len(line) - 1
	for i := 0; i < len(line); i++ {
		if line[i] == "." {
			for j > i {
				if line[j] != "." {
					Swap(line, i, j)
					break
				}
				j--
			}
		}
	}
	return line
}

func defragment2(line []string) []string {
	m := make(map[string]int)
	f := make(map[int]int)

	dot := 0
	for i, c := range line {
		if c == "." && dot == 0 {
			dot = i
			f[dot] = 1
			continue
		} else if c == "." && dot != 0 {
			f[dot]++
			continue
		}

		dot = 0

		if _, ok := m[c]; ok {
			m[c]++
		} else {
			m[c] = 1
		}
	}

	order := make([]int, 0, len(f))
	for k := range f {
		order = append(order, k)
	}
	sort.Slice(order, func(i, j int) bool {
		return order[i] < order[j]
	})

	done := make(map[string]bool)

	for i := len(line) - 1; i >= 0; i-- {
		if line[i] == "." {
			continue
		}
		c := line[i]
		l := 0
		if _, ok := m[c]; ok {
			l = m[c]
		}

		if dx := done[c]; dx {
			continue
		}

		for _, o := range order {
			oVal := f[o]
			if oVal >= l && o < i {
				first := i - l + 1

				for x := 0; x < l; x++ {
					Swap(line, first+x, o+x)
				}

				done[c] = true

				f[o+l] = f[o] - l
				delete(f, o)

				if f[o+l] == 0 {
					delete(f, o+l)
				}

				order = []int{}
				for k := range f {
					order = append(order, k)
				}
				sort.Slice(order, func(i, j int) bool {
					return order[i] < order[j]
				})

				break
			}
		}
	}
	return line
}

func checkSum(line []string) int {
	sum := 0
	for i, c := range line {
		if c == "." {
			continue
		}
		num := pkg.ParseToInt(c)
		sum += num * i
	}

	return sum

}

func Run(data []string) {
	num := pkg.ParseToInt("09")
	fmt.Println("Day ", num)
	data2 := make([]string, len(data))
	copy(data2, data)

	line := data[0]
	decoded := decode(line)
	defragmented1 := defragment1(decoded)
	sum1 := checkSum(defragmented1)
	fmt.Println(sum1)

	line2 := data2[0]
	decoded2 := decode(line2)
	defragmented2 := defragment2(decoded2)
	sum2 := checkSum(defragmented2)
	fmt.Println(sum2)
}
