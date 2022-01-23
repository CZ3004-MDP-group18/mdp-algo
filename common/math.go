package common

func Abs(value int) int {
	if value < 0 {
		value = -value
	}
	return value
}
