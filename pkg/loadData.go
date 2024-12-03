package pkg

import (
	"bufio"
	"fmt"
	"os"
)

func LoadByLine(day string, mode bool) []string {
	fileName := "example"
	if mode == true {
		fileName = "input"
	}
	filePath := "./days/Day" + day + "/" + fileName
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Cannot open file!")
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var lines []string

	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}

	file.Close()

	return lines
}
