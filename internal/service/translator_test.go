package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTranslator(t *testing.T) {
	t.Run("keyById test", func(t *testing.T) {
		testTable := []struct {
			id       uint64
			expected string
		}{
			{
				id:       1,
				expected: "b",
			},
			{
				id:       62,
				expected: "ba",
			},
			{
				id:       130,
				expected: "cg",
			},
		}

		for _, test := range testTable {
			assert.Equal(t, test.expected, keyById(test.id))
		}
	})

	t.Run("idByKey test", func(t *testing.T) {
		testTable := []struct {
			key      string
			expected uint64
		}{
			{
				key:      "b",
				expected: 1,
			},
			{
				key:      "ba",
				expected: 62,
			},
			{
				key:      "cg",
				expected: 130,
			},
		}

		for _, test := range testTable {
			assert.Equal(t, test.expected, idByKey(test.key))
		}
	})
}
