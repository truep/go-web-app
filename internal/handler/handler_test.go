package handler

import (
	"net/http"
	"testing"
)

func TestHandler_Hello(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *Handler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Hello(tt.args.w, tt.args.r)
		})
	}
}
