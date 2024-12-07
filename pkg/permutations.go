package pkg

func AllPermutations[T any](elements []T, N int) [][]T {
	all := [][]T{}
	var helper func(current []T, depth int)

	helper = func(current []T, depth int) {
		if depth == N {
			perms := make([]T, len(current))
			copy(perms, current)
			all = append(all, perms)
			return
		}

		for _, elem := range elements {
			x := append(current, elem)
			helper(x, depth+1)
		}
	}

	helper([]T{}, 0)

	return all
}
