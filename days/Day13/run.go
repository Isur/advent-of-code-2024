package Day13

import (
	"fmt"
	"regexp"

	"isur.dev/aoc2024/pkg"
)

type F struct {
	ax float64
	ay float64
	bx float64
	by float64
	x  float64
	y  float64
}

func calc(f F) int64 {
	W := f.ax*f.by - f.ay*f.bx

	W1 := f.x*f.by - f.y*f.bx
	W2 := f.ax*f.y - f.ay*f.x

	a := W1 / W
	b := W2 / W

	if float64(int64(a)) != a || float64(int64(b)) != b {
		return 0
	}

	return int64(a*3 + b)
}

func Run(data []string) {
	num := pkg.ParseToInt("13")
	fmt.Println("Day ", num)

	i := 0
	funcs := []F{}

	f := F{0, 0, 0, 0, 0, 0}
	for _, d := range data {
		reg, _ := regexp.Compile("\\d+")
		res := reg.FindAllString(d, 2)
		if i == 0 {
			f.ax = float64(pkg.ParseToInt(res[0]))
			f.ay = float64(pkg.ParseToInt(res[1]))
			i++
		} else if i == 1 {
			f.bx = float64(pkg.ParseToInt(res[0]))
			f.by = float64(pkg.ParseToInt(res[1]))
			i++
		} else if i == 2 {
			f.x = float64(pkg.ParseToInt(res[0])) + 10000000000000
			f.y = float64(pkg.ParseToInt(res[1])) + 10000000000000
			i++
		} else if i == 3 {
			funcs = append(funcs, f)
			i = 0
		}
	}

	sum := int64(0)
	for _, f := range funcs {
		sum += calc(f)
	}

	fmt.Println(sum)
}
