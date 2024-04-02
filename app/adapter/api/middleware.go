package api

import (
	"log/slog"
	"net/http"
)

type Serve func(w http.ResponseWriter, req *http.Request) error

func Middleware(next Serve, method string) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			slog.Error("invalid request method")
			return
		}
		err := next(w, r)
		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
