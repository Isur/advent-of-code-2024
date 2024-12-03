package main

import (
	"fmt"
	"os"

	"isur.dev/aoc2024/days/Day01"
	"isur.dev/aoc2024/days/Day02"
	"isur.dev/aoc2024/days/Day03"
	"isur.dev/aoc2024/days/Day04"
	"isur.dev/aoc2024/days/Day05"
	"isur.dev/aoc2024/days/Day06"
	"isur.dev/aoc2024/days/Day07"
	"isur.dev/aoc2024/days/Day08"
	"isur.dev/aoc2024/days/Day09"
	"isur.dev/aoc2024/days/Day10"
	"isur.dev/aoc2024/days/Day11"
	"isur.dev/aoc2024/days/Day12"
	"isur.dev/aoc2024/days/Day13"
	"isur.dev/aoc2024/days/Day14"
	"isur.dev/aoc2024/days/Day15"
	"isur.dev/aoc2024/days/Day16"
	"isur.dev/aoc2024/days/Day17"
	"isur.dev/aoc2024/days/Day18"
	"isur.dev/aoc2024/days/Day19"
	"isur.dev/aoc2024/days/Day20"
	"isur.dev/aoc2024/days/Day21"
	"isur.dev/aoc2024/days/Day22"
	"isur.dev/aoc2024/days/Day23"
	"isur.dev/aoc2024/days/Day24"
	"isur.dev/aoc2024/pkg"
)

type Settings struct {
	day  string
	mode bool
}

func setSettings() Settings {
	args := len(os.Args)
	if args < 2 {
		fmt.Println("Two args required, day and mode")
		os.Exit(1)
	}
	if args > 3 {
		fmt.Println("Too much args")
		fmt.Println("Two args required, day and mode")
		os.Exit(1)
	}

	mode := false
	if args == 3 {
		if os.Args[2] == "true" || os.Args[2] == "input" {
			mode = true
		} else if os.Args[2] == "false" || os.Args[2] == "example" {
			mode = false
		} else {
			fmt.Println("Wrong arg for mode, use one of [true, false, input, example]")
			os.Exit(1)
		}
	}

	settings := Settings{
		day:  os.Args[1],
		mode: mode,
	}

	return settings
}

func displaySettings(settings Settings) {
	fmt.Println("Day:  ", settings.day)
	fmt.Println("Mode: ", settings.mode)
	fmt.Println("==============")
}

func main() {
	settings := setSettings()
	displaySettings(settings)

	data := pkg.LoadByLine(settings.day, settings.mode)
	switch settings.day {
	case "01":
		Day01.Run(data)
	case "02":
		Day02.Run(data)
	case "03":
		Day03.Run(data)
	case "04":
		Day04.Run(data)
	case "05":
		Day05.Run(data)
	case "06":
		Day06.Run(data)
	case "07":
		Day07.Run(data)
	case "08":
		Day08.Run(data)
	case "09":
		Day09.Run(data)
	case "10":
		Day10.Run(data)
	case "11":
		Day11.Run(data)
	case "12":
		Day12.Run(data)
	case "13":
		Day13.Run(data)
	case "14":
		Day14.Run(data)
	case "15":
		Day15.Run(data)
	case "16":
		Day16.Run(data)
	case "17":
		Day17.Run(data)
	case "18":
		Day18.Run(data)
	case "19":
		Day19.Run(data)
	case "20":
		Day20.Run(data)
	case "21":
		Day21.Run(data)
	case "22":
		Day22.Run(data)
	case "23":
		Day23.Run(data)
	case "24":
		Day24.Run(data)
	default:
		fmt.Println("Day ", settings.day, " not exists here")
		os.Exit(1)
	}
}
