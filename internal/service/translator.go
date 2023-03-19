package service

import (
	"math"
	"strings"
)

const (
	alphabet       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	alphabetLength = uint64(len(alphabet))
)

func keyById(id uint64) string {
	var (
		digits  []uint64
		num     = id
		builder strings.Builder
	)

	if num == 0 {
		return string(alphabet[num])
	}

	for num > 0 {
		digits = append(digits, num%alphabetLength)
		num /= alphabetLength
	}

	reverse(digits)

	for _, digit := range digits {
		builder.WriteString(string(alphabet[digit]))
	}

	return builder.String()
}

func idByKey(key string) uint64 {
	var (
		result uint64 = 0
		length        = len(key)
	)

	for i, ch := range key {
		num := strings.Index(alphabet, string(ch))
		result += uint64(math.Pow(float64(alphabetLength), float64(length-i-1)) * float64(num))
	}

	return result
}

func reverse(s []uint64) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
