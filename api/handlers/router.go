package handlers

import (
	response "gomboc/api/response"
	"net/http"
)

func MethodNotAllowedHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		response.BadResponse(w, http.StatusMethodNotAllowed)
	}
}

func NotFoundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		response.BadResponse(w, http.StatusNotFound)
	}
}
