package Day14

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"

	"isur.dev/aoc2024/pkg"
)

type Robot struct {
	posX int
	posY int
	velX int
	velY int
}

type Room struct {
	sizeX int
	sizeY int
}

func calcPosition(room Room, robot Robot) Robot {
	robot.posX = robot.posX + robot.velX
	robot.posY = robot.posY + robot.velY

	if robot.posX >= room.sizeX {
		robot.posX = robot.posX - room.sizeX
	}

	if robot.posY >= room.sizeY {
		robot.posY = robot.posY - room.sizeY
	}

	if robot.posX < 0 {
		robot.posX = room.sizeX + robot.posX
	}

	if robot.posY < 0 {
		robot.posY = room.sizeY + robot.posY
	}

	return robot
}

func drawRobots(room [][]int, robots []Robot, p bool) [][]string {
	duplicate := make([][]string, len(room))
	for i := range room {
		duplicate[i] = make([]string, len(room[i]))
		line := []string{}
		for j := range room[i] {
			line = append(line, strconv.Itoa(room[i][j]))
		}
		duplicate = append(duplicate, line)
	}

	for row := range room {
		for col := range room[row] {
			r := 0
			for i := range robots {
				if robots[i].posX == col && robots[i].posY == row {
					r++
				}
			}
			if r > 0 {
				duplicate[row][col] = "#"
			} else {
				duplicate[row][col] = " "
			}
			if p {
				fmt.Print(r)
			}
		}
		if p {
			fmt.Println()
		}
	}

	return duplicate
}

func calcQuadrants(room Room, robots []Robot) int {
	first := 0
	second := 0
	third := 0
	fourth := 0

	hx := room.sizeX / 2
	hy := room.sizeY / 2

	for _, robot := range robots {
		if robot.posX < hx && robot.posY < hy {
			first++
		} else if robot.posX < hx && robot.posY > hy {
			third++
		} else if robot.posX > hx && robot.posY < hy {
			second++
		} else if robot.posX > hx && robot.posY > hy {
			fourth++
		}

	}

	return first * second * third * fourth
}

func save(data [][]string, fileName string) {
	height := len(data)
	width := len(data[0])

	img := image.NewGray(image.Rect(0, 0, width, height))

	for y, row := range data {
		for x, value := range row {
			if value == "#" {
				img.SetGray(x, y, color.Gray{Y: 255}) // Black
			} else {
				img.SetGray(x, y, color.Gray{Y: 0}) // White
			}
		}
	}

	outputFile, err := os.Create("days/Day14/output/" + fileName)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, img)
	if err != nil {
		panic(err)
	}

	fmt.Println("Image saved ", fileName)
}

func Run(data []string) {
	num := pkg.ParseToInt("14")
	fmt.Println("Day ", num)

	room := Room{101, 103}
	arr := [][]int{}
	for range room.sizeY {
		r := []int{}
		for range room.sizeX {
			r = append(r, 0)
		}
		arr = append(arr, r)
	}

	robots := []Robot{}

	for _, line := range data {
		d := strings.Split(line, " ")
		p := strings.Split(d[0], "=")
		pos := strings.Split(p[1], ",")
		v := strings.Split(d[1], "=")
		vel := strings.Split(v[1], ",")

		robot := Robot{
			posX: pkg.ParseToInt(pos[0]),
			posY: pkg.ParseToInt(pos[1]),
			velX: pkg.ParseToInt(vel[0]),
			velY: pkg.ParseToInt(vel[1]),
		}

		robots = append(robots, robot)
	}

	for i := 0; i < 10000; i++ {
		for x := range robots {
			robots[x] = calcPosition(room, robots[x])
		}
		dp := drawRobots(arr, robots, false)
		save(dp, strconv.Itoa(i+1)+".png")
	}

	res := calcQuadrants(room, robots)
	fmt.Println(res)
}
