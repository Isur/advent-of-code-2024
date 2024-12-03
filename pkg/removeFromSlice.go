package pkg

func RemoveFromSlice[T any](s []T, i int) []T {
	r := make([]T, 0)
	r = append(r, s[:i]...)
	return append(r, s[i+1:]...)
}
