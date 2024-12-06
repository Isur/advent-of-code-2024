package pkg

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
