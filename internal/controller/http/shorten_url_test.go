package http

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
	"url-shortener/internal/controller/http/mocks"
)

func TestHandler_ShortenUrl(t *testing.T) {
	testTable := []struct {
		name     string
		body     string
		err      error
		expected struct {
			code int
			body string
		}
	}{
		{
			name: "correct req",
			body: "https://google.com",
			err:  nil,
			expected: struct {
				code int
				body string
			}{
				code: 200,
				body: "url",
			},
		},
		{
			name: "incorrect req",
			body: "incorrect body",
			err:  nil,
			expected: struct {
				code int
				body string
			}{
				code: 400,
				body: "",
			},
		},
		{
			name: "service error",
			body: "https://google.com",
			err:  errors.New("error"),
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
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/", bytes.NewBufferString(test.body))
			rr := httptest.NewRecorder()

			svc := mocks.NewService(t)
			if test.body != "incorrect body" {
				svc.On("Shorten", test.body).Return(test.expected.body, test.err)
			}
			handler := New(svc)

			handler.ShortenUrl().ServeHTTP(rr, req)

			assert.Equal(t, test.expected.code, rr.Code)

			rBody, _ := io.ReadAll(rr.Body)

			assert.Equal(t, test.expected.body, string(rBody))
		})
	}
}
