package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"url-shortener/internal/service/mocks"
)

func TestShortener(t *testing.T) {
	t.Run("shorten test", func(t *testing.T) {
		testTable := []struct {
			id       uint64
			origUrl  string
			expected string
		}{
			{
				0,
				"https://google.com",
				"http://localhost:8080/a",
			},
			{
				1,
				"https://google.com",
				"http://localhost:8080/b",
			},
		}

		repo := mocks.NewRepository(t)
		s := New("8080", repo)

		for _, test := range testTable {
			repo.On("Set", test.id, test.origUrl).Return(nil)

			got, err := s.Shorten(test.origUrl)
			require.NoError(t, err)

			assert.Equal(t, test.expected, got)
		}

	})

	t.Run("get original test", func(t *testing.T) {
		testTable := []struct {
			id       uint64
			key      string
			expected string
		}{
			{
				0,
				"a",
				"https://google.com",
			},
			{
				1,
				"b",
				"https://habr.com",
			},
		}

		repo := mocks.NewRepository(t)
		s := New("8080", repo)

		for _, test := range testTable {
			repo.On("Get", test.id).Return(test.expected, nil)

			got, err := s.GetOriginal(test.key)
			require.NoError(t, err)

			assert.Equal(t, test.expected, got)
		}

	})
}
