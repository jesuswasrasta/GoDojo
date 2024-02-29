package StringCalculator

import (
	"strconv"
	"strings"
)

type IStringCalculator interface {
	add(text string) (int, error)
}

type StringCalculator struct {
}

func (c StringCalculator) add(text string) (int, error) {
	numbers := strings.Split(text, ",")

	total := 0
	for _, s := range numbers {
		num, _ := strconv.Atoi(s)
		total += num
	}
	return total, nil
}
