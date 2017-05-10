package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type handler struct {
	writer Writer
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		month := r.FormValue("month")
		if month == "" {
			month = time.Now().Format("01")
		}
		year := r.FormValue("year")
		if year == "" {
			year = time.Now().Format("2006")
		}
		monthInt, err := strconv.ParseInt(month, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		yearInt, err := strconv.ParseInt(year, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		amount, err := h.writer.AddThisMonth(monthInt, yearInt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(amount)
	case "POST":
		amount, err := strconv.ParseInt(r.FormValue("amount"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if amount > 1000000 {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		if err := h.writer.WriteAmount(amount); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandlerFromEnv() handler {
	return handler{
		WriterFromEnv(),
	}
}
