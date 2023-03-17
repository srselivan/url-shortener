package service

import (
	"math"
	"strings"
)

func resolve(str string) uint {
	var (
		result uint = 0
		length      = len(str)
	)

	for i, ch := range str {
		num := strings.Index(alphabet, string(ch))
		result += uint(math.Pow(float64(alphabetLength), float64(length-i-1)) * float64(num))
	}

	return result
}
