package Day12

import (
	"cmp"
	"fmt"
	"slices"

	"isur.dev/aoc2024/pkg"
)

type Point struct {
	x int
	y int
	n int
}

type VP struct {
	x int
	y int
}

var visitedPoints = make(map[VP]bool)

func enqueue(queue *[]Point, item *Point) {
	*queue = append(*queue, *item)
}

func dequeue(queue *[]Point) (*Point, error) {
	if len(*queue) == 0 {
		return &Point{}, fmt.Errorf("Queue is empty")
	}

	value := &(*queue)[0]
	*queue = (*queue)[1:]

	return value, nil
}

func isSame(data [][]string, point *Point, newPoint *Point, visited map[VP]bool) bool {
	_, ok := visited[VP{x: newPoint.x, y: newPoint.y}]
	if ok {
		if data[newPoint.x][newPoint.y] == data[point.x][point.y] {
			point.n += 1
		}
		return false
	}
	colSize := len(data)
	rowSize := len(data[0])

	char := data[point.x][point.y]

	if newPoint.x > colSize-1 || newPoint.x < 0 || newPoint.y > rowSize-1 || newPoint.y < 0 {
		return false
	}

	if char == data[newPoint.x][newPoint.y] {
		visited[VP{x: newPoint.x, y: newPoint.y}] = true
		point.n += 1
		return true
	}

	return false
}

func bfs(data [][]string, point *Point) ([]Point, string) {
	queue := []Point{}
	visited := make(map[VP]bool)
	enqueue(&queue, point)
	points := []Point{}
	for len(queue) > 0 {
		p, _ := dequeue(&queue)

		left := Point{x: p.x - 1, y: p.y, n: 0}
		right := Point{x: p.x + 1, y: p.y, n: 0}
		top := Point{x: p.x, y: p.y - 1, n: 0}
		bot := Point{x: p.x, y: p.y + 1, n: 0}

		if isSame(data, p, &left, visited) {
			enqueue(&queue, &left)
		}

		if isSame(data, p, &right, visited) {
			enqueue(&queue, &right)
		}

		if isSame(data, p, &top, visited) {
			enqueue(&queue, &top)
		}

		if isSame(data, p, &bot, visited) {
			enqueue(&queue, &bot)
		}
		visited[VP{x: p.x, y: p.y}] = true
		points = append(points, *p)
	}

	return points, data[point.x][point.y]
}

type Group struct {
	points []Point
	char   string
}

func perimeter(group Group) int {
	l := len(group.points)
	per := 0
	for _, p := range group.points {
		per += 4 - p.n
	}

	return per * l
}

func calcSides(group Group) int {
	maxX := slices.MaxFunc(group.points, func(a, b Point) int {
		return cmp.Compare(a.x, b.x)
	})
	maxY := slices.MaxFunc(group.points, func(a, b Point) int {
		return cmp.Compare(a.y, b.y)
	})
	minX := slices.MinFunc(group.points, func(a, b Point) int {
		return cmp.Compare(a.x, b.x)
	})
	minY := slices.MinFunc(group.points, func(a, b Point) int {
		return cmp.Compare(a.y, b.y)
	})

	sides := 0
	char := group.char
	area := [][]string{}
	for x := minX.x; x <= maxX.x; x++ {
		a := []string{}
		for y := minY.y; y <= maxY.y; y++ {
			a = append(a, "_")
		}
		area = append(area, a)
	}

	for _, poi := range group.points {
		area[poi.x-minX.x][poi.y-minY.y] = group.char
	}

	top := 0
	for col := range area[0] {
		for row := range area {
			curr := area[row][col]
			if row == 0 && col == 0 {
				if curr == char {
					top++
				}
			} else if row == 0 {
				if curr == char && area[row][col-1] == "_" {
					top++
				}
			} else if col == 0 {
				if curr == char && area[row-1][col] == "_" {
					top++
				}
			} else {
				if curr == char && area[row-1][col] == "_" {
					if area[row-1][col-1] == "_" && area[row][col-1] == "_" {
						top++
					}
					if area[row-1][col-1] == char && area[row][col-1] == char {
						top++
					}
					if area[row-1][col-1] == char && area[row][col-1] == "_" {
						top++
					}
				}
			}
		}
	}

	bot := 0
	for col := range area[0] {
		for row := len(area) - 1; row >= 0; row-- {
			curr := area[row][col]
			if row == len(area)-1 && col == 0 {
				if curr == char {
					bot++
				}
			} else if col == 0 {
				if curr == char && area[row+1][col] == "_" {
					bot++
				}
			} else if row == len(area)-1 {
				if curr == char && area[row][col-1] == "_" {
					bot++
				}
			} else {
				if curr == char && area[row+1][col] == "_" {
					if area[row+1][col-1] == "_" && area[row][col-1] == "_" {
						bot++
					}
					if area[row+1][col-1] == char && area[row][col-1] == char {
						bot++
					}
					if area[row+1][col-1] == char && area[row][col-1] == "_" {
						bot++
					}
				}
			}
		}
	}

	left := 0
	for row := range area {
		for col := range area[0] {
			curr := area[row][col]
			if row == 0 && col == 0 {
				if curr == char {
					left++
				}
			} else if row == 0 {
				if curr == char && area[row][col-1] == "_" {
					left++
				}
			} else if col == 0 {
				if curr == char && area[row-1][col] == "_" {
					left++
				}
			} else {
				if curr == char && area[row][col-1] == "_" {
					if area[row-1][col-1] == "_" && area[row-1][col] == "_" {
						left++
					}
					if area[row-1][col-1] == char && area[row-1][col] == char {
						left++
					}
					if area[row-1][col-1] == char && area[row-1][col] == "_" {
						left++
					}
				}
			}
		}
	}

	right := 0
	for row := range area {
		for col := len(area[0]) - 1; col >= 0; col-- {
			curr := area[row][col]
			if row == 0 && col == len(area[0])-1 {
				if curr == char {
					right++
				}
			} else if row == 0 {
				if curr == char && area[row][col+1] == "_" {
					right++
				}
			} else if col == len(area[0])-1 {
				if curr == char && area[row-1][col] == "_" {
					right++
				}
			} else {
				if curr == char && area[row][col+1] == "_" {
					if area[row-1][col+1] == "_" && area[row-1][col] == "_" {
						right++
					}
					if area[row-1][col+1] == char && area[row-1][col] == char {
						right++
					}
					if area[row-1][col+1] == char && area[row-1][col] == "_" {
						right++
					}
				}
			}
		}
	}

	sides = top + bot + left + right

	return sides * len(group.points)
}

func Run(data []string) {
	num := pkg.ParseToInt("12")
	fmt.Println("Day ", num)

	arr := pkg.Array2D(data)

	all := []Group{}

	for x := range arr {
		for y := range arr[x] {
			point := Point{x: x, y: y, n: 0}
			_, ok := visitedPoints[VP{point.x, point.y}]
			if ok {
				continue
			}
			points, char := bfs(arr, &point)
			slices.SortFunc(points, func(a, b Point) int {
				return cmp.Or(
					cmp.Compare(a.x, b.x),
					cmp.Compare(a.y, b.y),
				)
			})

			all = append(all, Group{points: points, char: char})
			for _, p := range points {
				visitedPoints[VP{p.x, p.y}] = true
			}
		}
	}

	sum2 := 0
	for a := range all {
		xd := calcSides(all[a])
		sum2 += xd
	}

	sum := 0
	for _, a := range all {
		sum += perimeter(a)
	}

	fmt.Println(sum, sum2)
}
