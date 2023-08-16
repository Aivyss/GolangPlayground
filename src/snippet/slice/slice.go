package slice

func DeepCopy[T any](slice []T) []T {
	return append([]T{}, slice...)
}

func Remove[T any](slice []T, index int) []T {
	var result []T

	if len(slice)-1 == index {
		result = append([]T{}, slice[:index]...)
	} else {
		left := append([]T{}, slice[:index]...)
		right := append([]T{}, slice[index+1:]...)
		result = append(left, right...)
	}

	return result
}
