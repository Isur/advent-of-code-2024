package pkg

func Abs(x, y int) int {
	if x > y {
		return x - y
	} else {
		return y - x
	}
}
