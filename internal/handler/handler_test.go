package handler

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-web-app/internal/api"
	"go-web-app/internal/api/mocks"
)

func TestHandler_Hello(t *testing.T) {
	tests := []struct {
		name       string
		customJoke string
		joke       *api.JokeResponse
		err        error
		codeWant   int
		bodyWant   string
	}{
		{
			name:       "Testing simple joke",
			customJoke: "",
			joke:       &api.JokeResponse{Joke: "test joke"},
			err:        nil,
			codeWant:   200,
			bodyWant:   "test joke",
		},
		{
			name:       "Testing custom joke",
			customJoke: "Custom joke",
			joke:       &api.JokeResponse{Joke: "some stupid joke"},
			err:        nil,
			codeWant:   200,
			bodyWant:   "Custom joke",
		},
		{
			name:     "Testing error case",
			joke:     &api.JokeResponse{},
			err:      errors.New("Some stupid error"),
			codeWant: http.StatusInternalServerError,
			bodyWant: http.StatusText(http.StatusInternalServerError),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiMock := mocks.Client{}
			apiMock.On("GetJoke").Return(tt.joke, tt.err)

			h := NewHandler(&apiMock, tt.customJoke)

			req, _ := http.NewRequest("GET", "/", nil)
			rr := httptest.NewRecorder()

			h.Hello(rr, req)

			gotRaw, _ := ioutil.ReadAll(rr.Body)
			got := string(gotRaw)

			if tt.err == nil {
				if got != tt.bodyWant {
					t.Errorf("wrong response body '%s' want '%s'", got, tt.bodyWant)
				}
			}

			if status := rr.Result().StatusCode; status != tt.codeWant {
				t.Errorf("wrong response status '%d' want '%d'", status, tt.codeWant)
			}

		})
	}
}
