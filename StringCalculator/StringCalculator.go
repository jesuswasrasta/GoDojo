package StringCalculator

import (
	"strconv"
	"strings"
)

func add(text string) (int, error) {
	strs := strings.Split(text, ",")

	total := 0
	for _, s := range strs {
		num, _ := strconv.Atoi(s)
		total += num
	}
	return total, nil
}
