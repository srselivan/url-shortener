package http

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
	"url-shortener/internal/controller/http/mocks"
	"url-shortener/internal/service/repository"
)

func TestHandler_GetOriginalUrl(t *testing.T) {
	testTable := []struct {
		name     string
		key      string
		err      error
		expected struct {
			code int
			body string
		}
	}{
		{
			name: "correct key",
			key:  "a",
			err:  nil,
			expected: struct {
				code int
				body string
			}{
				code: 200,
				body: "https://google.com",
			},
		},
		{
			name: "incorrect key",
			key:  "-",
			err:  repository.ErrorNotFound,
			expected: struct {
				code int
				body string
			}{
				code: 404,
				body: "",
			},
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/"+test.key, nil)
			rr := httptest.NewRecorder()

			svc := mocks.NewService(t)
			svc.On("GetOriginal", "").Return(test.expected.body, test.err)
			handler := New(svc)

			handler.GetOriginalUrl().ServeHTTP(rr, req)

			assert.Equal(t, test.expected.code, rr.Code)

			rBody, _ := io.ReadAll(rr.Body)

			assert.Equal(t, test.expected.body, string(rBody))
		})
	}
}
