package api

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func TraceRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, _ := uuid.NewUUID()
		trace := strings.Replace(id.String(), "-", "", -1)

		// add a request unique identification
		r.Header.Add("x-trace-id", trace)
		w.Header().Add("x-trace-id", trace)
		log.Trace().Msgf("Received a new request signed with id '%s'", trace)

		next.ServeHTTP(w, r)
	})
}

func PrepareResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func ObserverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL
		addr := r.RemoteAddr
		method := r.Method
		trace := r.Header.Get("x-trace-id")

		log.Trace().Msgf("Request %s: '%s' to '%s' from '%s'", trace, method, path, addr)
		next.ServeHTTP(w, r)
	})
}

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
