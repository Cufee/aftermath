package utils

func Batch[T any](values []T, size int) [][]T {
	var batched [][]T
	for i := 0; i < len(values); i += size {
		end := min(i+size, len(values))
		batched = append(batched, values[i:end])
	}
	return batched
}
