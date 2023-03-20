package http

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
	"url-shortener/internal/controller/http/mocks"
)

func TestHandler_GetOriginalUrl(t *testing.T) {
	testTable := []struct {
		key      string
		err      error
		expected struct {
			code int
			body string
		}
	}{
		{
			key: "a",
			err: nil,
			expected: struct {
				code int
				body string
			}{
				code: 200,
				body: "https://google.com",
			},
		},
		{
			key: "-",
			err: errors.New("error"),
			expected: struct {
				code int
				body string
			}{
				code: 500,
				body: "",
			},
		},
	}

	for _, test := range testTable {
		req := httptest.NewRequest("GET", "/"+test.key, nil)
		rr := httptest.NewRecorder()

		svc := mocks.NewService(t)
		svc.On("GetOriginal", "").Return(test.expected.body, test.err)
		handler := New(svc)

		handler.GetOriginalUrl().ServeHTTP(rr, req)

		assert.Equal(t, test.expected.code, rr.Code)

		rBody, _ := io.ReadAll(rr.Body)

		assert.Equal(t, test.expected.body, string(rBody))
	}
}
