package arrayutils

func Contains[T comparable](array []T, value T) bool {
	for i := range array {
		if value == array[i] {
			return true
		}
	}
	return false
}
