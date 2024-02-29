package StringCalculator

import (
	"strconv"
)

func add(text string) (int, error) {
	num, err := strconv.Atoi(text)
	return num, err
}
