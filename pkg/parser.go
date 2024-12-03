package pkg

import (
	"fmt"
	"os"
	"strconv"
)

func ParseToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return num
}
