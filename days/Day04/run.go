package Day04

import (
	"fmt"
	"regexp"
	"strings"
)

func Array2D(data []string) [][]string {
	var arr [][]string
	for _, line := range data {
		var arr1D []string
		for _, char := range line {
			arr1D = append(arr1D, string(char))
		}
		arr = append(arr, arr1D)
	}

	return arr
}

func lineCheck(line []string) int {
	r, _ := regexp.Compile("XMAS")
	r2, _ := regexp.Compile("SAMX")

	lineString := strings.Join(line, "")

	res := r.FindAllStringSubmatchIndex(lineString, 100)
	res2 := r2.FindAllStringSubmatchIndex(lineString, 100)

	return len(res) + len(res2)
}

func checkDiag(data [][]string) int {
	res := 0
	for d := 0; d < len(data)-3; d++ {
		for i := 0; i < len(data[0])-3; i++ {
			var diag string
			k := d
			for j := i; j < i+4; j++ {
				diag += data[k][j]
				if diag == "XMAS" || diag == "SAMX" {
					res++
				}
				k++
			}
		}
	}

	for d := len(data) - 1; d > 2; d-- {
		for i := 0; i < len(data[0])-3; i++ {
			var diag string
			k := d
			for j := i; j < i+4; j++ {
				diag += data[k][j]
				if diag == "XMAS" || diag == "SAMX" {
					res++
				}
				k--
			}
		}
	}

	return res
}

func checkDoubleMas(data [][]string) int {
	res := 0
	for x := 0; x < len(data)-2; x++ {
		for y := 0; y < len(data[0])-2; y++ {
			xd := false
			var diag string
			for k := 0; k < 3; k++ {
				diag += data[x+k][y+k]
			}
			if diag == "MAS" || diag == "SAM" {
				xd = true
			}
			var diag2 string
			for k := 0; k < 3; k++ {
				diag2 += data[x+k][y+2-k]
			}
			if diag2 == "MAS" || diag2 == "SAM" {
				if xd {
					res++
				}
			}
		}
	}

	return res
}

func Part2(data []string) {
	arr2D := Array2D(data)
	res := checkDoubleMas(arr2D)
	fmt.Println(res)

}

func Part1(data []string) {
	arr2D := Array2D(data)
	sum := 0
	for i := 0; i < len(arr2D); i++ {
		h := lineCheck(arr2D[i])
		sum = sum + h
	}

	for i := 0; i < len(arr2D[0]); i++ {
		var vert []string
		for j := 0; j < len(arr2D); j++ {
			vert = append(vert, arr2D[j][i])
		}
		v := lineCheck(vert)
		sum = sum + v
	}

	sum += checkDiag(arr2D)

	fmt.Println(sum)
}

func Run(data []string) {
	Part1(data)
	Part2(data)
}
