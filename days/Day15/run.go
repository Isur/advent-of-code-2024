package Day15

import (
	"fmt"
	"strings"

	"isur.dev/aoc2024/pkg"
)

const WALL = "#"
const BOX = "O"
const EMPTY = "."
const G = "@"
const UP = "^"
const DOWN = "v"
const LEFT = "<"
const RIGHT = ">"

func prepareData(data []string) ([]string, [][]string) {
	field := [][]string{}
	movesArr := []string{}

	f := true
	for _, line := range data {
		if line == "" {
			f = false
			continue
		}

		if f {
			items := strings.Split(line, "")
			kek := []string{}
			for _, item := range items {
				if item == WALL {
					kek = append(kek, "#")
					kek = append(kek, "#")
				} else if item == BOX {
					kek = append(kek, "[")
					kek = append(kek, "]")
				} else if item == EMPTY {
					kek = append(kek, ".")
					kek = append(kek, ".")
				} else if item == G {
					kek = append(kek, "@")
					kek = append(kek, ".")
				}
			}
			field = append(field, kek)
		} else {
			movesArr = append(movesArr, line)
		}
	}

	moves := strings.Split(strings.Join(movesArr, ""), "")

	return moves, field
}

type Pos struct {
	x int
	y int
}

func moveBoxRight(field [][]string, pos Pos) bool {
	x := pos.x
	y := pos.y

	for {
		if field[y][x] == "[" && field[y][x+1] == "]" {
			x = x + 2
		} else if field[y][x] == EMPTY {
			field[y][x] = "]"
			field[y][x-1] = "["
			field[y][x-2] = EMPTY
			x -= 2
			if x == pos.x {
				return true
			}
		} else if field[y][x] == WALL {
			return false
		}
	}
}

func moveBoxLeft(field [][]string, pos Pos) bool {
	x := pos.x
	y := pos.y

	for {
		if field[y][x] == "]" && field[y][x-1] == "[" {
			x = x - 2
		} else if field[y][x] == EMPTY {
			field[y][x] = "["
			field[y][x+1] = "]"
			field[y][x+2] = EMPTY
			x += 2
			if x == pos.x {
				return true
			}
		} else if field[y][x] == WALL {
			return false
		}
	}
}

func copyField(field [][]string) [][]string {
	duplicate := make([][]string, len(field))
	for i := range field {
		duplicate[i] = make([]string, len(field[i]))
		copy(duplicate[i], field[i])
	}
	return duplicate
}

func moveBoxUpDown(field [][]string, pos Pos, dir int) bool {
	x := pos.x
	y := pos.y

	if field[y][x] == "[" {
		l := moveBoxUpDown(field, Pos{x, y + dir}, dir)
		r := moveBoxUpDown(field, Pos{x + 1, y + dir}, dir)
		if l && r {
			field[y][x], field[y+dir][x] = field[y+dir][x], field[y][x]
			field[y][x+1], field[y+dir][x+1] = field[y+dir][x+1], field[y][x+1]

			return true
		}
	} else if field[y][x] == "]" {
		r := moveBoxUpDown(field, Pos{x, y + dir}, dir)
		l := moveBoxUpDown(field, Pos{x - 1, y + dir}, dir)
		if l && r {
			field[y][x], field[y+dir][x] = field[y+dir][x], field[y][x]
			field[y][x-1], field[y+dir][x-1] = field[y+dir][x-1], field[y][x-1]

			return true
		}
	}

	if field[y][x] == WALL {
		return false
	}

	if field[y][x] == EMPTY {
		return true
	}

	return false
}

func moving(field [][]string, pos Pos, dir []int) (Pos, [][]string) {
	y := dir[0]
	x := dir[1]

	if field[pos.y+y][pos.x+x] == EMPTY {
		field[pos.y][pos.x] = EMPTY
		field[pos.y+y][pos.x+x] = G
		pos.x = pos.x + x
		pos.y = pos.y + y
	} else if field[pos.y+y][pos.x+x] == WALL {
	} else if field[pos.y+y][pos.x+x] == "]" || field[pos.y+y][pos.x+x] == "[" {
		res := false
		if y == 0 && x == -1 {
			res = moveBoxLeft(field, Pos{x: pos.x + x, y: pos.y + y})
		}
		if y == 0 && x == 1 {
			res = moveBoxRight(field, Pos{x: pos.x + x, y: pos.y + y})
		}
		if y == 1 && x == 0 {
			fieldCopy := copyField(field)
			res = moveBoxUpDown(fieldCopy, Pos{x: pos.x + x, y: pos.y + y}, 1)
			if res {
				field = copyField(fieldCopy)
			}
		}
		if y == -1 && x == 0 {
			fieldCopy := copyField(field)
			res = moveBoxUpDown(fieldCopy, Pos{x: pos.x + x, y: pos.y + y}, -1)
			if res {
				field = copyField(fieldCopy)
			}
		}
		if res {
			field[pos.y][pos.x], field[pos.y+y][pos.x+x] = field[pos.y+y][pos.x+x], field[pos.y][pos.x]
			pos.x = pos.x + x
			pos.y = pos.y + y
		}
	}

	return pos, field
}

func move(field [][]string, move string, pos Pos) (Pos, [][]string) {
	if move == LEFT {
		pos, field = moving(field, pos, []int{0, -1})
	} else if move == UP {
		pos, field = moving(field, pos, []int{-1, 0})
	} else if move == DOWN {
		pos, field = moving(field, pos, []int{1, 0})
	} else if move == RIGHT {
		pos, field = moving(field, pos, []int{0, 1})
	}

	return pos, field
}

func printField(field [][]string) {
	for _, line := range field {
		for _, item := range line {
			fmt.Print(item, " ")
		}
		fmt.Println()
	}
}

func calcGPS(field [][]string) int {
	sum := 0
	for y := range field {
		for x := range field[y] {
			if field[y][x] == "[" {
				sum += 100*y + x
			}
		}
	}

	return sum
}

func Run(data []string) {
	num := pkg.ParseToInt("15")
	fmt.Println("Day ", num)

	moves, field := prepareData(data)

	pos := Pos{0, 0}
	for y, line := range field {
		for x, item := range line {
			if item == G {
				pos.x = x
				pos.y = y
			}
		}
	}

	for _, m := range moves {
		pos, field = move(field, m, pos)
	}

	res := calcGPS(field)
	fmt.Println(res)
}
